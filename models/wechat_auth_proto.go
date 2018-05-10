package models

import (
	"time"
)

//wechat认证通讯协议
type (
	WchatUserInfo struct {
		OpenID    string `json:"openId"`
		UnionID   string `json:"unionId"`
		NickName  string `json:"nickName"`
		Gender    int    `json:"gender"`
		City      string `json:"city"`
		Province  string `json:"province"`
		Country   string `json:"country"`
		AvatarURL string `json:"avatarUrl"`
		Language  string `json:"language"`
		Watermark struct {
			Timestamp int64  `json:"timestamp"`
			AppID     string `json:"appid"`
		} `json:"watermark"`
	}
	WchatSessionOpenId struct {
		SessionKey string `json:"session_key" orm:"column(session_key)"`
		Expires    int    `json:"expires_in" orm:"column(expires_in)"`
		XcxOpenId  string `json:"openid" orm:"column(openid)"`
	}
	WchatBizDataCrypt struct {
		AppId      string `json:"app_id"`
		SessionKey string `json:"session_key"`
	}
	UserWechatSessionKey struct {
		UserId            int
		SessionKey        string
		SessionExpires    int
		SessionKeyUpdTime time.Time
		OpenId            string
		UnionId           string
	}
)
