package oauth

import (
	"fmt"
	"net/http"
	"net/url"
	base "openapi/agent/web"
	"openapi/dependency"

	"openapi/dto/web"
)

// 文档：https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/account-permission/douyin-get-permission-code
// oauthRedirect 跳转到抖音二维码授权页面
func oauthRedirect(writer http.ResponseWriter, req *http.Request, clientKey, redirectUri, state, scope string) dependency.Catcher {
	param := web.OauthConnectRequest{
		ClientKey:    clientKey,
		RedirectUri:  redirectUri,
		ResponseType: "code",
		Scope:        scope,
		State:        state,
	}
	location := getRedirectURL(param)
	http.Redirect(writer, req, location, http.StatusFound)
	return nil
}

func getRedirectURL(param web.OauthConnectRequest) string {
	urlStr := url.QueryEscape(param.RedirectUri)
	return fmt.Sprintf(base.RedirectOauthURL, param.ClientKey, param.ResponseType, param.Scope, urlStr, param.State)
}
