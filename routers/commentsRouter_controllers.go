package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["gameLog/controllers:LogController"] = append(beego.GlobalControllerRouter["gameLog/controllers:LogController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gameLog/controllers:LogController"] = append(beego.GlobalControllerRouter["gameLog/controllers:LogController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gameLog/controllers:LogController"] = append(beego.GlobalControllerRouter["gameLog/controllers:LogController"],
		beego.ControllerComments{
			Method: "Client",
			Router: `/client`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gameLog/controllers:LogController"] = append(beego.GlobalControllerRouter["gameLog/controllers:LogController"],
		beego.ControllerComments{
			Method: "Server",
			Router: `/server`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
