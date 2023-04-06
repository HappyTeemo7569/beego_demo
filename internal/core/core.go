package core

// 全局错误代码
type ErrorCode int

const (
	ErrorCodeServer          = 500 //服务器异常
	ErrorCode_Auth           = 600 //鉴权失败
	ErrorCodeWallet          = 700 //扣费失败
	ErrorCodeWalletNotEnough = 701 //余额不足
)

type Error struct {
	Code ErrorCode   `json:"errorCode,omitempty"`
	Text string      `json:"txtMessage"`
	Data interface{} `json:"errorData,omitempty"`
}

func NewError(text string) *Error {
	return &Error{
		Code: ErrorCodeServer,
		Text: text,
		Data: nil,
	}
}

func (e *Error) AppendCode(code ErrorCode) *Error {
	e.Code = code
	return e
}

func (e *Error) AppendData(data interface{}) *Error {
	e.Data = data
	return e
}
