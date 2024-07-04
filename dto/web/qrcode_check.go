package web

import (
	"encoding/json"
	"net/url"
)

// check二维码
type TiktokCheckAuthCodeRequest struct {
	ClientKey            string          `json:"client_key"`
	Scope                string          `json:"scope"`
	Next                 string          `json:"next"`
	JumpType             string          `json:"jump_type"`
	Token                string          `json:"token"`
	State                string          `json:"state"`
	TimeStamp            string          `json:"time_stamp"`
	OptionalScopeCheck   string          `json:"optional_scope_check"`
	OptionalScopeUncheck string          `json:"optional_scope_uncheck"`
	CustomizeParams      CustomizeParams `json:"customize_params"`
}

func (r *TiktokCheckAuthCodeRequest) ToValues() url.Values {
	data := url.Values{}
	data.Set("client_key", r.ClientKey)
	data.Set("scope", r.Scope)
	data.Set("next", r.Next)
	data.Set("jump_type", r.JumpType)
	data.Set("token", r.Token)
	data.Set("state", r.State)
	data.Set("time_stamp", r.TimeStamp)
	by, _ := json.Marshal(r.CustomizeParams)
	data.Set("customize_params", string(by))
	return data
}

type TiktokCheckAuthCodeResponse struct {
	Data    TiktokCheckAuthCodeData `json:"data"`
	Message string                  `json:"message"`
}

type TiktokCheckAuthCodeData struct {
	Captcha         string  `json:"captcha"`
	Code            string  `json:"code"`
	ConfirmedScopes string  `json:"confirmed_scopes"`
	DescUrl         string  `json:"desc_url"`
	Description     string  `json:"description"`
	ErrorCode       ErrCode `json:"error_code"`
	RedirectUrl     string  `json:"redirect_url"`
	State           string  `json:"state"`
	Status          string  `json:"status"`
	OpenID          string  `json:"open_id"`
}
