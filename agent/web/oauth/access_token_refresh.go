package oauth

import (
	"context"
	"encoding/json"
	base "openapi/agent/web"
	"openapi/dependency"
	"openapi/dto/web"
	"openapi/pkg"
)

func refreshAccessToken(ctx context.Context, clientKey, refreshToken string) (*web.RefreshAccessTokenData, dependency.Catcher) {
	param := web.RefreshAccessTokenRequest{
		ClientKey:    clientKey,
		RefreshToken: refreshToken,
		GrantType:    "refresh_token",
	}
	reply, err := pkg.PostForm(ctx, base.RefreshAccessTokenURL, param.ToValues(), nil)
	if err != nil {
		return nil, dependency.ERROR_BUSI.Error(err)
	}
	if reply == nil {
		return nil, dependency.ERROR_BUSI.ErrorMsg("response nil")
	}
	res := &web.RefreshAccessTokenResponse{}
	if err = json.Unmarshal(reply, res); err != nil {
		return nil, dependency.ERROR_BUSI.Error(err)
	}
	if err = res.Data.ErrorCode.ToErr(res.Data.Description); err != nil {
		return nil, dependency.ERROR_AUTH.Error(err)
	}
	return &res.Data, nil
}
