package web

import "net/url"

type RenewRefreshTokenRequest struct {
	ClientKey    string `json:"client_key"`    // 应用唯一标识
	RefreshToken string `json:"refresh_token"` // 填写通过/oauth/access_token/ 获取到的 refresh_token 参数
}

func (r *RenewRefreshTokenRequest) ToValues() url.Values {
	data := url.Values{}
	data.Set("client_key", r.ClientKey)
	data.Set("refresh_token", r.RefreshToken)
	return data
}

type RenewRefreshTokenResponse struct {
	Data    RenewRefreshTokenData `json:"data"`
	Message string                `json:"message"`
}

type RenewRefreshTokenData struct {
	ErrorCode    ErrCode `json:"error_code"`
	RefreshToken string  `json:"refresh_token"`
	Description  string  `json:"description"`
	ExpiresIn    int64   `json:"expires_in"`
}
