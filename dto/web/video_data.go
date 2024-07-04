package web

import "net/url"

type VideoDataRequest struct {
	OpenID      string   `json:"open_id"`      // 通过/oauth/access_token/获取，用户唯一标志
	AccessToken string   `json:"access_token"` // 调用 /oauth/access_token/ 生成的 token
	ItemIDs     []string `json:"item_ids"`     // item_id 数组，仅能查询 access_token 对应用户上传的视频（与video_ids字段二选一，两者都传时平台只处理video_ids）
	VideoIDs    []string `json:"video_ids"`    // video_id 数组，仅能查询 access_token 对应用户上传的视频（与item_ids字段二选一，两者都传时平台只处理video_ids）
}

func (r *VideoDataRequest) ToValues() url.Values {
	u := url.Values{}
	u.Set("open_id", r.OpenID)
	return u
}

type VideoDataResponse struct {
	Data VideoDetailsData `json:"data"`
}

type VideoDetailsData struct {
	ErrorCode   int64       `json:"error_code"`
	Description string      `json:"description"`
	List        []VideoItem `json:"list"`
}

type VideoItem struct {
	Title       string      `json:"title"`
	CreateTime  int64       `json:"create_time"`
	VideoStatus int64       `json:"video_status"`
	ShareURL    string      `json:"share_url"`
	Cover       string      `json:"cover"`
	IsTop       bool        `json:"is_top"`
	Statistics  DStatistics `json:"statistics"`
	ItemID      string      `json:"item_id"`
	IsReviewed  bool        `json:"is_reviewed"`
	MediaType   int64       `json:"media_type"`
	VideoID     string      `json:"video_id"`
}

type DStatistics struct {
	DiggCount     int64 `json:"digg_count"`
	DownloadCount int64 `json:"download_count"`
	PlayCount     int64 `json:"play_count"`
	ShareCount    int64 `json:"share_count"`
	ForwardCount  int64 `json:"forward_count"`
	CommentCount  int64 `json:"comment_count"`
}
