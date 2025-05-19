package restaking

const (
	MethodDepositRestakeToken  = "depositRestakeToken"
	MethodWithdrawRestakeToken = "withdrawRestakeToken"
)

const (
	EventNameDepositRestakeToken  = "DepositRestakeToken"
	EventNameWithdrawRestakeToken = "WithdrawRestakeToken"
)

const (
	AbiDepositRestakeToken  = "depositRestakeToken(address,string,bytes,bytes)"
	AbiWithdrawRestakeToken = "withdrawRestakeToken(string,bytes,address,address,bytes32[],uint256[],bool)"
)
