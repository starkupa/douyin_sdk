package oauth

import (
	"context"
	"encoding/json"
	base "openapi/agent/web"
	"openapi/dependency"
	"openapi/pkg"
	"time"

	"openapi/dto/web"

	"github.com/spf13/cast"
)

func checkAuthQrCode(ctx context.Context, clientKey, token, toUri, state, scope string) (*web.TiktokCheckAuthCodeData, dependency.Catcher) {
	param := web.TiktokCheckAuthCodeRequest{
		ClientKey: clientKey,
		Scope:     scope,
		Next:      toUri,
		JumpType:  "native",
		Token:     token,
		State:     state,
		TimeStamp: cast.ToString(time.Now().UnixMicro()),
		CustomizeParams: web.CustomizeParams{
			Source:         "pc_auth",
			NotSkipConfirm: "true",
		},
	}
	reply, err := pkg.GET(ctx, base.CheckQrCodeURL, param.ToValues(), nil)
	if err != nil {
		return nil, dependency.ERROR_BUSI.Error(err)
	}
	if reply == nil {
		return nil, dependency.ERROR_BUSI.ErrorMsg("response nil")
	}
	res := &web.TiktokCheckAuthCodeResponse{}
	if err = json.Unmarshal(reply, res); err != nil {
		return nil, dependency.ERROR_BUSI.Error(err)
	}
	if err = res.Data.ErrorCode.ToErr(res.Data.Description); err != nil {
		return nil, dependency.ERROR_BUSI.Error(err)
	}
	return &res.Data, nil
}
