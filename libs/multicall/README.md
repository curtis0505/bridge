# MultiCall


### Using MultiCall 

```go
// 초기화 해야 함
multicall.Init(client, p.cache)

multiCall := multicall.New()

convertToNpKlay := stakingtypes.OutputConvertKlayToNpKlay{}
getTotalPoolKlay := stakingtypes.OutputGetTotalPooledKlay{}
getTotalStake := stakingtypes.OutputGetTotalStake{}
getProtocolFeeBP := stakingtypes.OutputGetProtocolFeeBP{}

multiCall.AddCallUnmarshaler(itemInfo.Chain, contractInfo.Address, "convertKlayToNpKlay", contractInfo.Abi, &convertToNpKlay, big.NewInt(1e18))
multiCall.AddCallUnmarshaler(itemInfo.Chain, contractInfo.Address, "getTotalPooledKlay", contractInfo.Abi, &getTotalPoolKlay)
multiCall.AddCallUnmarshaler(itemInfo.Chain, contractInfo.Address, "getTotalStake", contractInfo.Abi, &getTotalStake)
multiCall.AddCallUnmarshaler(itemInfo.Chain, contractInfo.Address, "protocolFeeBP", contractInfo.Abi, &getProtocolFeeBP)

err = p.proxy.CallMsgUnmarshal(
	ctx, 
	itemInfo.Chain,                     //chain
	"",                                 // from address
	multicall.Address[types.ChainKLAY], // to address
	multicall.Aggregate,                // method : aggregate
	multicall.Abi,                      // multiCall abi
	multiCall,                          // multiCall is Unmarshaler
	multiCall,                          // multiCall is Data
)

// or 
err = p.proxy.CallMsgMultiCall(
ctx,
chain,
multiCall,
)


```

### Batch

- 내부적으로 MultiCall을 Chunk로 처리하도록 함.
```go
// 초기화 해야 함
multicall.Init(client, p.cache)

multiCall := multicall.New()
size := 10000

balanceOf := make([]tokentypes.OutputBalanceOf{}, size)

for i := 0; i < 10000; i++ {
    multiCall.AddCallUnmarshaler(chain, address, "balanceOf", Abi, &balanceOf[i], common.HexToAddress("0x..."))
}

err = p.proxy.CallMsgMultiCall(
	ctx,
	chain,
	multiCall,
)

```