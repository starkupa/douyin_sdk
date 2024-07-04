package web

import (
	"encoding/json"
	"net/url"
)

// 获取授权码
type TiktokGetAuthCodeRequest struct {
	ClientKey            string          `json:"client_key"`
	Scope                string          `json:"scope"`     // 权限
	Next                 string          `json:"next"`      // 前端跳转地址
	State                string          `json:"state"`     // 用于保持请求和回调的状态
	JumpType             string          `json:"jump_type"` // 固定 native
	OptionalScopeCheck   string          `json:"optional_scope_check"`
	OptionalScopeUncheck string          `json:"optional_scope_uncheck"`
	CustomizeParams      CustomizeParams `json:"customize_params"`
}

func (r *TiktokGetAuthCodeRequest) ToValues() url.Values {
	data := url.Values{}
	data.Set("client_key", r.ClientKey)
	data.Set("scope", r.Scope)
	data.Set("next", r.Next)
	data.Set("jump_type", r.JumpType)
	data.Set("state", r.State)
	by, _ := json.Marshal(r.CustomizeParams)
	data.Set("customize_params", string(by))
	return data
}

type CustomizeParams struct {
	CommentID      string `json:"comment_id"`
	Source         string `json:"source"`           // use "pc_auth"
	NotSkipConfirm string `json:"not_skip_confirm"` // "true"
	EnterFrom      string `json:"enter_from"`
}

type TiktokGetAuthCodeResponse struct {
	Data    TiktokGetAuthCodeData `json:"data"`
	Message string                `json:"message"`
}

type TiktokGetAuthCodeData struct {
	Captcha        string  `json:"captcha"`
	DescUrl        string  `json:"desc_url"`
	Description    string  `json:"description"`
	ErrorCode      ErrCode `json:"error_code"`
	IsFrontier     bool    `json:"is_frontier"`
	Qrcode         string  `json:"qrcode"`
	QrcodeIndexUrl string  `json:"qrcode_index_url"`
	Token          string  `json:"token"`
}
