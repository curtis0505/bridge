package web

import (
	"fmt"
	"github.com/curtis0505/bridge/libs/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	Success                        = NewResult(0, "Success")                           // 성공
	Failed                         = NewResult(1, "Failed")                            // 실패
	IpInvalid                      = NewResult(2, "IpInvalid")                         // 요청한 아이피가 정상적이지 않습니다.
	UserIDNotFound                 = NewResult(13, "UserIDNotFound")                   // 유저 아이디가 존재하지 않습니다
	AccessTokenInvalid             = NewResult(101, "AccessTokenInvalid")              // 접속 토큰이 유효하지 않음
	UserNotFound                   = NewResult(102, "UserNotFound")                    // 유저 정보를 찾을 수 없음
	UserExistMnemonic              = NewResult(103, "UserExistMnemonic")               // 유저의 니모닉이 이미 있습니다.
	OTPEncryptionFail              = NewResult(104, "OTPEncryptionFail")               // OTP를 암호화 하는 도중 실패하였습니다.
	UserAddressNotFound            = NewResult(106, "UserAddressNotFound")             // 유저 어드레스(체인)이 아무것도 없을 경우 입니다.
	UserNotSleep                   = NewResult(107, "UserNotSleep")                    // 유저가 휴면 상태가 아닙니다.
	OTPDeleteFail                  = NewResult(109, "OTPDeleteFail")                   // OTP를 초기화하는데 실패하였습니다.
	OTPInvalid                     = NewResult(110, "OTPInvalid")                      // OTP가 유효하지 않습니다.
	KYCNotApprovedUser             = NewResult(112, "KYCNotApprovedUser")              // kyc 승인이 되지않았습니다.
	UserNotExist                   = NewResult(114, "UserNotExist")                    // 유저가 존재하지 않습니다.
	UserAlreadyExist               = NewResult(113, "UserAlreadyExist")                // 유저가 이미 존재합니다.
	AlreadyTransaction             = NewResult(309, "AlreadyTransaction")              // 이미 진해중인 트랜잭션이 존재한다.
	TxFail                         = NewResult(313, "TxFail")                          // 트렌잭션이 실패했습니다.
	Forbidden                      = NewResult(403, "Forbidden")                       // 금지된 국가
	Maintenance                    = NewResult(500, "Maintenance")                     // 전체점검 중
	TxInvalid                      = NewResult(701, "TxInvalid")                       // 트랜잭션이 올바르지 않습니다. (검증 실패)
	FunctionBlocked                = NewResult(707, "FunctionBlocked")                 // 기능이 막혀있어서 사용할 수 없습니다.
	InvalidMinVal                  = NewResult(802, "InvalidMinVal")                   // 최소 수량보다 작은 경우 발생한다.
	InvalidMaxVal                  = NewResult(803, "InvalidMinVal")                   // 최소 수량보다 많은 경우 발생한다.
	NotKYCApprovedFriendCode       = NewResult(900, "NotKYCApprovedFriendCode")        //
	LockTxChainErr                 = NewResult(1000, "LockTxChainErr")                 // 체인별로 트랜잭션 요청을 막기위한 코드
	NotFoundChain                  = NewResult(1001, "NotFoundChain")                  // 없는 체인
	SwapV2Error                    = NewResult(1100, "SwapV2Error")                    // SwapV2 Error
	SwapV2ApproveError             = NewResult(1101, "SwapV2ApproveError")             // SwapV2 Error
	BridgeError                    = NewResult(1200, "BridgeError")                    // Bridge Error
	DBError                        = NewResult(1300, "DBError")                        // DB Error
	InvalidStatusForSpeedupError   = NewResult(1400, "InvalidStatusForSpeedupError")   // speedup을 위한 status가 유효하지 않다.
	InvalidNonceForSpeedupError    = NewResult(1401, "InvalidNonceForSpeedupError")    // speedup을 위한 nonce가 유효하지 않다.
	InvalidGasPriceForSpeedupError = NewResult(1402, "InvalidGasPriceForSpeedupError") // speedup을 위한 gasPrice가 유효하지 않다.
	DetectAddress                  = NewResult(1500, "DetectAddress")                  // 이상한 어드레스 감지되었습니다.
	NotMatchAddress                = NewResult(1501, "NotMatchAddress")                // 서버의 주소와 앱에서 요청한 주소가 다른 경우 발생
	NeopinUserAddress              = NewResult(1502, "NeopinUserAddress")              // 주소 import 할 때 네오핀 유저 주소가 등록된 경우 발생
	InvalidWhiteList               = NewResult(2000, "InvalidWhiteList")               // 화이트리스트에 없는 주소
	InvalidApiKey                  = NewResult(2001, "InvalidApiKey")                  // api key 가 유효하지 않다.
	InvalidVaspCode                = NewResult(2002, "InvalidVaspCode")                // vasp code 가 유효하지 않다.
	BlacklistAddress               = NewResult(2003, "BlacklistAddress")               // 블랙리스트에 있는 주소
	InvalidAccessToken             = NewResult(4010, "InvalidAccessToken")             // access_token 이 유효하지 않다.
	ExpireAccessToken              = NewResult(4011, "ExpireAccessToken")              // access_token 이 만료되었다.
	NotLastRegisteredAccessToken   = NewResult(4012, "NotLastRegisteredAccessToken")   // 서버에 마지막에 등록된 access_token 이 아니다(중복 로그인)
	InvalidRefreshToken            = NewResult(4013, "InvalidRefreshToken")            // refresh_token 이 유효하지 않다.
	ExpiredRefreshToken            = NewResult(4014, "ExpiredRefreshToken")            // refresh_token 이 만료되었다.
	NotLastRegisteredRefreshToken  = NewResult(4015, "NotLastRegisteredRefreshToken")  // 서버에 마지막에 등록된 refresh_token 이 아니다(중복 로그인)
	PanicError                     = NewResult(5000, "PanicError")                     // 서버 패닉시 발생하는 공통 코드

)

func RespOk(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, data)
}

func RespError(ctx *gin.Context, status int, err error) {
	logger.Error("Request error", "path", ctx.FullPath(), "status", status, "error", err, "requestId", ctx.GetString("X-Request-Id"))
	ctx.AbortWithStatusJSON(status, NewRespHeader(Failed, err.Error()))
}

type Result struct {
	Desc string
	Code int
}

var resultCodeSpace = make(map[int]Result)

func NewResult(code int, desc string) Result {
	return registerResult(code, desc)
}

func registerResult(code int, desc string) Result {
	if result, ok := resultCodeSpace[code]; ok {
		panic(fmt.Sprintf("duplicated result code: %d, desc: %s, got: %s", result.Code, result.String(), desc))
	}

	result := Result{
		Code: code,
		Desc: desc,
	}

	resultCodeSpace[code] = result
	return result
}

func (result Result) String() string {
	return result.Desc
}

// RespHeader 모든 응답의 헤더
type RespHeader struct {
	Code    int    `json:"result"`
	Message string `json:"resultString"`
	Desc    string `json:"desc"`
}

// NewRespHeader : RespHeader 객체 생성 및 반환
func NewRespHeader(result Result, desc ...string) *RespHeader {
	return &RespHeader{
		Code:    result.Code,
		Message: result.String(),
		Desc:    strings.Join(desc, ","),
	}
}
