package web

const (
	// 抖音开放平台授权
	RedirectOauthURL = "https://open.douyin.com/platform/oauth/connect?client_key=%s&response_type=%s&scope=%s&redirect_uri=%s&state=%s"
	// 抖音开放平台获取access_token
	AccessTokenURL = "https://open.douyin.com/oauth/access_token"
	// 抖音开放平台刷新refresh_token
	RenewRefreshTokenURL = "https://open.douyin.com/oauth/renew_refresh_token"
	// 抖音开放平台刷新access_token
	RefreshAccessTokenURL = "https://open.douyin.com/oauth/refresh_token"

	// 获取二维码
	GetQrCodeURL = "https://open.douyin.com/oauth/get_qrcode"
	// 检测二维码
	CheckQrCodeURL = "https://open.douyin.com/oauth/check_qrcode"

	// 查询用户
	UserInfoURL = "https://open.douyin.com/oauth/userinfo"
)

const (
	// 查询授权账号视频列表
	VideoListURL = "https://open.douyin.com/api/douyin/v1/video/video_list"
	// 查询授权账号特定视频数据
	VideoDataURL = "https://open.douyin.com/api/douyin/v1/video/video_data"
)
