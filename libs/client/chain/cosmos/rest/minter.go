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

type Minter struct {
	client  *client
	address string
}

func newMinter(client *client, address string) *Minter {
	return &Minter{
		client:  client,
		address: address,
	}
}

func (m *Minter) GetAddress() string {
	return m.address
}

func (m *Minter) GetChainId(chainName string) (string, error) {
	return "", fmt.Errorf("not supported")
}

func (m *Minter) GetChainFee(chainName string) (*big.Int, error) {
	response := bridgetypes.QueryParentChainResponse{}
	err := m.client.CallWasm(context.Background(), m.address, bridgetypes.QueryParentChainRequest{
		ParentChainName: chainName, //minter 기준 parent는 ETH
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

func (m *Minter) GetTaxRateBP() (*big.Int, error) {
	response := bridgetypes.QueryConfigResponse{}
	err := m.client.CallWasm(context.Background(), m.address, bridgetypes.QueryConfigRequest{}, &response)
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

func (m *Minter) IsSupportChain(chainName string) (bool, error) {
	response := bridgetypes.QueryParentChainResponse{}
	err := m.client.CallWasm(context.Background(), m.address, bridgetypes.QueryParentChainRequest{
		ParentChainName: chainName,
	}, &response)
	if err != nil {
		return false, commontypes.WrapError("CallWasm", err)
	}

	return response.Support, nil
}

func (m *Minter) GetInputDataMint(fromChainName string, from []byte, toAddr string, token []byte, chainId string, txHash string, amount, decimal *big.Int) ([]byte, error) {
	return []byte{}, fmt.Errorf("not supported")
}

func (m *Minter) Burn(tokenAddr string, toChainName string, to []byte, amount *big.Int, coinAmount *big.Int, proof []byte, account *types.Account) (*types.Transaction, error) {
	return nil, fmt.Errorf("not supported")
}

func (m *Minter) GetMinterAddressByProof() []byte {
	_, minter, _ := bech32.DecodeAndConvert(strings.ToLower(m.GetAddress()))
	return minter
}

func (m *Minter) GetFeeTax(toChainName string) (*big.Int, *big.Int, error) {
	chainFee, err := m.GetChainFee(toChainName)
	if err != nil {
		return nil, nil, err
	}

	taxRateBp, err := m.GetTaxRateBP()
	if err != nil {
		return nil, nil, err
	}

	return chainFee, taxRateBp, nil
}
