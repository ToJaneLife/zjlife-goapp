package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"zjlife/logger"
	"zjlife/models"
)

// Operations about object
type UserController struct {
	beego.Controller
}

// @Title Login
// @Description 登录系统
// @Param	username	query string	true	"username"
// @Param	password	query string	true	"password"
// @Param	weinfo		query interface{}	true	"weinfo"
// @Success 00 login success
// @Failure 01 user not exist
// @router /login [post]
func (u *UserController) Login() {
	var req ReqLogin
	json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	resp := new(Resp)

	if len(req.UserName) != 0 || len(req.WuAuth.Code) != 0 {
		result := models.Login(req.UserName, req.Password, req.WuAuth.Code)
		resp.Desc = "成功"
		resp.Code = 00
		resp.Data = result
		u.Data["json"] = resp
	} else {
		resp.Desc = "无效参数"
		resp.Code = 2
		resp.Data = nil
		u.Data["json"] = resp
	}
	u.ServeJSON()
}

// @Title Update
// @Description 登录系统
// @Param	token	query string	true	"token"
// @Success 00 update success
// @router /update [post]
func (u *UserController) Update() {
	var req ReqUpdate
	json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	logger.Debug(req.Token)

	result := models.Update(req.Token, Struct2Map(req.UserInfo))
	resp := new(Resp)
	resp.Desc = "成功"
	resp.Code = 00
	resp.Data = result
	u.Data["json"] = resp
	u.ServeJSON()
}

// @Title GetUserInfo
// @Description 获取用户信息
// @Param	token		query string	true	"token"
// @Param	userid	query string	true	"userid"
// @Success 00 info success
// @router /info [post]
func (u *UserController) GetUserInfo() {
	var req ReqGetInfo
	json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	logger.Debug(req.Token)

	result := models.GetUserInfo(req.Token, req.Userid)
	resp := new(Resp)
	resp.Desc = "成功"
	resp.Code = 00
	resp.Data = result
	u.Data["json"] = resp
	u.ServeJSON()
}
