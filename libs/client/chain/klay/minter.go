package klay

import (
	"errors"
	"fmt"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/bridge/abi/klay"
	"github.com/kaiachain/kaia/accounts/abi"
	"github.com/kaiachain/kaia/common"
	"github.com/kaiachain/kaia/common/hexutil"
	"math/big"
)

type Minter struct {
	Client  *client
	Address common.Address
	Abi     abi.ABI
}

func newMinter(client *client, address string) *Minter {
	return &Minter{
		Client:  client,
		Address: common.HexToAddress(address),
		Abi:     klay.Minter,
	}
}

func (m *Minter) GetAddress() string {
	return m.Address.String()
}

func (m *Minter) GetTokenAddress(chainName string, token []byte) (common.Address, error) {
	msg, err := m.Client.callMsg(m.Abi, m.Address.String(), "getTokenAddress", chainName, token)
	if err != nil {
		return common.Address{}, err
	}

	result := msg[0].(common.Address)
	return result, nil
}

func (m *Minter) GetChainId(chainName string) (string, error) {
	msg, err := m.Client.callMsg(m.Abi, m.Address.String(), "getChainId", chainName)
	if err != nil {
		return "", err
	}

	chainId := msg[0].([32]byte)
	return hexutil.Encode(chainId[:]), nil
}

func (m *Minter) GetChainFee(chainId string) (*big.Int, error) {
	var decodeData [32]byte

	_decodeData, err := hexutil.Decode(chainId)
	if err != nil {
		return nil, err
	}

	copy(decodeData[:], _decodeData[:])

	msg, err := m.Client.callMsg(m.Abi, m.Address.String(), "getChainFee", decodeData)
	if err != nil {
		return nil, err
	}

	chainFee, ok := msg[0].(*big.Int)
	if ok == false {
		return nil, err
	}

	return chainFee, nil
}

func (m *Minter) GetTaxRateBP() (*big.Int, error) {
	msg, err := m.Client.callMsg(m.Abi, m.Address.String(), "taxRateBP")
	if err != nil {
		return nil, err
	}

	taxRateB, ok := msg[0].(*big.Int)
	if ok == false {
		return nil, err
	}

	return taxRateB, nil
}

func (m *Minter) IsSupportChain(chainId string) (bool, error) {
	chainIdB, err := hexutil.Decode(chainId)
	if err != nil {
		return false, err
	}
	chainIdByte32 := [32]byte{}
	copy(chainIdByte32[:], chainIdB)

	msg, err := m.Client.callMsg(m.Abi, m.Address.String(), "isSupportChain", chainIdByte32)
	if err != nil {
		return false, err
	}

	return msg[0].(bool), nil
}

func (m *Minter) GetInputDataMint(fromChainName string, from []byte, toAddr string, token []byte, chainId string, txHash string, amount, decimal *big.Int) ([]byte, error) {
	txInfo1 := [32]byte{}
	txInfo2 := [32]byte{}
	chainIdB, _ := hexutil.Decode(chainId)
	copy(txInfo1[:], chainIdB)
	copy(txInfo2[:], common.HexToHash(txHash).Bytes())
	txInfo := [][32]byte{txInfo1, txInfo2}

	tokenInfo := make([]*big.Int, 0)
	tokenInfo = append(tokenInfo, amount)
	tokenInfo = append(tokenInfo, decimal)

	inputData, err := m.Abi.Pack("mint", fromChainName, from, common.HexToAddress(toAddr), token, txInfo, tokenInfo)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return inputData, nil
}

func (m *Minter) Burn(tokenAddr string, toChainName string, to []byte, amount *big.Int, coinAmount *big.Int, proof []byte, account *types.Account) (*types.Transaction, error) {
	return nil, errors.New("not implemented")
}

func (m *Minter) GetMinterAddressByProof() []byte {
	return common.HexToAddress(m.GetAddress()).Bytes()
}

func (m *Minter) GetFeeTax(toChainName string) (*big.Int, *big.Int, error) {
	chainId, err := m.GetChainId(toChainName)
	if err != nil {
		return nil, nil, err
	}

	chainFee, err := m.GetChainFee(chainId)
	if err != nil {
		return nil, nil, err
	}

	taxRateBp, err := m.GetTaxRateBP()
	if err != nil {
		return nil, nil, err
	}

	return chainFee, taxRateBp, nil
}
