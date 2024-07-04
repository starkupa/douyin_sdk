package oauth

import (
	"context"
	"encoding/json"
	base "openapi/agent/web"
	"openapi/dependency"
	"openapi/dto/web"
	"openapi/pkg"
)

func getAccessToken(ctx context.Context, clientKey, clientSecret, code string) (*web.AccessTokenData, dependency.Catcher) {
	param := web.AccessTokenRequest{
		ClientKey:    clientKey,
		ClientSecret: clientSecret,
		Code:         code,
		GrantType:    "authorization_code",
	}
	reply, err := pkg.PostForm(ctx, base.AccessTokenURL, param.ToValues(), nil)
	if err != nil {
		return nil, dependency.ERROR_BUSI.Error(err)
	}
	if reply == nil {
		return nil, dependency.ERROR_BUSI.ErrorMsg("response nil")
	}
	res := &web.AccessTokenResponse{}
	if err = json.Unmarshal(reply, res); err != nil {
		return nil, dependency.ERROR_BUSI.Error(err)
	}
	if err = res.Data.ErrorCode.ToErr(res.Data.Description); err != nil {
		return nil, dependency.ERROR_AUTH.Error(err)
	}
	return &res.Data, nil
}
