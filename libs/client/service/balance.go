package service

import (
	"context"
	"fmt"
	"github.com/curtis0505/bridge/libs/account"
	"github.com/curtis0505/bridge/libs/cache"
	"github.com/curtis0505/bridge/libs/client/chain/cosmos"
	"github.com/curtis0505/bridge/libs/common"
	mongoServiceDB "github.com/curtis0505/bridge/libs/database/mongo/service_db"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"github.com/curtis0505/bridge/libs/multicall"
	"github.com/curtis0505/bridge/libs/types"
	"github.com/curtis0505/bridge/libs/types/cosmos/cw20"
	tokentypes "github.com/curtis0505/bridge/libs/types/token"
	"github.com/curtis0505/bridge/libs/util"
	"github.com/shopspring/decimal"
	"math/big"
	"sort"
	"strings"
	"sync"
	"time"
)

// BalanceAll
type BalanceAll struct {
	*sync.WaitGroup
	mutex   *sync.RWMutex
	Balance []*Balance
}

func NewBalanceAll() *BalanceAll {
	return &BalanceAll{
		WaitGroup: &sync.WaitGroup{},
		mutex:     &sync.RWMutex{},
		Balance:   make([]*Balance, 0),
	}
}

func (b *BalanceAll) Len() int {
	return len(b.Balance)
}

func (b *BalanceAll) Less(i, j int) bool {
	if b.Balance[i].Tokens.Order != b.Balance[j].Tokens.Order {
		return b.Balance[i].Tokens.Order < b.Balance[j].Tokens.Order
	}

	if b.Balance[i].Tokens.ChainOrder != b.Balance[j].Tokens.ChainOrder {
		return b.Balance[i].Tokens.ChainOrder < b.Balance[j].Tokens.ChainOrder
	}

	return b.Balance[i].Tokens.CurrencyID < b.Balance[j].Tokens.CurrencyID
}

func (b *BalanceAll) Swap(i, j int) {
	b.Balance[i], b.Balance[j] = b.Balance[j], b.Balance[i]
}

func (b *BalanceAll) AddBalance(balance *Balance) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if balance.Amount == nil {
		balance.Amount = big.NewInt(0)
	}
	b.Balance = append(b.Balance, balance)
}

// TotalPrice returns all [(amount * token) price]
func (b *BalanceAll) TotalPrice() float64 {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	var total float64
	for _, balance := range b.Balance {
		total += balance.Price()
	}

	return total
}

