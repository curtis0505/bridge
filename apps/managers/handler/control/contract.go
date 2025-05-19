package control

import (
	"context"
	"fmt"
	eventtypes "github.com/curtis0505/bridge/apps/managers/handler/event/types"
	"github.com/curtis0505/bridge/libs/cache"
	"github.com/curtis0505/bridge/libs/types"
	common "github.com/curtis0505/bridge/libs/types/bridge/abi"
	"math/big"
)

func (p *ControlHandler) Pause(ctx context.Context, chain string) (string, error) {
	event := eventtypes.EventPaused

	if ok := p.CheckPendingTx(chain, event); ok {
		return "", fmt.Errorf("already sent pause transaction")
	}

	data, err := types.PackAbi(chain, common.GetAbiToMap(common.MultiSigWalletAbi), "pause")
	if err != nil {
		return "", fmt.Errorf("PackWithChain: %w", err)
	}

	account := p.account[chain]

	option, err := p.client.GetTransactionOption(ctx, chain, account.Address)
	if err != nil {
		return "", fmt.Errorf("GetTransactionOption: %w", err)
	}

	contract, err := cache.ContractCache().GetContractByContractID(chain, "multisigwallet")
	if err != nil {
		return "", fmt.Errorf("GetContractByContractID: %w", err)
	}

	tx, err := p.client.GetTransactionData(chain, &types.RequestTransaction{
		From:      account.Address,
		To:        contract.Address,
		Nonce:     option.Nonce,
		GasPrice:  option.GasPrice,
		GasFeeCap: option.GasFeeCap,
		GasTipCap: option.GasTipCap,
		GasLimit:  types.GasLimit,
		Value:     big.NewInt(0),
		Data:      data,
	})
	if err != nil {
		return "", fmt.Errorf("GetTransactionData: %w", err)
	}

	chainId, err := p.client.GetChainID(ctx, chain)
	if err != nil {
		return "", fmt.Errorf("GetChainID: %w", err)
	}

	signedTx, err := account.Sign(tx, chainId)
	if err != nil {
		return "", fmt.Errorf("Sign: %w", err)
	}

	resp, err := p.client.RawSendTxAsyncByTx(context.Background(), chain, signedTx)
	if err != nil {
		return "", fmt.Errorf("RawSendTxAsyncByTx: %w", err)
	}

	p.logger.Info(
		"event", "Pause",
		"chain", chain,
		"from", account.Address,
		"txHash", resp.TxHash.String(),
	)

	p.AddPendingTx(types.ChainKLAY, eventtypes.EventPaused, resp.TxHash.String())

	return resp.TxHash.String(), nil
}

func (p *ControlHandler) Unpause(ctx context.Context, chain string) (string, error) {
	event := eventtypes.EventUnpaused

	if ok := p.CheckPendingTx(chain, event); ok {
		return "", fmt.Errorf("already sent unpause transaction")
	}

	data, err := types.PackAbi(chain, common.GetAbiToMap(common.MultiSigWalletAbi), "unpause")
	if err != nil {
		return "", fmt.Errorf("PackWithChain: %w", err)
	}

	account := p.account[chain]

	option, err := p.client.GetTransactionOption(ctx, chain, account.Address)
	if err != nil {
		return "", fmt.Errorf("GetTransactionOption: %w", err)
	}

	contract, err := cache.ContractCache().GetContractByContractID(chain, "multisigwallet")
	if err != nil {
		return "", fmt.Errorf("GetContractByContractID: %w", err)
	}

	tx, err := p.client.GetTransactionData(chain, &types.RequestTransaction{
		From:      account.Address,
		To:        contract.Address,
		Nonce:     option.Nonce,
		GasPrice:  option.GasPrice,
		GasFeeCap: option.GasFeeCap,
		GasTipCap: option.GasTipCap,
		GasLimit:  types.GasLimit,
		Value:     big.NewInt(0),
		Data:      data,
	})
	if err != nil {
		return "", fmt.Errorf("GetTransactionData: %w", err)
	}

	chainId, err := p.client.GetChainID(ctx, chain)
	if err != nil {
		return "", fmt.Errorf("GetChainID: %w", err)
	}

	signedTx, err := account.Sign(tx, chainId)
	if err != nil {
		return "", fmt.Errorf("Sign: %w", err)
	}

	resp, err := p.client.RawSendTxAsyncByTx(ctx, chain, signedTx)
	if err != nil {
		return "", fmt.Errorf("RawSendTxAsyncByTx: %w", err)
	}

	p.logger.Info(
		"event", "Unpause",
		"chain", chain,
		"from", account.Address,
		"txHash", resp.TxHash.String(),
	)

	p.AddPendingTx(types.ChainKLAY, eventtypes.EventUnpaused, resp.TxHash.String())

	return resp.TxHash.String(), nil

}
