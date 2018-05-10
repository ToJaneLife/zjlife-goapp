package controllers

import (
	"zjlife/models"

	"github.com/astaxie/beego"
)

// Operations about object
type WeatherController struct {
	beego.Controller
}

// @Title SearchWeather
// @Description 查询实时天气和天气预报
// @Param	city		query 	string	true		"The city for query"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /weatherInfo [get]
func (u *WeatherController) SearchWeather() {
	city := u.GetString("city")
	data := make(map[string]interface{})

	if len(city) != 0 {
		result := models.SearchWeather(city)
		u.Data["json"] = result
	} else {
		data["desc"] = "缺少city 字段"
		data["code"] = 2
		data["data"] = "没有数据"
		u.Data["json"] = data
	}
	u.ServeJSON()
}
