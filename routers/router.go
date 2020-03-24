package routers

import (
	"myLittleBear/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"get:Get")
    beego.Router("/login",&controllers.MainController{},"get:LoginGet;post:LoginPost")

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/test1",
			beego.NSInclude(
				&controllers.MainController{},
			),
		),
		beego.NSNamespace("/test2",
			beego.NSInclude(
				&controllers.MainController{},
			),
		),
		beego.NSNamespace("/test3",
			beego.NSInclude(
				&controllers.MainController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
