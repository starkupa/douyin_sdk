package oauth

import (
	"context"
	"encoding/json"
	base "openapi/agent/web"
	"openapi/dependency"
	"openapi/dto/web"
	"openapi/pkg"
)

func getAuthQrCode(ctx context.Context, clientKey, toUri, state, scope string) (*web.TiktokGetAuthCodeData, dependency.Catcher) {
	param := web.TiktokGetAuthCodeRequest{
		ClientKey: clientKey,
		Scope:     scope,
		Next:      toUri,
		JumpType:  "native",
		State:     state,
		CustomizeParams: web.CustomizeParams{
			Source:         "pc_auth",
			NotSkipConfirm: "true",
		},
	}
	reply, err := pkg.GET(ctx, base.GetQrCodeURL, param.ToValues(), nil)
	if err != nil {
		return nil, dependency.ERROR_BUSI.Error(err)
	}
	if reply == nil {
		return nil, dependency.ERROR_BUSI.ErrorMsg("response nil")
	}
	res := &web.TiktokGetAuthCodeResponse{}
	if err = json.Unmarshal(reply, res); err != nil {
		return nil, dependency.ERROR_BUSI.Error(err)
	}
	if err = res.Data.ErrorCode.ToErr(res.Data.Description); err != nil {
		return nil, dependency.ERROR_BUSI.Error(err)
	}
	return &res.Data, nil
}
