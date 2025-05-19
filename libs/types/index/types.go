package index

import (
	"encoding/json"
	"github.com/curtis0505/bridge/libs/common"
	ethercommon "github.com/ethereum/go-ethereum/common"
	"math/big"
)

const (
	Index = "INDEX"
)

const (
	IndexRWA      = "INDEX-rwa"
	IndexEthereum = "INDEX-ethereum"
	IndexMeme     = "INDEX-meme"

	ExchangeIssuanceID   = "exchangeissuance"
	StreamingFeeModuleID = "streamingfeemodule"
	UniswapQuoterID      = "quoteruni"
	UniswapV2FactoryID   = "uniswapv2factoryuni"
	UniswapV3FactoryID   = "uniswapv3factoryuni"
)

const (
	Maintenance = "index"
)

const (
	None = uint8(iota)
	UniV2
	UniV3
)

type SwapData struct {
	Path     [][20]byte `abi:"path" json:"path"`
	Fees     []*big.Int `abi:"fees" json:"fees"`
	Exchange uint8      `abi:"exchange" json:"exchange"`
}

func NewSwapData(exchange uint8, Fees []*big.Int, paths []common.Address) SwapData {
	swapData := SwapData{
		Fees:     Fees,
		Exchange: exchange,
	}
	swapData.SetPath(paths...)

	return swapData
}

func (s *SwapData) SetPath(addresses ...common.Address) {
	for _, address := range addresses {
		s.Path = append(s.Path, ethercommon.BytesToAddress(address.Bytes()))
	}
}

func (s *SwapData) GetPath() []common.Address {
	var paths []common.Address
	for _, path := range s.Path {
		paths = append(paths, ethercommon.BytesToAddress(path[:]))
	}
	return paths
}

func (s *SwapData) MarshalJSON() ([]byte, error) {
	swapData := struct {
		Path     []string   `json:"path"`
		Fees     []*big.Int `json:"fees"`
		Exchange uint8      `json:"exchange"`
	}{}

	for _, path := range s.Path {
		swapData.Path = append(swapData.Path, ethercommon.BytesToAddress(path[:]).String())
	}

	swapData.Fees = s.Fees
	swapData.Exchange = s.Exchange

	return json.Marshal(swapData)
}

func (s *SwapData) UnmarshalJSON(bz []byte) error {
	swapData := struct {
		Path     []string   `json:"path"`
		Fees     []*big.Int `json:"fees"`
		Exchange uint8      `json:"exchange"`
	}{}

	err := json.Unmarshal(bz, &swapData)
	if err != nil {
		return err
	}
	for _, path := range swapData.Path {
		s.Path = append(s.Path, common.HexToBytes(path))
	}
	s.Fees = swapData.Fees
	s.Exchange = swapData.Exchange
	return nil
}

type Addresses struct {
	UniV2Router common.Address
	UniV3Router common.Address
	UniV3Quoter common.Address
	WETH        common.Address
}
