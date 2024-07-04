package busi

import (
	"context"
	"encoding/json"
	base "openapi/agent/web"
	"openapi/dependency"
	"openapi/dto/web"
	"openapi/pkg"
)

// 文档 https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/video-management/douyin/search-video/account-video-list

func videoList(ctx context.Context, accessToken, openId string, cursor, count int64) (*web.VideoListData, dependency.Catcher) {
	param := web.VideoListRequest{
		OpenID:      openId,
		AccessToken: accessToken,
		Cursor:      cursor,
		Count:       count,
	}
	header := map[string]string{
		"Access-Token": accessToken,
		"Content-Type": "application/json",
	}
	reply, err := pkg.GET(ctx, base.VideoListURL, param.ToValues(), header)
	if err != nil {
		return nil, dependency.ERROR_BUSI.Error(err)
	}
	if reply == nil {
		return nil, dependency.ERROR_BUSI.ErrorMsg("response nil")
	}
	res := &web.VideoListResponse{}
	if err = json.Unmarshal(reply, res); err != nil {
		return nil, dependency.ERROR_BUSI.Error(err)
	}
	if res.Data.ErrorCode != 0 {
		return nil, dependency.ERROR_BUSI.ErrorMsg(res.Data.Description)
	}
	return &res.Data, nil
}
