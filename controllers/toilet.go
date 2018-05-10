package controllers

import (
	"zjlife/models"

	"github.com/astaxie/beego"
)

// Operations about object
type ToiletController struct {
	beego.Controller
}

// @Title SearchToilet
// @Description 查询位置附近的厕所
// @Param	location		query 	string	true		"The location for query"
// @Success 1000 {string} query success
// @Failure 999 error
// @router /around [post]
func (t *ToiletController) SearchToilet() {
	location := t.GetString("location")
	data := make(map[string]interface{})

	if len(location) != 0 {
		result := models.SearchToilet(location)
		t.Data["json"] = result
	} else {
		data["desc"] = "缺少location 字段"
		data["code"] = 2
		data["data"] = "没有数据"
		t.Data["json"] = data
	}
	t.ServeJSON()
}
