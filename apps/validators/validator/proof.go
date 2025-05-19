package validator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/curtis0505/bridge/libs/common"
	"github.com/curtis0505/bridge/libs/dto"
	"github.com/curtis0505/bridge/libs/elog"
	commontypes "github.com/curtis0505/bridge/libs/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types/bridge"
	"github.com/curtis0505/bridge/libs/types/token"
	klaycommon "github.com/klaytn/klaytn/common"
	"io"
	"net/http"
)

type ProofResp struct {
	Proof string `json:"proof"`
}

func (p *Validator) getProof(commonLog bridgetypes.CommonEventLog, contractMethod Method) ([]byte, error) {
	if len(p.cfg.Server.ProofURL) == 0 {
		return []byte{}, nil
	}

	fromChainContract, err := p.Contract.GetChain(commontypes.Chain(commonLog.FromChainName))
	if err != nil {
		return nil, err
	}

	toChainContract, err := p.Contract.GetChain(commontypes.Chain(commonLog.ToChainName))
	if err != nil {
		return nil, err
	}

	sendTokenAddr := commonLog.FromTokenAddr.String()
	elog.Debug("getProof", "commonLog", commonLog, "FromChainName", commonLog.FromChainName, "sendTokenAddr", sendTokenAddr)

	var vaultAddress, minterAddress, msgReceiver string

	switch commonLog.Version {
	case bridgetypes.VersionBridge:
		vaultAddress = fromChainContract.Vault
		minterAddress = fromChainContract.Minter

		switch contractMethod.ExecuteMethod {
		case bridgetypes.MethodNameBridgeBurn:
			msgReceiver = toChainContract.Minter
		case bridgetypes.MethodNameBridgeWithdraw:
			msgReceiver = toChainContract.Vault
		}
	case bridgetypes.VersionRestakeBridge:
		vaultAddress = fromChainContract.RestakeVault
		minterAddress = fromChainContract.RestakeMinter

		switch contractMethod.ExecuteMethod {
		case bridgetypes.MethodNameBridgeBurn:
			msgReceiver = toChainContract.RestakeMinter
		case bridgetypes.MethodNameBridgeWithdraw:
			msgReceiver = toChainContract.RestakeVault
		}
	default:
		return nil, fmt.Errorf("commonLog.Version is empty")
	}

	var vault, minter, txHash []byte
	if commontypes.GetChainType(commonLog.FromChainName) == commontypes.ChainTypeCOSMOS {
		//response := bridge.QueryCoinResponse{} //cosmos 컨트랙트 상 vault만 coin bridge 가능
		//p.client.CallWasm(context.Background(), commonLog.FromChainName, p.client.Vault(commonLog.FromChainName).GetAddress(), bridge.QueryCoinRequest{
		//	ChildChainName: commonLog.ToChainName,
		//	ChildTokenAddr: strings.ToLower(commonLog.ToTokenAddr.String()),
		//}, &response)
		//
		////coin인 경우, sendTokenAddr denom 으로 변경
		//if response.NativeDenom == token.DenomByChain(commonLog.FromChainName) {
		//	sendTokenAddr = token.DenomByChain(commonLog.FromChainName)
		//}

		if commonLog.FromTokenAddr.String() == bridgetypes.CoinAddress {
			sendTokenAddr = token.DenomByChain(commonLog.FromChainName)
		}

		_, vault, err = bech32.DecodeAndConvert(vaultAddress)
		if err != nil {
			return nil, fmt.Errorf("bech32 decode error: %v", err)
		}
		_, minter, err = bech32.DecodeAndConvert(minterAddress)
		if err != nil {
			return nil, fmt.Errorf("bech32 decode error: %v", err)
		}

		txHash = klaycommon.Hex2Bytes(commonLog.TxHash[:])
	} else {
		vault = common.HexToAddress(commonLog.ToChainName, vaultAddress).Bytes()
		minter = common.HexToAddress(commonLog.ToChainName, minterAddress).Bytes()
		txHash = klaycommon.Hex2Bytes(commonLog.TxHash[2:])
	}

	verification := dto.ProofBridgeReq{
		SenderVault:  vault,
		SenderMinter: minter,
		SenderHash:   txHash,
		//SenderHash:      commonLog.TxHash,
		SenderChainName: commonLog.FromChainName,
		MsgReceiver:     msgReceiver,
		SendTokenAddr:   sendTokenAddr,
		//RcvTokenAddr:    commonLog.TokenAddr,
		RcvAmount:    commonLog.Amount,
		RcvChainName: commonLog.ToChainName,
		Version:      commonLog.Version,
	}

	proof, err := bridgeProof(
		verification,
		p.cfg.Server.ProofURL,
		p.Account[commonLog.ToChainName].Address,
	)
	if err != nil {
		return []byte{}, err
	}

	return proof, nil
}

func bridgeProof(verification dto.ProofBridgeReq, path, address string) ([]byte, error) {
	body, err := json.Marshal(verification)
	if err != nil {
		return []byte{}, err
	}
	url := fmt.Sprintf("%s/%s", path, address)
	respBody, err := getProofApi(body, url)
	if err != nil {
		return []byte{}, err
	}

	proofResp := ProofResp{}
	err = json.Unmarshal(respBody, &proofResp)
	if err != nil {
		return []byte{}, err
	}

	proof, err := commontypes.HexUtilDecode(verification.RcvChainName, proofResp.Proof)
	if err != nil {
		return []byte{}, err
	}

	return proof, nil
}

func getProofApi(body []byte, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(body))
	if err != nil {
		return []byte{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp == nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return respBody, nil
}
