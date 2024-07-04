package web

import (
	"fmt"
	"net/url"
)

type VideoListRequest struct {
	OpenID      string `json:"open_id"`      // 通过/oauth/access_token/获取，用户唯一标志
	Count       int64  `json:"count"`        // 每页数量
	AccessToken string `json:"access_token"` // 调用 /oauth/access_token/ 生成的 token
	Cursor      int64  `json:"cursor"`       // 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
}

func (r *VideoListRequest) ToValues() url.Values {
	data := url.Values{}
	data.Set("count", fmt.Sprintf("%v", r.Count))
	data.Set("open_id", r.OpenID)
	data.Set("cursor", fmt.Sprintf("%v", r.Cursor))
	return data
}

type VideoListResponse struct {
	Data VideoListData `json:"data"`
}

type VideoListData struct {
	ErrorCode   int64      `json:"error_code"`  // 错误码
	Description string     `json:"description"` // 错误描述
	HasMore     bool       `json:"has_more"`    // 是否还有更多数据
	List        []ListItem `json:"list"`        // 数据
	Cursor      int64      `json:"cursor"`      // 用于下一页请求的cursor
}

type ListItem struct {
	Title       string      `json:"title"`        // 视频标题
	IsTop       bool        `json:"is_top"`       // 是否置顶
	CreateTime  int64       `json:"create_time"`  // 视频创建时间戳
	IsReviewed  bool        `json:"is_reviewed"`  // 表示是否审核结束。审核通过或者失败都会返回true，审核中返回false
	VideoStatus int64       `json:"video_status"` // 表示视频状态。1:细化为5、6、7三种状态;2:不适宜公开;4:审核中;5:公开视频;6:好友可见;7:私密视频
	ShareURL    string      `json:"share_url"`    // 视频播放页面。视频播放页可能会失效，请在观看视频前调用/video/data/获取最新的播放页。
	ItemID      string      `json:"item_id"`      // 视频id
	MediaType   int64       `json:"media_type"`   // 媒体类型
	Cover       string      `json:"cover"`        // 视频封面
	Statistics  LStatistics `json:"statistics"`   // 统计数据
	VideoID     string      `json:"video_id"`     // 视频真实id
}

type LStatistics struct {
	ForwardCount  int64 `json:"forward_count"`  // 转发数
	CommentCount  int64 `json:"comment_count"`  // 评论数
	DiggCount     int64 `json:"digg_count"`     // 点赞数
	DownloadCount int64 `json:"download_count"` // 下载数
	PlayCount     int64 `json:"play_count"`     // 播放数，只有作者本人可见。公开视频设为私密后，播放数也会返回0
	ShareCount    int64 `json:"share_count"`    // 分享数
}
