package types

import (
	"fmt"
)

type TokenType int8

var (
	NotImplemented = fmt.Errorf("not implemented")
)

const (
	TokenTypeNotToken TokenType = iota
	TokenTypeERC20
	TokenTypeERC721
)

// SendTxResultType : reqtx 결과 타입
type SendTxResultType uint8

const (
	SendTxResultType_Success SendTxResultType = 1
	//요청 성공
	SendTxResultType_Reverted = 2 //revert가 됨.
	SendTxResultType_Timeout  = 3 //블럭에 포함되지까지 시간이 타임아웃됨.

	SendTxResultType_NotOpenYet             = 11 //아직 정상적인 로딩이 끝나지 않음.
	SendTxResultType_SynchronizingWithChain = 12 //현재 체인과 동기화중임.

	SendTxResultType_DecodeRawTxError  = 21 //tx를 decode하는데 실패함.
	SendTxResultType_TxTypeError       = 22 //tx타입이 TxTypeFeeDelegatedSmartContractExecution가 아님.
	SendTxResultType_UnknownChainID    = 23 //알수없는 체인 id
	SendTxResultType_InvalidSender     = 24 //tx의 from이 제대로 설정되어있지않음.
	SendTxResultType_FeepayerSignError = 25 //fee payer의 서명이 실패.
	SendTxResultType_FeepayerLimit     = 26 //fee payer의 횟수 제한.
	SendTxResultType_InvalidGasPrice   = 27 //gas price 가 올바르지 않는 경우(ex. gas price 가 0인 경우)

	SendTxResultType_NotMember       = 31 //맴버가 아님.
	SendTxResultType_BlackListMember = 32 //블릭리스트임.
	SendTxResultType_NotDeployer     = 33 //컨트랙트 배포자가 아님.

	SendTxResultType_SameTxPending      = 41 //동일한 tx가 대기중임.
	SendTxResultType_SameAddressPending = 42 //동일한 address의 tx가 현재 대기중임.
	SendTxResultType_ExceededTxDaily    = 43 //하루 사용가능한 tx량 초과
	SendTxResultType_TxDataFiltered     = 45 //tx data의 method단위로 tx발생이 불가능한경우

	SendTxResultType_SendTxError = 51 //노드에 sendrawTx실패
)

func ResultToString(result SendTxResultType) string {
	switch result {
	case SendTxResultType_Success:
		return "SendTxResultType_Success"
	case SendTxResultType_Reverted:
		return "SendTxResultType_Reverted"
	case SendTxResultType_Timeout:
		return "SendTxResultType_Timeout"
	case SendTxResultType_NotOpenYet:
		return "SendTxResultType_NotOpenYet"
	case SendTxResultType_SynchronizingWithChain:
		return "SendTxResultType_SynchronizingWithChain"
	case SendTxResultType_DecodeRawTxError:
		return "SendTxResultType_DecodeRawTxError"
	case SendTxResultType_TxTypeError:
		return "SendTxResultType_TxTypeError"
	case SendTxResultType_UnknownChainID:
		return "SendTxResultType_UnknownChainID"
	case SendTxResultType_InvalidSender:
		return "SendTxResultType_InvalidSender"
	case SendTxResultType_FeepayerSignError:
		return "SendTxResultType_FeepayerSignError"
	case SendTxResultType_FeepayerLimit:
		return "SendTxResultType_FeepayerLimit"
	case SendTxResultType_InvalidGasPrice:
		return "SendTxResultType_InvalidGasPrice"
	case SendTxResultType_NotMember:
		return "SendTxResultType_NotMember"
	case SendTxResultType_BlackListMember:
		return "SendTxResultType_BlackListMember"
	case SendTxResultType_NotDeployer:
		return "SendTxResultType_NotDeployer"
	case SendTxResultType_SameTxPending:
		return "SendTxResultType_SameTxPending"
	case SendTxResultType_SameAddressPending:
		return "SendTxResultType_SameAddressPending"
	case SendTxResultType_ExceededTxDaily:
		return "SendTxResultType_ExceededTxDaily"
	case SendTxResultType_TxDataFiltered:
		return "SendTxResultType_TxDataFiltered"
	case SendTxResultType_SendTxError:
		return "SendTxResultType_SendTxError"
	}
	return "unknown result type"
}
