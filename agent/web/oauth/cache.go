package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"openapi/dto/web"
)

const (
	accessTokenInfoKey  = "douyin_access_token"
	refreshTokenInfoKey = "douyin_refresh_token"
)

func (o *Oauth) getAccessTokenKey() string {
	return fmt.Sprintf("%v:%v:%v", o.keyPrefix, accessTokenInfoKey, o.identity)
}

func (o *Oauth) getRefreshTokenKey() string {
	return fmt.Sprintf("%v:%v:%v", o.keyPrefix, refreshTokenInfoKey, o.identity)
}

func (o *Oauth) GetCacheRefreshToken(ctx context.Context) (*web.AccessTokenCoreData, error) {
	if o.cache == nil {
		return nil, fmt.Errorf("cache is nil")
	}
	key := o.getRefreshTokenKey()
	if i := o.cache.Get(ctx, key); i == "" {
		return nil, nil
	} else {
		data := &web.AccessTokenCoreData{}
		if err := json.Unmarshal([]byte(i), data); err != nil {
			return nil, err
		}
		return data, nil
	}
}

func (o *Oauth) GetCacheAccessToken(ctx context.Context) (*web.AccessTokenCoreData, error) {
	if o.cache == nil {
		return nil, fmt.Errorf("cache is nil")
	}
	key := o.getAccessTokenKey()
	if i := o.cache.Get(ctx, key); i == "" {
		return nil, nil
	} else {
		data := &web.AccessTokenCoreData{}
		if err := json.Unmarshal([]byte(i), data); err != nil {
			return nil, err
		}
		return data, nil
	}
}

func (o *Oauth) setCacheAccessToken(ctx context.Context, token *web.AccessTokenCoreData, expire time.Duration) error {
	if o.cache == nil {
		return fmt.Errorf("cache nil")
	}
	if o.identity == "" {
		o.identity = token.OpenID
	}
	key := o.getAccessTokenKey()
	by, _ := json.Marshal(token)
	return o.cache.Set(ctx, key, string(by), expire)
}

func (o *Oauth) setCacheRefreshToken(ctx context.Context, token *web.AccessTokenCoreData, expire time.Duration) error {
	if o.cache == nil {
		return fmt.Errorf("cache nil")
	}
	if o.identity == "" {
		o.identity = token.OpenID
	}
	key := o.getRefreshTokenKey()
	by, _ := json.Marshal(token)
	return o.cache.Set(ctx, key, string(by), expire)
}