func (p *Client) BalanceAll(ctx context.Context, addressInfo account.AddressInfo) (*BalanceAll, error) {
	balanceAll := NewBalanceAll()

	for _, chain := range p.GetChains() {
		balanceAll.Add(1)
		go func(chain string) {
			defer balanceAll.Done()

			timeOutCtx, cancel := context.WithTimeout(ctx, time.Second*5)
			defer cancel()

			coinInfo, err := cache.TokenCache().GetTokensByCurrencyID(ctx, tokentypes.CurrencyIdByChain(chain))
			if err != nil {
				logger.Error("BalanceAll", logger.BuildLogInput().WithError(fmt.Errorf("GetTokensByCurrencyID: %w", err)).WithChain(chain))
				return
			}

			balance, err := p.BalanceAt(timeOutCtx, chain, addressInfo.GetAddressByChain(chain), nil)
			if err != nil {
				logger.Error("BalanceAll", logger.BuildLogInput().WithError(fmt.Errorf("BalanceAt: %w", err)).WithChain(chain).WithAddress(addressInfo.GetAddressByChain(chain)))
			}

			balanceAll.AddBalance(NewBalance(
				coinInfo,
				balance,
				cache.PriceCache().GetPriceBySymbolWithNoErr(coinInfo.Symbol),
				err,
			))

			tokenList := cache.TokenCache().GetTokensListByChain(ctx, chain, 2)

			if len(tokenList) == 0 {
				return
			}

			c, err := cache.ContractCache().GetContractByContractID(ctx, chain, multicall.ContractID)
			if err != nil {
				logger.Error("BalanceAll", logger.BuildLogInput().WithError(fmt.Errorf("GetContractByContractID: %w", err)).WithChain(chain))
				return
			}
			switch types.GetChainType(chain) {
			// TODO: 체인 추가시 체크 필요
			case types.ChainTypeEVM:
				multiCall := multicall.New()
				balanceOf := make([]tokentypes.OutputBalanceOf, len(tokenList))
				for i, token := range tokenList {
					contract, err := cache.ContractCache().GetContractByAddress(ctx, chain, token.Address)
					if err != nil {
						logger.Error("BalanceAll", logger.BuildLogInput().WithError(fmt.Errorf("GetContractByAddress: %w", err)).WithChain(chain).WithAddress(addressInfo.GetAddressByChain(chain)))
						continue
					}

					err = multiCall.AddCallUnmarshaler(chain, token.Address, "balanceOf", contract.ABI, &balanceOf[i], common.HexToAddress(chain, addressInfo.GetAddressByChain(chain)))
					if err != nil {
						logger.Error("BalanceAll", logger.BuildLogInput().WithError(fmt.Errorf("AddCallUnmarshaler: %w", err)).WithChain(chain).WithAddress(addressInfo.GetAddressByChain(chain)))
						continue
					}
				}
				err = p.CallMsgUnmarshal(timeOutCtx, chain, "", c.Address, multicall.Aggregate, multicall.Abi, multiCall, multiCall)
				if err != nil {
					logger.Error("BalanceAll", logger.BuildLogInput().WithError(fmt.Errorf("CallMsgUnmarshal: %w", err)).WithChain(chain).WithAddress(addressInfo.GetAddressByChain(chain)))
				}

				for i, tokenInfo := range tokenList {
					if common.EmptyAddress(common.HexToAddress(chain, addressInfo.GetAddressByChain(chain))) {
						balanceOf[i].Amount = big.NewInt(0)
					}

					balanceAll.AddBalance(NewBalance(
						tokenInfo,
						balanceOf[i].Amount,
						cache.PriceCache().GetPriceBySymbolWithNoErr(tokenInfo.Symbol),
						err,
					))
				}

			case types.ChainTypeTVM:
				multiCall := multicall.New()
				balanceOf := make([]tokentypes.OutputBalanceOf, len(tokenList))
				for i, token := range tokenList {
					contract, err := cache.ContractCache().GetContractByAddress(ctx, chain, token.Address)
					if err != nil {
						logger.Error("BalanceAll", logger.BuildLogInput().WithError(fmt.Errorf("GetContractByAddress: %w", err)).WithChain(chain).WithAddress(addressInfo.GetAddressByChain(chain)))
						continue
					}

					err = multiCall.AddCallUnmarshaler(chain, token.Address, "balanceOf", contract.ABI, &balanceOf[i], common.HexToAddress(chain, addressInfo.GetAddressByChain(chain)))
					if err != nil {
						logger.Error("BalanceAll", logger.BuildLogInput().WithError(fmt.Errorf("AddCallUnmarshaler: %w", err)).WithChain(chain).WithAddress(addressInfo.GetAddressByChain(chain)))
						continue
					}
				}
				err = p.CallMsgUnmarshal(timeOutCtx, chain, "", c.Address, multicall.Aggregate, multicall.Abi, multiCall, multiCall)
				if err != nil {
					fmt.Println(err)
					logger.Error("BalanceAll", logger.BuildLogInput().WithError(fmt.Errorf("CallMsgUnmarshal: %w", err)).WithChain(chain).WithAddress(addressInfo.GetAddressByChain(chain)))
				}

				for i, tokenInfo := range tokenList {
					if common.EmptyAddress(common.HexToAddress(chain, addressInfo.GetAddressByChain(chain))) {
						balanceOf[i].Amount = big.NewInt(0)
					}

					balanceAll.AddBalance(NewBalance(
						tokenInfo,
						balanceOf[i].Amount,
						cache.PriceCache().GetPriceBySymbolWithNoErr(tokenInfo.Symbol),
						err,
					))
				}
			case types.ChainTypeCOSMOS:
				client, err := cosmos.ProxyClient(p, chain)
				if err != nil {
					logger.Error("BalanceAll", logger.BuildLogInput().WithError(fmt.Errorf("ProxyClient: %w", err)).WithChain(chain).WithAddress(addressInfo.GetAddressByChain(chain)))
					return
				}

				coins, err := client.Balances(timeOutCtx, addressInfo.GetAddressByChain(chain))
				if err != nil {
					logger.Error("BalanceAll", logger.BuildLogInput().WithError(fmt.Errorf("GetAddressByChain: %w", err)).WithChain(chain).WithAddress(addressInfo.GetAddressByChain(chain)))
				}

				for _, tokenInfo := range tokenList {
					switch tokenInfo.Type {
					case mongoServiceDB.TokenTypeCoin:
						ok, coin := coins.Find(tokenInfo.Denom)
						if !ok {
							balanceAll.AddBalance(NewBalance(
								tokenInfo,
								big.NewInt(0),
								cache.PriceCache().GetPriceBySymbolWithNoErr(tokenInfo.Symbol),
								err,
							))
						} else {
							balanceAll.AddBalance(NewBalance(
								tokenInfo,
								coin.Amount.BigInt(),
								cache.PriceCache().GetPriceBySymbolWithNoErr(tokenInfo.Symbol),
								err,
							))
						}
					case mongoServiceDB.TokenTypeToken:
						// CW-20
						var queryBalanceResp cw20.QueryBalanceResponse
						err := client.CallWasm(timeOutCtx, tokenInfo.Address, cw20.QueryBalanceRequest{
							Address: addressInfo.GetAddressByChain(chain),
						}, &queryBalanceResp)
						if err != nil {
							logger.Error("BalanceAll", logger.BuildLogInput().WithError(fmt.Errorf("CallWasm: %w", err)).WithChain(chain).WithAddress(addressInfo.GetAddressByChain(chain)).WithData("contractAddress", tokenInfo.Address))
						}
						cw20Balance := util.ToDecimal(queryBalanceResp.Balance, 0)
						// i.e) NPT.fns -> NPT
						priceSymbol := strings.Split(tokenInfo.Symbol, ".")[0]
						balanceAll.AddBalance(NewBalance(
							tokenInfo,
							cw20Balance.BigInt(),
							cache.PriceCache().GetPriceBySymbolWithNoErr(priceSymbol),
							err,
						))
					default:
						// TODO: define IBC token type
						logger.Warn("Balance", logger.BuildLogInput().WithError(fmt.Errorf("unknown token type")).WithChain(tokenInfo.Chain).WithData("contractAddress", tokenInfo.Type, "contractAddress", tokenInfo.Address))
					}
				}
			}
		}(chain)
	}
	balanceAll.Wait()
	sort.Sort(balanceAll)

	return balanceAll, nil
}

