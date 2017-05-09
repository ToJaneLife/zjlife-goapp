package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["zjlife/controllers:ExpressController"] = append(beego.GlobalControllerRouter["zjlife/controllers:ExpressController"],
		beego.ControllerComments{
			Method: "EbusinessOrderHandle",
			Router: `/ebusinessOrderHandle`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["zjlife/controllers:WeatherController"] = append(beego.GlobalControllerRouter["zjlife/controllers:WeatherController"],
		beego.ControllerComments{
			Method: "SearchWeather",
			Router: `/weatherInfo`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
