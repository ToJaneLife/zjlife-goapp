package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "tolife.yuhanle.com"
	c.Data["Email"] = "shiboven@foxmail.com"
	c.TplName = "index.tpl"
}
