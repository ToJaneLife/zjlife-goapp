package controllers

import (
	"reflect"
)

type (
	Resp struct {
		Data interface{} `json:"data"`
		Desc string      `json:"desc"`
		Code int         `json:"code"`
	}
	Auth struct {
		Token string `json:"token"`
	}
	UserInfo struct {
		Userid   int    `json:"userid"`
		UserName string `json:"username"`
		Avatar   string `json:"avatar"`
		NickName string `json:"nick_name"`
		Password string `json:"password"`
	}
	WchatAuth struct {
		AppId         string `json:"app_id"`
		Code          string `json:"code"`
		Iv            string `json:"iv"`
		EncryptedData string `json:"encrypted_data"`
	}
	ReqLogin struct {
		UserName string    `json:"username"`
		Password string    `json:"password"`
		WuAuth   WchatAuth `json:"wuAuth"`
	}
	ReqUpdate struct {
		Auth
		UserInfo UserInfo `json:"userinfo"`
	}
	ReqGetInfo struct {
		Auth
		Userid string `json:"userid"`
	}
)

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
