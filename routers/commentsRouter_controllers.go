package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["zjlife/controllers:ExpressController"] = append(beego.GlobalControllerRouter["zjlife/controllers:ExpressController"],
		beego.ControllerComments{
			Method: "EbusinessOrderHandle",
			Router: `/ebusinessOrderHandle`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zjlife/controllers:ToiletController"] = append(beego.GlobalControllerRouter["zjlife/controllers:ToiletController"],
		beego.ControllerComments{
			Method: "SearchToilet",
			Router: `/around`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zjlife/controllers:UserController"] = append(beego.GlobalControllerRouter["zjlife/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUserInfo",
			Router: `/info`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zjlife/controllers:UserController"] = append(beego.GlobalControllerRouter["zjlife/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zjlife/controllers:UserController"] = append(beego.GlobalControllerRouter["zjlife/controllers:UserController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["zjlife/controllers:WeatherController"] = append(beego.GlobalControllerRouter["zjlife/controllers:WeatherController"],
		beego.ControllerComments{
			Method: "SearchWeather",
			Router: `/weatherInfo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
