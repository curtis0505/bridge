package ether

import (
	"context"
	"errors"
	"fmt"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/bridge/abi/ether"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethertypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type FxERC20RootTunnel struct {
	Client  *client
	Address common.Address
	Abi     abi.ABI
}

func newFxERC20RootTunnel(client *client, address string) *FxERC20RootTunnel {
	return &FxERC20RootTunnel{
		Client:  client,
		Address: common.HexToAddress(address),
		Abi:     ether.FxERC20RootTunnel,
	}
}

func (f *FxERC20RootTunnel) Deposit(tokenAddr string, user []byte, amount *big.Int, data []byte, account *types.Account) (*types.Transaction, error) {
	return nil, errors.New("not implemented")
}

// ReceiveMessage is an exit contract call for withdraw flow
func (f *FxERC20RootTunnel) ReceiveMessage(inputData []byte, account *types.Account) (*types.Transaction, error) {
	ctx := context.Background()

	data, err := f.Abi.Pack("receiveMessage", inputData)
	if err != nil {
		return nil, fmt.Errorf("pack : %w", err)
	}

	nonce, err := f.Client.PendingNonceAt(ctx, account.Address)
	if err != nil {
		return nil, fmt.Errorf("PendingNonceAt: %w", err)
	}

	option, err := f.Client.GetTransactionOption(ctx, account.Address)
	if err != nil {
		return nil, fmt.Errorf("GetTransactionOption: %w", err)
	}

	toAddress := common.HexToAddress(f.Address.String())
	tx := types.NewTx(types.ChainETH, &ethertypes.DynamicFeeTx{
		Nonce:     nonce,
		GasFeeCap: option.GasFeeCap,
		GasTipCap: option.GasTipCap,
		Gas:       types.GasLimit,
		To:        &toAddress,
		Value:     big.NewInt(0),
		Data:      data,
	})

	chainID, err := f.Client.GetChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("GetChainID: %w", err)
	}

	signedTx, err := account.Sign(tx, chainID)
	if err != nil {
		return nil, fmt.Errorf("sign: %w", err)
	}

	_, err = f.Client.RawTxAsyncByTx(ctx, signedTx)
	if err != nil {
		return nil, fmt.Errorf("RawTxAsyncByTx: %w", err)
	}

	return signedTx, nil
}

func (f *FxERC20RootTunnel) GetChainFee() (*big.Int, error) {
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

func (f *FxERC20RootTunnel) GetTaxRateBP() (*big.Int, error) {
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
