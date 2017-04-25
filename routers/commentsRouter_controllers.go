package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["zjlife/controllers:UserController"] = append(beego.GlobalControllerRouter["zjlife/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["zjlife/controllers:UserController"] = append(beego.GlobalControllerRouter["zjlife/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["zjlife/controllers:UserController"] = append(beego.GlobalControllerRouter["zjlife/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["zjlife/controllers:UserController"] = append(beego.GlobalControllerRouter["zjlife/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["zjlife/controllers:UserController"] = append(beego.GlobalControllerRouter["zjlife/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["zjlife/controllers:UserController"] = append(beego.GlobalControllerRouter["zjlife/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["zjlife/controllers:UserController"] = append(beego.GlobalControllerRouter["zjlife/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["zjlife/controllers:WeatherController"] = append(beego.GlobalControllerRouter["zjlife/controllers:WeatherController"],
		beego.ControllerComments{
			Method: "SearchWeather",
			Router: `/weatherInfo`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
