package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/orm"
	"io"
	"time"
	"zjlife/logger"
)

func Login(username string, password string, code string) (res map[string]interface{}) {
	o := orm.NewOrm()
	user := UserInfo{UserName: username, Password: password}
	_, wui, sk, expire := weChatAuthS(code, "", "", username, user.Userid)

	user.Expires = expire
	user.SessionKey = sk
	user.XcxOpenId = wui.OpenID

	h := md5.New()
	io.WriteString(h, code+time.Now().String())
	user.Token = fmt.Sprintf("%x", h.Sum(nil))
	if created, id, err := o.ReadOrCreate(&user, "UserName"); err == nil {
		if created {
			logger.Debug("New Insert an object. Id:", id)
		} else {
			logger.Error("Get an object. Id:", id)
		}
	} else {
		logger.Fatal(err)
	}

	res = make(map[string]interface{})
	res["token"] = user.Token
	res["userinfo"] = user
	return res
}

func Update(token string, userinfo map[string]interface{}) (res map[string]interface{}) {
	ui := UserInfo{}
	ui.Token = token

	o := orm.NewOrm()
	res = make(map[string]interface{})
	if o.Read(&ui, "token") == nil {
		ui.NickName = userinfo["NickName"].(string)
		ui.Avatar = userinfo["Avatar"].(string)
		if _, err := o.Update(&ui); err == nil {
			res["userinfo"] = ui
		} else {
			logger.Error(err)
		}
	}

	return res
}

func GetUserInfo(token string, userid string) (res map[string]interface{}) {
	ui := UserInfo{}
	ui.Token = token

	o := orm.NewOrm()
	res = make(map[string]interface{})
	if err := o.Read(&ui, "token"); err == nil {
		res["userinfo"] = ui
	} else {
		logger.Error(err)
	}

	return res
}
