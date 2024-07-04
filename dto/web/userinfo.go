package web

import "net/url"

type UserInfoRequest struct {
	OpenID      string `json:"open_id"`      // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string `json:"access_token"` // 调用 /oauth/access_token/ 生成的 token
}

func (r *UserInfoRequest) ToValues() url.Values {
	data := url.Values{}
	data.Set("open_id", r.OpenID)
	data.Set("access_token", r.AccessToken)
	return data
}

type UserInfoResponse struct {
	Data    *UserInfoData `json:"data"`
	Message string        `json:"message"`
}

type UserInfoData struct {
	Avatar      string `json:"avatar"`   // 头像
	Nickname    string `json:"nickname"` // 昵称
	ErrorCode   int64  `json:"error_code"`
	Description string `json:"description"`
	SensitiveInformation
}

type SensitiveInformation struct {
	LogID        string `json:"log_id"`
	OpenID       string `json:"open_id"`
	UnionID      string `json:"union_id"`
	ClientKey    string `json:"client_key"`
	EAccountRole string `json:"e_account_role"`
}
