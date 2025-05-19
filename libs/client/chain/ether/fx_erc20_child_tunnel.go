package ether

import (
	"errors"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/bridge/abi/ether"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type FxERC20ChildTunnel struct {
	Client  *client
	Address common.Address
	Abi     abi.ABI
}

func newFxERC20ChildTunnel(client *client, address string) *FxERC20ChildTunnel {
	return &FxERC20ChildTunnel{
		Client:  client,
		Address: common.HexToAddress(address),
		Abi:     ether.FxERC20ChildTunnel,
	}
}

func (f *FxERC20ChildTunnel) Withdraw(childTokenAddr string, amount *big.Int, account *types.Account) (*types.Transaction, error) {
	return nil, errors.New("not implemented")
}

func (f *FxERC20ChildTunnel) WithdrawTo(childTokenAddr string, receiverAddr string, amount *big.Int, account *types.Account) (*types.Transaction, error) {
	return nil, errors.New("not implemented")
}

func (f *FxERC20ChildTunnel) MapToken() (*types.Transaction, error) {
	return nil, nil
}

func (f *FxERC20ChildTunnel) GetChainFee() (*big.Int, error) {
	msg, err := f.Client.callMsg(f.Abi, f.Address.String(), "getChainFee")
	if err != nil {
		return nil, err
	}

	chainFee, ok := msg[0].(*big.Int)
	if ok == false {
		return nil, err
	}

	return chainFee, nil
}

func (f *FxERC20ChildTunnel) GetTaxRateBP() (*big.Int, error) {
	msg, err := f.Client.callMsg(f.Abi, f.Address.String(), "taxRateBP")
	if err != nil {
		return nil, err
	}

	taxRateB, ok := msg[0].(*big.Int)
	if ok == false {
		return nil, err
	}

	return taxRateB, nil
}
