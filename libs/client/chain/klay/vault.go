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

type Vault struct {
	Client  *client
	Address common.Address
	Abi     abi.ABI
}

func newVault(client *client, address string) *Vault {
	return &Vault{
		Client:  client,
		Address: common.HexToAddress(address),
		Abi:     klay.Vault,
	}
}

func (v *Vault) GetAddress() string {
	return v.Address.String()
}

func (v *Vault) GetChainId(chainName string) (string, error) {
	msg, err := v.Client.callMsg(v.Abi, v.Address.String(), "getChainId", chainName)
	if err != nil {
		return "", err
	}

	chainId := msg[0].([32]byte)
	return hexutil.Encode(chainId[:]), nil
}

func (v *Vault) GetChainFee(chainId string) (*big.Int, error) {
	var decodeData [32]byte

	_decodeData, err := hexutil.Decode(chainId)
	if err != nil {
		return nil, err
	}

	copy(decodeData[:], _decodeData[:])

	msg, err := v.Client.callMsg(v.Abi, v.Address.String(), "getChainFee", decodeData)
	if err != nil {
		return nil, err
	}

	chainFee, ok := msg[0].(*big.Int)
	if ok == false {
		return nil, err
	}

	return chainFee, nil
}

func (v *Vault) GetTaxRateBP() (*big.Int, error) {
	msg, err := v.Client.callMsg(v.Abi, v.Address.String(), "taxRateBP")
	if err != nil {
		return nil, err
	}

	taxRateB, ok := msg[0].(*big.Int)
	if ok == false {
		return nil, err
	}

	return taxRateB, nil
}

func (v *Vault) IsSupportChain(chainId string) (bool, error) {

	chainIdB, err := hexutil.Decode(chainId)
	if err != nil {
		return false, err
	}
	chainIdByte32 := [32]byte{}
	copy(chainIdByte32[:], chainIdB)

	msg, err := v.Client.callMsg(v.Abi, v.Address.String(), "isSupportChain", chainIdByte32)
	if err != nil {
		return false, err
	}

	return msg[0].(bool), nil
}

func (v *Vault) GetInputDataWithdraw(fromChainName string, from []byte, toAddr string, token []byte, chainId string, txHash string, amount, decimal *big.Int) ([]byte, error) {
	txInfo1 := [32]byte{}
	txInfo2 := [32]byte{}
	chainIdB, _ := hexutil.Decode(chainId)
	copy(txInfo1[:], chainIdB)
	copy(txInfo2[:], common.HexToHash(txHash).Bytes())
	txInfo := [][32]byte{txInfo1, txInfo2}

	tokenInfo := make([]*big.Int, 0)
	tokenInfo = append(tokenInfo, amount)
	tokenInfo = append(tokenInfo, decimal)

	inputData, err := v.Abi.Pack("withdraw", fromChainName, from, common.HexToAddress(toAddr), common.BytesToAddress(token), txInfo, tokenInfo)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return inputData, nil
}

func (v *Vault) Deposit(toChainName string, to []byte, amount *big.Int, proof []byte, account *types.Account) (*types.Transaction, error) {
	return nil, errors.New("not implemented")
}

func (v *Vault) DepositToken(tokenAddr string, toChainName string, to []byte, amount *big.Int, proof []byte, account *types.Account) (*types.Transaction, error) {
	return nil, errors.New("not implemented")
}

func (v *Vault) GetVaultAddressByProof() []byte {
	return common.HexToAddress(v.GetAddress()).Bytes()
}

func (v *Vault) GetFeeTax(toChainName string) (*big.Int, *big.Int, error) {
	chainId, err := v.GetChainId(toChainName)
	if err != nil {
		return nil, nil, err
	}

	chainFee, err := v.GetChainFee(chainId)
	if err != nil {
		return nil, nil, err
	}

	taxRateBp, err := v.GetTaxRateBP()
	if err != nil {
		return nil, nil, err
	}
	return chainFee, taxRateBp, nil
}
