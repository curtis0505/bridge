package types

type Chain string

const (
	// EVM Chains
	ChainKLAYLegacy = "KLY" // address_info 에서 필드로 사용중
	ChainKLAY       = "KLAY"
	ChainKAIA       = "KAIA"
	ChainMATIC      = "MATIC"
	ChainETH        = "ETH"
	ChainARB        = "ARB"
	ChainBASE       = "BASE"

	// TVM Chains
	ChainTRX = "TRX"

	// COSMOS Chains
	ChainATOM  = "ATOM"
	ChainFNSA  = "FNSA"
	ChainTFNSA = "TFNSA"
	ChainKAVA  = "KAVA"
	ChainOSMO  = "OSMO"
)

var (
	SupportedChains = []Chain{
		ChainKLAY,
		ChainMATIC,
		ChainETH,
		ChainTRX,
		ChainATOM,
		ChainFNSA,
		ChainTFNSA,
		ChainARB,
		ChainBASE,
	}
)

func (c Chain) String() string {
	return string(c)
}

var (
	BlockPerYear = map[string]int64{
		ChainKLAY:  86400 * 365 * 100,  // 1sec
		ChainMATIC: 43200 * 365 * 100,  // 2sec
		ChainETH:   7200 * 365 * 100,   // 12sec
		ChainTRX:   28800 * 365 * 100,  // 3sec
		ChainFNSA:  28800 * 365 * 100,  // 3sec
		ChainARB:   345600 * 365 * 100, // 0.25sec
		ChainBASE:  43200 * 365 * 100,  // 2sec
	}
	//(60 / BlockTime) * 60 * 24 * 365

	BlockTime = map[string]int64{
		ChainKLAY:  1,  //1sec
		ChainMATIC: 2,  //2sec
		ChainETH:   12, //12sec
		ChainTRX:   3,  //3sec
		ChainFNSA:  3,  //3sec
		//ChainARB:   0.25,
		ChainBASE: 2,
	}
)

type ChainType string

const (
	ChainTypeUnknown = ChainType("UNKNOWN")
	ChainTypeEVM     = ChainType("EVM")
	ChainTypeCOSMOS  = ChainType("ATOM")
	ChainTypeTVM     = ChainType("TVM")
)

func GetChainType(chain string) ChainType {
	if isEVM(chain) {
		return ChainTypeEVM
	}

	if isCOSMOS(chain) {
		return ChainTypeCOSMOS
	}

	if chain == ChainTRX {
		return ChainTypeTVM
	}

	return ChainTypeUnknown
}

func isEVM(chain string) bool {
	return chain == ChainKLAY || chain == ChainETH || chain == ChainMATIC || chain == ChainARB || chain == ChainBASE
}

func isCOSMOS(chain string) bool {
	return chain == ChainATOM || chain == ChainKAVA || chain == ChainOSMO ||
		chain == ChainFNSA || chain == ChainTFNSA
}
