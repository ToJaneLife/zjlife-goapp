package controllers

import (
	"zjlife/models"

	"github.com/astaxie/beego"
)

// Operations about object
type ExpressController struct {
	beego.Controller
}

// @Title EbusinessOrderHandle
// @Description 查询快递订单和状态
// @Param	RequestData		query 	string	true		"The RequestData for query"
// @Param	RequestType		query 	string	true		"The RequestType for query"
// @Success 200 {string} success
// @Failure 401 wei chaxun dao danhao
// @router /ebusinessOrderHandle [post]
func (e *ExpressController) EbusinessOrderHandle() {
	requestData := e.GetString("RequestData")
	requestType := e.GetString("RequestType")

	data := make(map[string]interface{})

	if (len(requestData) != 0) {
		result := models.EbusinessOrderHandle(requestData, requestType)
		e.Data["json"] = result
	} else {
		data["desc"] = "缺少requestData 字段"
		data["code"] = 2
		data["data"] = "没有数据"
		e.Data["json"] = data
	}
	e.ServeJSON()
}