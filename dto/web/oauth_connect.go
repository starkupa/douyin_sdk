package web

type OauthConnectRequest struct {
	ClientKey     string `json:"client_key"`     // 应用唯一标识
	ResponseType  string `json:"response_type"`  // 写死为 code 即可
	Scope         string `json:"scope"`          // 应用授权作用域，多个授权作用域以英文逗号（,）分隔
	OptionalScope string `json:"optional_scope"` // 应用授权可选作用域
	RedirectUri   string `json:"redirect_uri"`   // 授权成功后的回调地址
	State         string `json:"state"`          // 用于保持请求和回调的状态
}
