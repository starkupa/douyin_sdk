package dependency

import (
	"context"
	"net/http"
	"openapi/dto/web"
)

type OpenApi interface {
	// 跳转到抖音二维码授权页面
	Redirect(writer http.ResponseWriter, req *http.Request, clientKey, redirectUri, state, scope string) Catcher
	// 获取二维码
	GetAuthQrCode(ctx context.Context, redirectUri, state, scope string) (*web.TiktokGetAuthCodeData, Catcher)
	// check二维码
	CheckAuthQrCode(ctx context.Context, token, redirectUri, state, scope string) (*web.TiktokCheckAuthCodeData, Catcher)
	// 获取 access_token
	GetAccessToken(ctx context.Context, code string) (*web.AccessTokenData, Catcher)
	// 从cache获取 access_token，若无则从抖音获取
	GetAndRefreshAccessToken(ctx context.Context) (*web.AccessTokenData, Catcher)
	// 刷新access_token
	RefreshAccessToken(ctx context.Context, refreshToken string) (*web.RefreshAccessTokenData, Catcher)
	// 刷新 refresh_token
	RenewRefreshToken(ctx context.Context, refreshToken string) (*web.RenewRefreshTokenData, Catcher)
	// 检测二维码，授权保存 access_token
	CheckQrCodeAndSaveToken(ctx context.Context, token, redirectUri, state, scope string) (*web.TiktokCheckAuthCodeData, Catcher)
	// 获取视频列表
	VideoList(ctx context.Context, cursor, count int64) (*web.VideoListData, Catcher)
	// 获取视频数据
	VideoData(ctx context.Context, videoIds []string, itemIds ...string) ([]web.VideoItem, Catcher)
	// 获取用户信息
	Userinfo(ctx context.Context) (*web.UserInfoData, Catcher)
}
