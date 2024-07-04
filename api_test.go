package douyin_openapi

import (
	"context"
	"fmt"
	"openapi/agent/web/busi"
	"openapi/dependency"
	"openapi/pkg"
	"os"
	"testing"
)

var (
	api dependency.OpenApi
	ctx context.Context
	// 开放平台登记的回调地址
	redirectUri  = ""
	clientKey    = ""
	clientSecret = ""
	cacheKeyUser = "starkupa"
	cacheKeyBusi = "douyin_busi"
	// 账号权限
	scope = ""
)

func init() {
	cache, err := pkg.NewCache("127.0.0.1:6379", "root", "", 0)
	if err != nil {
		panic(err)
	}
	api = busi.NewDouyin(clientKey, clientSecret, cacheKeyUser, cacheKeyBusi, cache)
	ctx = context.Background()
}

func TestQrcode_Get(t *testing.T) {
	res, e := api.GetAuthQrCode(ctx, redirectUri, "0", scope)
	if e != nil {
		t.Fatal(e.Error())
	}
	t.Log(pkg.ToJsonString(res))
	// 存html，打开二维码
	s := fmt.Sprintf(`<img src="data:image/png;base64, %v" alt="Image">`, res.Qrcode)
	_ = os.WriteFile("qrcode.html", []byte(s), 0644)
}

func TestQrcode_check(t *testing.T) {
	// 取上面的token
	token := "72f9295e953fbbc40a37497bd56c637d_hl"
	res, e := api.CheckQrCodeAndSaveToken(ctx, token, redirectUri, "0", scope)
	if e != nil {
		t.Fatal(e.Error())
	}
	t.Log(pkg.ToJsonString(res))
}

func TestUserInfo(t *testing.T) {
	res, e := api.Userinfo(ctx)
	if e != nil {
		t.Fatal(e.Error())
	}
	t.Log(pkg.ToJsonString(res))
}

func TestVideoList(t *testing.T) {
	res, e := api.VideoList(ctx, 0, 10)
	if e != nil {
		t.Fatal(e.Error())
	}
	t.Log(pkg.ToJsonString(res))
}

func TestVideoData(t *testing.T) {
	// list读取的
	videoIds := []string{"7387337942391475494", "7387337781044940069"}
	res, e := api.VideoData(ctx, videoIds)
	if e != nil {
		t.Fatal(e.Error())
	}
	t.Log(pkg.ToJsonString(res))
}
