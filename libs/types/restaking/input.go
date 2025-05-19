package restaking

import (
	"github.com/curtis0505/bridge/libs/common"
)

// DepositRestakeTokenInput function depositRestakeToken(address _restakeTokenAddr, string memory _toChainName, bytes memory _to, bytes memory _proof)
type DepositRestakeTokenInput struct {
	RestakeTokenAddr common.Address `abi:"_restakeTokenAddr"`
	ToChainName      string         `abi:"_toChainName"`
	To               []byte         `abi:"_to"`
	Proof            []byte         `abi:"_proof"`
}

// WithdrawRestakeTokenInput function withdrawRestakeToken(string memory fromChainName, bytes memory from, address toAddr, address tokenAddr, bytes32[] memory txInfo, uint256[] memory tokenInfo, bool isDirect)
type WithdrawRestakeTokenInput struct {
	FromChainName string         `abu:"fromChainName"`
	From          []byte         `abu:"from"`
	ToAddr        common.Address `abu:"toAddr"`
	TokenAddr     common.Address `abu:"tokenAddr"`
	TxInfo        []byte         `abu:"txInfo"`
	TokenInfo     []byte         `abu:"tokenInfo"`
	IsDirect      bool           `abu:"isDirect"`
}
