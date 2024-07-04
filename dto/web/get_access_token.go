package web

import "net/url"

type AccessTokenRequest struct {
	ClientKey    string `json:"client_key"`    // 应用唯一标识
	ClientSecret string `json:"client_secret"` // 应用唯一标识对应的密钥
	Code         string `json:"code"`          // 用户授权码
	GrantType    string `json:"grant_type"`    // 固定值"authorization_code"
}

func (r *AccessTokenRequest) ToValues() url.Values {
	data := url.Values{}
	data.Set("client_key", r.ClientKey)
	data.Set("client_secret", r.ClientSecret)
	data.Set("code", r.Code)
	data.Set("grant_type", "authorization_code")
	return data
}

type AccessTokenResponse struct {
	Data    AccessTokenData `json:"data"`
	Message string          `json:"message"`
}

type AccessTokenData struct {
	AccessTokenCoreData
	Description string  `json:"description"`
	ErrorCode   ErrCode `json:"error_code"`
	LogID       string  `json:"log_id"`
}

type AccessTokenCoreData struct {
	AccessToken      string `json:"access_token"`       // 接口调用凭证
	ExpiresIn        int64  `json:"expires_in"`         // access_token接口调用凭证超时时间，单位（秒)
	OpenID           string `json:"open_id"`            // 授权用户唯一标识
	RefreshExpiresIn int64  `json:"refresh_expires_in"` // refresh_token凭证超时时间，单位（秒)
	RefreshToken     string `json:"refresh_token"`      // 用户刷新access_token
	Scope            string `json:"scope"`              // 用户授权的作用域(Scope)，使用逗号（,）分隔，开放平台几乎每个接口都需要特定的Scope。
}
