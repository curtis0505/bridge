package types

type ResponseHeader struct {
	ResultCode ResponseCode `json:"result"`
	Error      string       `json:"error,omitempty"`
}

func NewResponseHeader(code ResponseCode, err error) *ResponseHeader {
	header := ResponseHeader{
		ResultCode: code,
	}
	if err != nil {
		header.Error = err.Error()
	}
	return &header
}

func NewResponseSuccess() *ResponseHeader {
	return NewResponseHeader(Success, nil)
}

type ResponseCode int

const (
	Success ResponseCode = iota
	Failed
)
