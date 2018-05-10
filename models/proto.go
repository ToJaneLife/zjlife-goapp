package models

import (
	"github.com/astaxie/beego/orm"
)

type (
	Auth struct {
		Token string `json:"-" orm:"column(token)"`
	}
	UserInfo struct {
		Userid   int    `json:"userid" orm:"pk; auto; column(userid)"`
		UserName string `json:"username" orm:"column(username)"`
		Avatar   string `json:"avatar" orm:"column(avatar)"`
		NickName string `json:"nick_name" orm:"column(nick_name)"`
		Password string `json:"-" orm:"column(password)"`
		Auth
		WchatSessionOpenId
	}
	WchatAuth struct {
		AppId         string `json:"app_id"`
		Code          string `json:"code"`
		Iv            string `json:"iv"`
		EncryptedData string `json:"encrypted_data"`
	}
	ReqLogin struct {
		UserName string `json:"username"`
		Password string `json:"password"`
		WchatAuth
	}
)

// 高德天气结构
type AMapWeather struct {
	Count     string          `json:"count"`
	Forecasts []AMapForecasts `json:"forecasts"`
	Info      string          `json:"info"`
	Infocode  string          `json:"infocode"`
	Status    string          `json:"status"`
}

type AMapForecasts struct {
	Adcode     string      `json:"adcode"`
	Casts      []AMapCasts `json:"casts"`
	City       string      `json:"city"`
	Province   string      `json:"province"`
	Reporttime string      `json:"reporttime"`
}

type AMapCasts struct {
	Date         string `json:"date"`
	Daypower     string `json:"daypower"`
	Daytemp      string `json:"daytemp"`
	Dayweather   string `json:"dayweather"`
	Daywind      string `json:"daywind"`
	Nightpower   string `json:"nightpower"`
	Nighttemp    string `json:"nighttemp"`
	Nightweather string `json:"nightweather"`
	Nightwind    string `json:"nightwind"`
	Week         string `json:"week"`
}

// 自定义结构
type Weather struct {
	City       string      `json:"city"`
	Live       Live        `json:"live"`
	Forecasts  []Forecasts `json:"forecasts"`
	Province   string      `json:"province"`
	Reporttime string      `json:"reporttime"`
	Adcode     string      `json:"adcode"`
}

type AMapLives struct {
	Lives []Live `json:"lives"`
}

type Live struct {
	Weather            string `json:"weather"`
	Temperature        string `json:"temperature"`
	Winddirection      string `json:"winddirection"`
	Windpower          string `json:"windpower"`
	Humidity           string `json:"humidity"`
	BackgroundImageUrl string `json:"backgroundImageUrl"`
}

type Forecasts struct {
	Date         string `json:"date"`
	Daypower     string `json:"daypower"`
	Daytemp      string `json:"daytemp"`
	Dayweather   string `json:"dayweather"`
	Daywind      string `json:"daywind"`
	Nightpower   string `json:"nightpower"`
	Nighttemp    string `json:"nighttemp"`
	Nightweather string `json:"nightweather"`
	Nightwind    string `json:"nightwind"`
	Week         string `json:"week"`
}

func init() {
	orm.RegisterModel(new(UserInfo))
}
