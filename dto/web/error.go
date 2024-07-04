package web

import "fmt"

type ErrCode int64

func (e ErrCode) ToInt64() int64 {
	return int64(e)
}

func (e ErrCode) ToErr(msg ...string) error {
	if e == 0 {
		return nil
	}
	v, ok := ErrorMap[e]
	if ok {
		return fmt.Errorf("%v, %v", v, msg)
	}
	return fmt.Errorf("%v", msg)
}

const (
	SystemError     ErrCode = 10001 // 系统异常
	BadRequestError ErrCode = 10002 // 参数错误
	NoPrivilege     ErrCode = 10004 // 没有权限
	MissParam       ErrCode = 10005 // 缺少参数
	CodeExpired     ErrCode = 10007 // 授权码过期
	RefreshExpired  ErrCode = 10010 // refresh_token过期
	ClientInvalid   ErrCode = 10013 // 应用信息无效
	ClientKeyMissed ErrCode = 10014 // client_key不匹配
	RefreshTooMany  ErrCode = 10020 // refresh_token刷新次数过多
)

var ErrorMap = map[ErrCode]error{
	SystemError:     fmt.Errorf("授权系统异常"),
	BadRequestError: fmt.Errorf("参数错误"),
	NoPrivilege:     fmt.Errorf("没有权限"),
	MissParam:       fmt.Errorf("缺少参数"),
	CodeExpired:     fmt.Errorf("授权码过期，请重新授权"),
	RefreshExpired:  fmt.Errorf("token已过期"),
	ClientInvalid:   fmt.Errorf("应用信息无效"),
	ClientKeyMissed: fmt.Errorf("token 已过期，请重新授权"),
	RefreshTooMany:  fmt.Errorf("无法再刷新，请重新授权"),
}

const (
	BaseErrCode = -1
)
