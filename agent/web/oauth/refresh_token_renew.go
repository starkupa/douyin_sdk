package oauth

import (
	"context"
	"encoding/json"
	base "openapi/agent/web"
	"openapi/dependency"
	"openapi/dto/web"
	"openapi/pkg"
)

// 文档：https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/account-permission/refresh-token

func renewRefreshToken(ctx context.Context, clientKey, refreshToken string) (*web.RenewRefreshTokenData, dependency.Catcher) {
	param := web.RenewRefreshTokenRequest{
		ClientKey:    clientKey,
		RefreshToken: refreshToken,
	}
	reply, err := pkg.PostForm(ctx, base.RenewRefreshTokenURL, param.ToValues(), nil)
	if err != nil {
		return nil, dependency.ERROR_AUTH.Error(err)
	}
	if reply == nil {
		return nil, dependency.ERROR_AUTH.ErrorMsg("response nil")
	}
	res := &web.RenewRefreshTokenResponse{}
	if err = json.Unmarshal(reply, res); err != nil {
		return nil, dependency.ERROR_AUTH.Error(err)
	}
	if err = res.Data.ErrorCode.ToErr(res.Data.Description); err != nil {
		return nil, dependency.ERROR_AUTH.Error(err)
	}
	return &res.Data, nil
}
