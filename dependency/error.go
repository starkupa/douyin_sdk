package dependency

import "fmt"

type IErrorType int32

const (
	ERROR_BUSI             IErrorType = 10000
	IErrorType_SystemError IErrorType = 10001 // 系统错误
	IErrorType_Expired     IErrorType = 10002 // 授权过期
	ERROR_AUTH             IErrorType = 10003 // 未授权
)

func (e IErrorType) Error(err error) *IError {
	return &IError{
		ICode: int32(e),
		error: err,
	}
}

func (e IErrorType) ErrorMsg(msg string) *IError {
	return &IError{
		ICode: int32(e),
		error: fmt.Errorf(msg),
	}
}

type IError struct {
	ICode int32
	error
}

func (e IError) Code() int32 {
	return e.ICode
}

func (e IError) Error() error {
	return e.error
}

type Catcher interface {
	Code() int32
	Error() error
}
