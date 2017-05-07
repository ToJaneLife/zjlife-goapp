package models

import (
	"encoding/json"

	"github.com/astaxie/beego/httplib"
)

// 高德天气结构
type AMapWeather struct {
	Count string `json:"count"`
	Forecasts []AMapForecasts `json:"forecasts"`
	Info string `json:"info"`
	Infocode string `json:"infocode"`
	Status string `json:"status"`
}

type AMapForecasts struct {
	Adcode string `json:"adcode"`
	Casts []AMapCasts `json:"casts"`
	City string `json:"city"`
	Province string `json:"province"`
	Reporttime string `json:"reporttime"`
}

type AMapCasts struct {
	Date string `json:"date"`
  	Daypower string `json:"daypower"`
  	Daytemp string `json:"daytemp"`
  	Dayweather string `json:"dayweather"`
  	Daywind string `json:"daywind"`
  	Nightpower string `json:"nightpower"`
  	Nighttemp string `json:"nighttemp"`
  	Nightweather string `json:"nightweather"`
  	Nightwind string `json:"nightwind"`
  	Week string `json:"week"`
}

// 自定义结构
type Weather struct {
	City string `json:"city"`
	Live Live `json:"live"`
	Forecasts []Forecasts `json:"forecasts"`
	Province string `json:"province"`
	Reporttime string `json:"reporttime"`
	Adcode string `json:"adcode"`
}

    // "province": "上海",
    // "city": "上海市",
    // "adcode": "310000",
    // "weather": "多云",
    // "temperature": "23",
    // "winddirection": "北",
    // "windpower": "5",
    // "humidity": "44",
    // "reporttime": "2017-05-06 18:00:00"

type AMapLives struct {
	Lives []Live `json:"lives"`
}

type Live struct {
	Weather string `json:"weather"`
	Temperature string `json:"temperature"`
	Winddirection string `json:"winddirection"`
	Windpower string `json:"windpower"`
	Humidity string `json:"humidity"`
}

type Forecasts struct {
	Date string `json:"date"`
  	Daypower string `json:"daypower"`
  	Daytemp string `json:"daytemp"`
  	Dayweather string `json:"dayweather"`
  	Daywind string `json:"daywind"`
  	Nightpower string `json:"nightpower"`
  	Nighttemp string `json:"nighttemp"`
  	Nightweather string `json:"nightweather"`
  	Nightwind string `json:"nightwind"`
  	Week string `json:"week"`
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
