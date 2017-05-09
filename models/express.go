package models

import (
	"encoding/base64"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/astaxie/beego/httplib"
)

var (
	appKey string
	eBusinessID string
)

func init() {
	appKey = "ee1b1872-1d38-45de-bb1f-7e5a40d41667"
	eBusinessID = "1286240"
}

func EbusinessOrderHandle(requestData string) (result map[string]interface{}) {
	requestType := "2002"
	w := md5.New()
	io.WriteString(w, requestData + appKey)   //将str写入到w中
	md5str := fmt.Sprintf("%x", w.Sum(nil))  //w.Sum(nil)将w的hash转成[]byte格式
	dataSign := base64.StdEncoding.EncodeToString([]byte(md5str))

	//生成要访问的url
	url := "http://api.kdniao.cc/Ebusiness/EbusinessOrderHandle.aspx"
	//提交请求
	req := httplib.Post(url)
	req.Param("RequestData", requestData)
	req.Param("EBusinessID", eBusinessID)
	req.Param("RequestType", requestType)
	req.Param("DataSign", dataSign)
	req.Param("DataType", "2")
	
	_, err := req.String()
	if err != nil {
		panic(err)
	}

	req.ToJSON(&result)

	return result
}