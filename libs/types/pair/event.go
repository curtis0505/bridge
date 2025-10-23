package pair

import (
	"github.com/kaiachain/kaia/common"
	"math/big"
)

// EventCreatePairByUser (address indexed user, address tokenA, address tokenB, uint256 amountA, uint256 amountB, uint256 liquidity, address protocolFeeTo, uint256 protocolFee, address protocolFeeAddr);
type EventCreatePairByUser struct {
	User            common.Address
	TokenA          common.Address `abi:"tokenA"`
	TokenB          common.Address `abi:"tokenB"`
	AmountA         *big.Int       `abi:"amountA"`
	AmountB         *big.Int       `abi:"amountB"`
	Liquidity       *big.Int       `abi:"liquidity"`
	ProtocolFeeTo   common.Address `abi:"protocolFeeTo"`
	ProtocolFee     *big.Int       `abi:"protocolFee"`
	ProtocolFeeAddr common.Address `abi:"protocolFeeAddr"`
}

// EventCreatePairETHByUser (address indexed user, address token, uint256 amountToken, uint256 amountETH, uint256 liquidity, address protocolFeeTo, uint256 protocolFee, address protocolFeeAddr);
type EventCreatePairETHByUser struct {
	User            common.Address
	Token           common.Address `abi:"token"`
	AmountToken     *big.Int       `abi:"amountToken"`
	AmountETH       *big.Int       `abi:"amountETH"`
	Liquidity       *big.Int       `abi:"liquidity"`
	ProtocolFeeTo   common.Address `abi:"protocolFeeTo"`
	ProtocolFee     *big.Int       `abi:"protocolFee"`
	ProtocolFeeAddr common.Address `abi:"protocolFeeAddr"`
}