type Balance struct {
	Tokens     *mongoServiceDB.Tokens
	Amount     *big.Int
	TokenPrice float64
	Err        string
}

func NewBalance(token *mongoServiceDB.Tokens, amount *big.Int, price float64, err error) *Balance {
	balance := &Balance{
		Tokens:     token,
		Amount:     amount,
		TokenPrice: price,
	}

	if err != nil {
		balance.Err = err.Error()
	}

	return balance
}

// Price returns amount * token price
func (b *Balance) Price() float64 {
	if b.Amount == nil {
		return 0
	}

	if b.Amount.Cmp(big.NewInt(0)) == 0 {
		return 0
	}
	price, _ := util.ToEtherWithDecimal(b.Amount, b.Tokens.Decimal).
		Mul(decimal.NewFromFloat(b.TokenPrice)).
		Round(6).
		Float64()
	return price
}

func (b *Balance) Error() string {
	return b.Err
}

func (p *Client) Balance(ctx context.Context, address, currencyId string) (*Balance, error) {
	tokenInfo, err := cache.TokenCache().GetTokensByCurrencyID(ctx, currencyId)
	if err != nil {
		return nil, err
	}

	if tokenInfo.Type == 1 {
		balance, err := p.BalanceAt(ctx, tokenInfo.Chain, address, nil)
		if err != nil {
			logger.Error("Balance", logger.BuildLogInput().WithError(fmt.Errorf("BalanceAt: %w", err)).WithChain(tokenInfo.Chain))
			balance = big.NewInt(0)
		}
		return NewBalance(
			tokenInfo,
			balance,
			cache.PriceCache().GetPriceBySymbolWithNoErr(tokenInfo.Symbol),
			err,
		), nil
	} else {
		switch types.GetChainType(tokenInfo.Chain) {
		// TODO: 체인 추가시 체크 필요
		case types.ChainTypeEVM:
			contractInfo, err := cache.ContractCache().GetContractByAddress(ctx, tokenInfo.Chain, tokenInfo.Address)
			if err != nil {
				logger.Error("Balance", logger.BuildLogInput().WithError(fmt.Errorf("GetContractByAddress: %w", err)).WithChain(tokenInfo.Chain))
				return nil, err
			}

			balanceOf := tokentypes.OutputBalanceOf{}
			err = p.CallMsgUnmarshalContract2(ctx, contractInfo, "balanceOf", &balanceOf, common.HexToAddress(tokenInfo.Chain, address))
			if err != nil {
				logger.Error("Balance", logger.BuildLogInput().WithError(fmt.Errorf("CallMsgUnmarshalContract2: %w", err)).WithChain(tokenInfo.Chain))
			}
			return NewBalance(
				tokenInfo,
				balanceOf.Amount,
				cache.PriceCache().GetPriceBySymbolWithNoErr(tokenInfo.Symbol),
				err,
			), nil
		case types.ChainTypeCOSMOS:
			client, err := cosmos.ProxyClient(p, tokenInfo.Chain)
			if err != nil {
				logger.Error("Balance", logger.BuildLogInput().WithError(fmt.Errorf("ProxyClient: %w", err)).WithChain(tokenInfo.Chain))
				return nil, err
			}

			switch tokenInfo.Type {
			case mongoServiceDB.TokenTypeCoin:
				balance, err := client.Balance(ctx, address, tokenInfo.Denom)
				if err != nil {
					return NewBalance(
						tokenInfo,
						big.NewInt(0),
						cache.PriceCache().GetPriceBySymbolWithNoErr(tokenInfo.Symbol),
						err,
					), nil
				}

				return NewBalance(
					tokenInfo,
					balance,
					cache.PriceCache().GetPriceBySymbolWithNoErr(tokenInfo.Symbol),
					err,
				), nil
			case mongoServiceDB.TokenTypeToken:
				// CW-20
				var queryBalanceResp cw20.QueryBalanceResponse
				err := client.CallWasm(ctx, tokenInfo.Address, cw20.QueryBalanceRequest{
					Address: address,
				}, &queryBalanceResp)
				if err != nil {
					logger.Error("BalanceAll", logger.BuildLogInput().WithError(fmt.Errorf("BalanceAt: %w", err)).WithChain(tokenInfo.Chain).WithAddress(address).WithData("contractAddress", tokenInfo.Address))
				}
				cw20Balance := util.ToDecimal(queryBalanceResp.Balance, 0)
				// i.e) NPT.fns -> NPT
				priceSymbol := strings.Split(tokenInfo.Symbol, ".")[0]
				return NewBalance(
					tokenInfo,
					cw20Balance.BigInt(),
					cache.PriceCache().GetPriceBySymbolWithNoErr(priceSymbol),
					err,
				), nil
			default:
				logger.Warn("Balance", logger.BuildLogInput().WithError(fmt.Errorf("unknown token type")).WithChain(tokenInfo.Chain).WithAddress(address).WithData("type", tokenInfo.Type, "contractAddress", tokenInfo.Address))
			}

		default:
			return nil, err
		}
	}
	return nil, nil
}

