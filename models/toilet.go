package models

import (
	"fmt"

	"github.com/astaxie/beego/httplib"
)

func SearchToilet(location string) (result map[string]interface{}) {
	fmt.Printf(location)
	//生成要访问的url
	url := "http://restapi.amap.com/v3/place/around?key=2749361c6fa6a5c850dd426c9d07827d&keywords=厕所&radius=2000&location=" + location
	//提交请求
	req := httplib.Get(url)
	
	str, err := req.String()
	if err != nil {
		panic(err)
	}

	fmt.Printf(url)

	fmt.Printf(str)

	req.ToJSON(&result)

	return result
}