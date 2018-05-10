package models

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/httplib"
)

var (
	background map[string]string
)

func init() {
	background = make(map[string]string)

	background["大雨"] = "http://7xqhcq.com1.z0.glb.clouddn.com/zj-weapp/thunderstorm-358992_1280.jpg"
	background["中雨"] = "http://7xqhcq.com1.z0.glb.clouddn.com/zj-weapp/weatherman-849792_1280.jpg"
	background["小雨"] = "http://7xqhcq.com1.z0.glb.clouddn.com/zj-weapp/rain-985874_1280.jpg"

	background["阵雨"] = "http://omiz2siz5.bkt.clouddn.com/zj-weapp/xiaoyu.jpg"
	background["暴雨"] = "http://7xqhcq.com1.z0.glb.clouddn.com/zj-weapp/lightning-503157_1920.jpg"
	background["雷阵雨"] = "http://7xqhcq.com1.z0.glb.clouddn.com/zj-weapp/lightning-2102435_1920.jpg"

	background["晴"] = "http://omiz2siz5.bkt.clouddn.com/zj-weapp/qing2.jpg"
	background["阴"] = "http://omiz2siz5.bkt.clouddn.com/zj-weapp/yin.jpg"
	background["多云"] = "http://omiz2siz5.bkt.clouddn.com/zj-weapp/duoyun2.jpg"

	background["雾"] = "http://omiz2siz5.bkt.clouddn.com/zj-weapp/wu.jpg"
	background["小雪"] = "http://omiz2siz5.bkt.clouddn.com/zj-weapp/xue.jpg"
	background["大雪"] = "http://omiz2siz5.bkt.clouddn.com/zj-weapp/xue.jpg"

	background["暴雪"] = "http://omiz2siz5.bkt.clouddn.com/zj-weapp/xue.jpg"
}

func SearchWeather(city string) map[string]interface{} {
	//生成要访问的url
	url := "http://restapi.amap.com/v3/weather/weatherInfo?key=2749361c6fa6a5c850dd426c9d07827d&extensions=all&city=" + city
	//提交请求
	req := httplib.Get(url)

	str, err := req.String()
	if err != nil {
		panic(err)
	}

	//实时天气
	url2 := "http://restapi.amap.com/v3/weather/weatherInfo?key=2749361c6fa6a5c850dd426c9d07827d&extensions=base&city=" + city
	//提交请求
	req2 := httplib.Get(url2)

	str2, err2 := req2.String()
	if err2 != nil {
		panic(err2)
	}

	var newWeather AMapWeather
	var newLives AMapLives

	data := make(map[string]interface{})

	// 将高德返回的数据 转化成自定义的结构
	if err := json.Unmarshal([]byte(str), &newWeather); err == nil {
		lastRes := WeatherTanslate(newWeather)
		if err2 := json.Unmarshal([]byte(str2), &newLives); err2 == nil {
			lastRes.Live = newLives.Lives[0]

			value, _ := background[lastRes.Live.Weather]
			lastRes.Live.BackgroundImageUrl = value
			fmt.Sprintf(lastRes.Live.BackgroundImageUrl)
		}

		data["code"] = newWeather.Infocode
		data["data"] = lastRes
		data["desc"] = newWeather.Info
	}

	return data
}

func WeatherTanslate(data AMapWeather) (result Weather) {
	result.City = data.Forecasts[0].City
	result.Reporttime = data.Forecasts[0].Reporttime
	result.Adcode = data.Forecasts[0].Adcode
	result.Province = data.Forecasts[0].Province

	if len(data.Forecasts[0].Casts) != 0 {
		cast := make([]Forecasts, 0)
		for i := 0; i < len(data.Forecasts[0].Casts); i++ {
			amap := data.Forecasts[0].Casts[i]

			temp := Forecasts{}

			temp.Date = amap.Date
			temp.Daypower = amap.Daypower
			temp.Daytemp = amap.Daytemp
			temp.Dayweather = amap.Dayweather
			temp.Daywind = amap.Daywind
			temp.Nightpower = amap.Nightpower
			temp.Nighttemp = amap.Nighttemp
			temp.Nightweather = amap.Nightweather
			temp.Nightwind = amap.Nightwind
			temp.Week = WeekTranslate(amap.Week)

			cast = append(cast, temp)
		}

		result.Forecasts = cast
	}

	return result
}

func WeekTranslate(week string) (result string) {
	switch week {
	case "1":
		result = "星期一"
	case "2":
		result = "星期二"
	case "3":
		result = "星期三"
	case "4":
		result = "星期四"
	case "5":
		result = "星期五"
	case "6":
		result = "星期六"
	case "7":
		result = "星期日"
	}

	return result
}