func (p *Client) Allowance(ctx context.Context, addressInfo account.AddressInfo, tokenInfo *mongoServiceDB.Tokens, spenderAddress string) (*big.Int, error) {
	var tokenAllowance *big.Int
	chain := tokenInfo.Chain
	ownerAddress := common.HexToAddress(chain, addressInfo.GetAddressByChain(chain))
	switch types.GetChainType(chain) {
	// TODO: 체인 추가시 체크 필요
	case types.ChainTypeEVM:
		contractInfo, err := cache.ContractCache().GetContractByAddress(ctx, chain, tokenInfo.Address)
		if err != nil {
			logger.Error("TokenAllowance", logger.BuildLogInput().WithError(fmt.Errorf("GetContractByAddress: %w", err)).WithChain(tokenInfo.Chain))
			return nil, err
		}

		msg, err := p.CallMsg(ctx, chain, "", tokenInfo.Address, "allowance", contractInfo.ABI, ownerAddress, common.HexToAddress(chain, spenderAddress))
		if err != nil {
			logger.Error("Allowance", logger.BuildLogInput().WithError(fmt.Errorf("CallMsg: %w", err)).WithChain(tokenInfo.Chain).WithData("tokenAddress", tokenInfo.Address, "spender", spenderAddress))
			return nil, err
		}
		var output tokentypes.OutputAllowance
		output.Unmarshal(msg)
		tokenAllowance = output.Amount
	case types.ChainTypeCOSMOS:
		client, err := cosmos.ProxyClient(p, chain)
		if err != nil {
			logger.Error("BalanceAll", logger.BuildLogInput().WithError(fmt.Errorf("ProxyClient: %w", err)).WithChain(tokenInfo.Chain).WithAddress(addressInfo.GetAddressByChain(chain)))
			return nil, err
		}
		var queryAllowanceResp cw20.QueryAllowanceResponse

		err = client.CallWasm(ctx, tokenInfo.Address, cw20.QueryAllowanceRequest{
			Owner:   ownerAddress.String(),
			Spender: spenderAddress,
		}, &queryAllowanceResp)
		if err != nil {
			return nil, err
		}

		tokenAllowance = util.ToDecimal(queryAllowanceResp.Allowance, tokenInfo.Decimal).BigInt()
	default:
		logger.Warn("TokenAllowance", logger.BuildLogInput().WithError(fmt.Errorf("ToDecimal:  not support chain type")).WithChain(tokenInfo.Chain))
		tokenAllowance = big.NewInt(0)
	}

	return tokenAllowance, nil
}
