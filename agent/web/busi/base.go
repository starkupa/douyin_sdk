package busi

import (
	"context"
	"openapi/agent/web/oauth"
	"openapi/dependency"
	"openapi/dto/web"
)

type Douyin struct {
	*oauth.Oauth
}

// NewDouyin 抖音业务
//
//	@param clientKey    应用key
//	@param clientSecret 应用密钥
//	@param identity     身份标识，key对应用户
//	@param keyPrefix    key业务前缀
//	@param cache        缓存
//	@return Douyin
func NewDouyin(clientKey, clientSecret, identity, keyPrefix string, cache dependency.Cache) *Douyin {
	au := oauth.NewOauth(clientKey, clientSecret, identity, keyPrefix, cache)
	return &Douyin{Oauth: au}
}

// VideoList 获取视频列表
func (d *Douyin) VideoList(ctx context.Context, cursor, count int64) (*web.VideoListData, dependency.Catcher) {
	token, ex := d.GetAndRefreshAccessToken(ctx)
	if ex != nil {
		return nil, ex
	}
	if token == nil {
		return nil, dependency.ERROR_AUTH.ErrorMsg("token is nil")
	}
	res, ex := videoList(ctx, token.AccessToken, token.OpenID, cursor, count)
	if ex != nil {
		return nil, ex
	}
	return res, nil
}

// VideoData 获取视频数据
func (d *Douyin) VideoData(ctx context.Context, videoIds []string, itemIds ...string) ([]web.VideoItem, dependency.Catcher) {
	token, ex := d.GetAndRefreshAccessToken(ctx)
	if ex != nil {
		return nil, ex
	}
	if token == nil {
		return nil, dependency.ERROR_AUTH.ErrorMsg("token is nil")
	}
	res, ex := videoData(ctx, token.AccessToken, token.OpenID, videoIds, itemIds...)
	if ex != nil {
		return nil, ex
	}
	return res, nil
}

// Userinfo 获取用户信息
func (d *Douyin) Userinfo(ctx context.Context) (*web.UserInfoData, dependency.Catcher) {
	token, ex := d.GetAndRefreshAccessToken(ctx)
	if ex != nil {
		return nil, ex
	}
	if token == nil {
		return nil, dependency.ERROR_AUTH.ErrorMsg("token is nil")
	}
	res, ex := userInfo(ctx, token.AccessToken, token.OpenID)
	if ex != nil {
		return nil, ex
	}
	return res, nil
}
