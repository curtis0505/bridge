package rest

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/curtis0505/bridge/libs/types"
	commontypes "github.com/curtis0505/bridge/libs/types"
	bridgetypes "github.com/curtis0505/bridge/libs/types/cosmos/bridge"
	"math/big"
	"strings"
)

type Vault struct {
	client  *client
	address string
}

func newVault(client *client, address string) *Vault {
	return &Vault{
		client:  client,
		address: address,
	}
}

func (v *Vault) GetAddress() string {
	return v.address
}

func (v *Vault) GetChainId(chainName string) (string, error) {
	return "", fmt.Errorf("not supported")
}

func (v *Vault) GetChainFee(chainName string) (*big.Int, error) {
	response := bridgetypes.QueryChildChainResponse{}
	err := v.client.CallWasm(context.Background(), v.address, bridgetypes.QueryChildChainRequest{
		ChildChainName: chainName, //vault 기준 ETH는 child chain
	}, &response)
	if err != nil {
		return nil, commontypes.WrapError("CallWasm", err)
	}

	fee, ok := new(big.Int).SetString(response.Fee, 10)
	if !ok {
		return nil, commontypes.WrapError("SetString", fmt.Errorf("chain fee convert error"))
	}

	return fee, nil
}

func (v *Vault) GetTaxRateBP() (*big.Int, error) {
	response := bridgetypes.QueryConfigResponse{}
	err := v.client.CallWasm(context.Background(), v.address, bridgetypes.QueryConfigRequest{}, &response)
	if err != nil {
		return nil, commontypes.WrapError("CallWasm", err)
	}

	tax, ok := new(big.Float).SetString(response.TaxRate)
	if !ok {
		return nil, commontypes.WrapError("SetString", fmt.Errorf("chain fee convert error"))
	}
	taxRateBp, ok := new(big.Int).SetString(tax.Mul(tax, big.NewFloat(10000)).String(), 10)
	if !ok {
		return nil, commontypes.WrapError("SetString", fmt.Errorf("chain fee convert error"))
	}
	return taxRateBp, nil
}

func (v *Vault) IsSupportChain(chainName string) (bool, error) {
	response := bridgetypes.QueryChildChainResponse{}
	err := v.client.CallWasm(context.Background(), v.address, bridgetypes.QueryChildChainRequest{
		ChildChainName: chainName, //vault 기준 ETH는 child chain
	}, &response)
	if err != nil {
		return false, commontypes.WrapError("CallWasm", err)
	}

	return response.Support, nil
}

func (v *Vault) GetInputDataWithdraw(fromChainName string, from []byte, toAddr string, token []byte, chainId string, txHash string, amount, decimal *big.Int) ([]byte, error) {
	return []byte{}, fmt.Errorf("not supported")
}

func (v *Vault) Deposit(toChainName string, to []byte, amount *big.Int, proof []byte, account *commontypes.Account) (*commontypes.Transaction, error) {
	return nil, fmt.Errorf("not supported")
}

func (v *Vault) DepositToken(tokenAddr string, toChainName string, to []byte, amount *big.Int, proof []byte, account *types.Account) (*types.Transaction, error) {
	return nil, fmt.Errorf("not supported")
}

func (v *Vault) GetVaultAddressByProof() []byte {
	_, vault, _ := bech32.DecodeAndConvert(strings.ToLower(v.GetAddress()))
	return vault
}

func (v *Vault) GetFeeTax(toChainName string) (*big.Int, *big.Int, error) {
	chainFee, err := v.GetChainFee(toChainName)
	if err != nil {
		return nil, nil, err
	}

	taxRateBp, err := v.GetTaxRateBP()
	if err != nil {
		return nil, nil, err
	}
	return chainFee, taxRateBp, nil
}
