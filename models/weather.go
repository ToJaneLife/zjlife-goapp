package models

import (
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego/httplib"
)

var (
	
)

type Weather struct {
	count int32 `json:"count"`
	forecasts Forecasts `json:"forecasts"`
	info string `json:"info"`
	infocode int32 `json:"infocode"`
	status int32 `json:"status"`
}

type Forecasts struct {
	adcode int32 `json:"adcode"`
	casts []Casts `json:"casts"`
}

type Casts struct {
	date string `json:"date"`
  	daypower float64 `json:"daypower"`
  	daytemp float64 `json:"daytemp"`
  	dayweather string `json:"dayweather"`
  	daywind string `json:"daywind"`
  	nightpower float64 `json:"nightpower"`
  	nighttemp float64 `json:"nighttemp"`
  	nightweather string `json:"nightweather"`
  	nightwind string `json:"nightwind"`
  	week int32 `json:"week"`
}

func init() {
	
}

func SearchWeather(city string) (result map[string]interface{}) {
	//生成要访问的url
	url := "http://restapi.amap.com/v3/weather/weatherInfo?key=2749361c6fa6a5c850dd426c9d07827d&extensions=all&city=" + city
	//提交请求
	req := httplib.Get(url)
	
	str, err := req.String()
	if err != nil {
		panic(err)
	}

	fmt.Println(str)
	
	req.ToJSON(&result)

	var newWeather Weather

	// 将高德返回的数据 转化成自定义的结构
	if err := json.Unmarshal([]byte(str), &newWeather); err == nil {
	    fmt.Println(newWeather)
	    fmt.Println(newWeather.info)
	}

	data := make(map[string]interface{})
	data["code"] = result["infocode"]
	data["data"] = result["forecasts"]
	data["desc"] = result["info"]

	return data
}
