package api

import (
	"fmt"
	"github.com/curtis0505/bridge/apps/validators/validator"
	"github.com/curtis0505/bridge/libs/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type ResultCode int

const (
	Success                    ResultCode = 0
	Failed                                = 1     // 요청이 실패하였습니다.
	InvalidTransaction                    = 10000 // 트랜잭션이 올바르지 않을때
	PendingTransaction                    = 10001 // 트랜잭션이 펜딩중일 때
	TransactionAlreadyExecuted            = 10002 // 이미 처리한 트랜잭션 일때
	PanicError                            = 50000 // 서버 패닉시 발생하는 공통 코드
)

func (r ResultCode) toString() string {
	switch r {
	case Success:
		return "Success"
	case Failed:
		return "Failed"
	case InvalidTransaction:
		return "InvalidTransaction"
	case PendingTransaction:
		return "PendingTransaction"
	case TransactionAlreadyExecuted:
		return "TransactionAlreadyExecuted"
	case PanicError:
		return "InternalServerError"
	}
	return ""
}

// BaseResponse 모든 응답의 헤더
type BaseResponse struct {
	Result       ResultCode `json:"result"`
	ResultString string     `json:"resultString,omitempty"`
	Desc         string     `json:"desc,omitempty"`
}

// NewBaseResponse : BaseResponse 객체 생성 및 반환
func NewBaseResponse(resultCode ResultCode, desc ...string) *BaseResponse {
	responseBase := BaseResponse{
		Result: resultCode,
	}

	// TODO prod 가 아닐때 아래 데이터가 추가되고, prod 이면 result code 만 표시
	if false {
		responseBase.ResultString = resultCode.toString()
		responseBase.Desc = strings.Join(desc, ",")
	}

	return &responseBase
}

type ValidatorAddressByChainResponse struct {
	*BaseResponse
	Address string `json:"address"`
}

type ValidatorInfo struct {
	ChainSymbol string `json:"chainSymbol"`
	Address     string `json:"address"`
}

type ValidatorAddressResponse struct {
	*BaseResponse
	Validator []ValidatorInfo `json:"validator"`
}

type CacheTransactionResponse struct {
	*BaseResponse
	TransactionList map[string]*validator.PublicTransactionHistory `json:"transactionList"`
}

type GenerateKeyResponse struct {
	*BaseResponse
	Keystore string `json:"keystore"`
}

func Response200(c *gin.Context, responseData interface{}) {
	c.JSON(http.StatusOK, responseData)
}

func Response422(c *gin.Context, err error) {
	logger.Warn(
		fmt.Sprintf("[%s] %s", c.Request.Method, c.FullPath()),
		"Header", c.Request.Header,
		"Url", c.Request.URL,
		//"Body", string(jsonData),
		"err", err,
	)
	c.JSON(http.StatusUnprocessableEntity, "Unprocessable Content")
	c.Abort()
}

func Response400(c *gin.Context) {
	c.String(http.StatusNotFound, "400 Bad Request")
	c.Abort()
}

func Response404(c *gin.Context) {
	c.String(http.StatusNotFound, "404 page not found")
	c.Abort()
}

func ResponseException(c *gin.Context, resultCode ResultCode, err error) {
	logger.Warn(
		fmt.Sprintf("[%s] %s", c.Request.Method, c.FullPath()),
		"Header", c.Request.Header,
		"Url", c.Request.URL,
		//"Body", string(jsonData),  // TODO 사용자가 입력한 body 로그 찍기
		"resultCode", resultCode,
		"resultCode.toString()", resultCode.toString(),
		"err", err,
	)

	errMessage := ""
	if err != nil {
		errMessage = err.Error()
	}

	c.JSON(http.StatusOK, NewBaseResponse(resultCode, errMessage))
	c.Abort()
}
