package oauth

import (
	"context"
	"net/http"
	"openapi/dependency"
	"openapi/dto/web"
	"sync"
	"time"
)

// 若 access_token 未过期，刷新 refresh_token 不会改变原来的 access_token，但超时时间会更新，相当于续期。最多能再获取 5 次新的 refresh_token，最长续期为 15 + 30 + 30 * 5 = 195 天。
// 若 access_token 已过期，调用接口会报错（error_code=10008 或 2190008）。可以通过刷新 access_token 接口进行刷新。刷新后重新获得一个有效期为 15 天的 access_token，但 refresh_token 的有效期保持不变。
// 若 refresh_token 未过期，可以通过刷新 refresh_token 获取新的 refresh_token。
// 若 refresh_token 已过期，获取 access_token 会报错（error_code=10010），且不能通过刷新 refresh_token 获取新的 refresh_token。此时需要重新引导用户授权。

// Oauth 保存用户授权信息
type Oauth struct {
	Identity
	cache        dependency.Cache
	clientKey    string
	clientSecret string
	lock         *sync.Mutex
}

type Identity struct {
	keyPrefix string // cache key prefix
	identity  string // eg. application inner user id
}

// NewOauth 实例化授权信息
func NewOauth(clientKey, clientSecret, identity, keyPrefix string, cache dependency.Cache) *Oauth {
	id := Identity{keyPrefix, identity}
	return &Oauth{
		clientKey:    clientKey,
		clientSecret: clientSecret,
		Identity:     id,
		cache:        cache,
		lock:         &sync.Mutex{},
	}
}

func (o *Oauth) GetClientKey() string {
	return o.clientKey
}

func (o *Oauth) GetClientSecret() string {
	return o.clientSecret
}

// Redirect 跳转到抖音二维码授权页面
func (o *Oauth) Redirect(writer http.ResponseWriter, req *http.Request, clientKey, redirectUri, state, scope string) dependency.Catcher {
	return oauthRedirect(writer, req, clientKey, redirectUri, state, scope)
}

// getAuthQrCode 获取二维码
func (o *Oauth) GetAuthQrCode(ctx context.Context, redirectUri, state, scope string) (*web.TiktokGetAuthCodeData, dependency.Catcher) {
	return getAuthQrCode(ctx, o.GetClientKey(), redirectUri, state, scope)
}

// checkAuthQrCode check二维码
func (o *Oauth) CheckAuthQrCode(ctx context.Context, token, redirectUri, state, scope string) (*web.TiktokCheckAuthCodeData, dependency.Catcher) {
	return checkAuthQrCode(ctx, o.GetClientKey(), token, redirectUri, state, scope)
}

// CheckQrCodeAndSaveToken 检测二维码，授权保存 access_token
func (o *Oauth) CheckQrCodeAndSaveToken(ctx context.Context, token, redirectUri, state, scope string) (*web.TiktokCheckAuthCodeData, dependency.Catcher) {
	reply, ex := o.CheckAuthQrCode(ctx, token, redirectUri, state, scope)
	if ex != nil {
		return nil, ex
	}
	if reply.Status != "confirmed" {
		return reply, nil
	}
	ac, ex := o.GetAccessToken(ctx, reply.Code)
	if ex != nil {
		return nil, dependency.ERROR_BUSI.Error(ex.Error())
	}
	reply.OpenID = ac.OpenID
	return reply, nil
}

// getAccessToken 获取 access_token
func (o *Oauth) GetAccessToken(ctx context.Context, code string) (*web.AccessTokenData, dependency.Catcher) {
	data, ex := getAccessToken(ctx, o.GetClientKey(), o.GetClientSecret(), code)
	if ex != nil {
		return nil, ex
	}
	// 缓存 access_token
	if err := o.setCacheAccessToken(ctx, &data.AccessTokenCoreData, time.Duration(data.ExpiresIn-3600)*time.Second); err != nil {
		return nil, dependency.ERROR_AUTH.Error(err)
	}
	// 缓存 refresh_token
	if err := o.setCacheRefreshToken(ctx, &data.AccessTokenCoreData, time.Duration(data.RefreshExpiresIn-3600)*time.Second); err != nil {
		return nil, dependency.ERROR_AUTH.Error(err)
	}
	return data, nil
}

// GetAndRefreshAccessToken 从cache获取 access_token，若无则从抖音获取
func (o *Oauth) GetAndRefreshAccessToken(ctx context.Context) (*web.AccessTokenData, dependency.Catcher) {
	// 先从cache中取access_token
	if at, err := o.GetCacheAccessToken(ctx); err != nil {
		return nil, dependency.ERROR_AUTH.Error(err)
	} else if at != nil {
		return &web.AccessTokenData{AccessTokenCoreData: *at}, nil
	}

	// 加上lock，是为了防止在并发获取token时，cache刚好失效，导致从抖音服务器上获取到不同token
	o.lock.Lock()
	defer o.lock.Unlock()

	// 双检，防止重复从抖音服务器获取
	if at, err := o.GetCacheAccessToken(ctx); err != nil {
		return nil, dependency.ERROR_AUTH.Error(err)
	} else if at != nil {
		return &web.AccessTokenData{AccessTokenCoreData: *at}, nil
	}

	// cache失效，通过refresh_token刷新access_token
	ak, err := o.GetCacheRefreshToken(ctx)
	if err != nil {
		return nil, dependency.ERROR_AUTH.Error(err)
	}
	// 若refresh_token失效，返回重新授权
	if ak == nil {
		return nil, dependency.ERROR_AUTH.ErrorMsg("refresh_token is nil")
	}
	// 刷新access_token
	newAt, ex := o.RefreshAccessToken(ctx, ak.RefreshToken)
	if ex != nil {
		return nil, ex
	}
	if err = o.setCacheAccessToken(ctx, &newAt.AccessTokenCoreData, time.Duration(newAt.ExpiresIn-3600)*time.Second); err != nil {
		return nil, dependency.ERROR_AUTH.Error(err)
	}
	if err = o.setCacheRefreshToken(ctx, &newAt.AccessTokenCoreData, time.Duration(newAt.RefreshExpiresIn-3600)*time.Second); err != nil {
		return nil, dependency.ERROR_AUTH.Error(err)
	}
	return &newAt.AccessTokenData, nil
}

// refreshAccessToken 刷新access_token
func (o *Oauth) RefreshAccessToken(ctx context.Context, refreshToken string) (*web.RefreshAccessTokenData, dependency.Catcher) {
	return refreshAccessToken(ctx, o.GetClientKey(), refreshToken)
}

// renewRefreshToken 刷新 refresh_token
func (o *Oauth) RenewRefreshToken(ctx context.Context, refreshToken string) (*web.RenewRefreshTokenData, dependency.Catcher) {
	return renewRefreshToken(ctx, o.GetClientKey(), refreshToken)
}
