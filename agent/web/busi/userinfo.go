package busi

import (
	"context"
	"encoding/json"
	base "openapi/agent/web"
	"openapi/dependency"
	"openapi/dto/web"
	"openapi/pkg"
)

// 文档：https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/account-management/get-account-open-info

func userInfo(ctx context.Context, accessToken, openId string) (*web.UserInfoData, dependency.Catcher) {
	param := web.UserInfoRequest{
		OpenID:      openId,
		AccessToken: accessToken,
	}
	reply, err := pkg.PostForm(ctx, base.UserInfoURL, param.ToValues(), nil)
	if err != nil {
		return nil, dependency.ERROR_AUTH.Error(err)
	}
	if reply == nil {
		return nil, dependency.ERROR_AUTH.ErrorMsg("response nil")
	}
	res := &web.UserInfoResponse{}
	if err = json.Unmarshal(reply, res); err != nil {
		return nil, dependency.ERROR_AUTH.Error(err)
	}
	if res.Data.ErrorCode != 0 {
		return nil, dependency.ERROR_AUTH.ErrorMsg(res.Data.Description)
	}
	return res.Data, nil
}
