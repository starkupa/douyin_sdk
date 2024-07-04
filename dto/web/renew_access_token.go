package web

import "net/url"

type RefreshAccessTokenRequest struct {
	ClientKey    string `json:"client_key"`    // 应用唯一标识
	RefreshToken string `json:"refresh_token"` // 填写通过 access_token 获取到的 refresh_token 参数
	GrantType    string `json:"grant_type"`    // 固定值“refresh_token”
}

func (r *RefreshAccessTokenRequest) ToValues() url.Values {
	data := url.Values{}
	data.Set("client_key", r.ClientKey)
	data.Set("refresh_token", r.RefreshToken)
	data.Set("grant_type", "refresh_token")
	return data
}

type RefreshAccessTokenResponse struct {
	Data    RefreshAccessTokenData `json:"data"`
	Message string                 `json:"message"`
}

type RefreshAccessTokenData struct {
	AccessTokenData
}

func (a AccessTokenData) SetRefreshAccessTokenData() *RefreshAccessTokenData {
	return &RefreshAccessTokenData{
		AccessTokenData: a,
	}
}
