package busi

import (
	"context"
	"encoding/json"
	base "openapi/agent/web"
	"openapi/dependency"
	"openapi/dto/web"
	"openapi/pkg"
)

// 文档：https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/video-management/douyin/search-video/video-data

func videoData(ctx context.Context, accessToken, openId string, videoIds []string, itemIds ...string) ([]web.VideoItem, dependency.Catcher) {
	if len(videoIds) == 0 && len(itemIds) == 0 {
		return nil, nil
	}
	param := web.VideoDataRequest{
		OpenID:      openId,
		AccessToken: accessToken,
		VideoIDs:    videoIds,
		ItemIDs:     itemIds,
	}
	if len(videoIds) > 0 && len(itemIds) > 0 {
		param.ItemIDs = nil
	}
	header := map[string]string{
		"Access-Token": accessToken,
	}
	reply, err := pkg.PostJson(ctx, base.VideoDataURL, param, param.ToValues(), header)
	if err != nil {
		return nil, dependency.ERROR_AUTH.Error(err)
	}
	if reply == nil {
		return nil, dependency.ERROR_AUTH.ErrorMsg("response nil")
	}
	res := &web.VideoDataResponse{}
	if err = json.Unmarshal(reply, res); err != nil {
		return nil, dependency.ERROR_AUTH.Error(err)
	}
	if res.Data.ErrorCode != 0 {
		return nil, dependency.ERROR_AUTH.ErrorMsg(res.Data.Description)
	}
	return res.Data.List, nil
}
