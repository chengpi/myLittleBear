package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["myLittleBear/controllers:MainController"] = append(beego.GlobalControllerRouter["myLittleBear/controllers:MainController"],
        beego.ControllerComments{
            Method: "Test1",
            Router: `/test1`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["myLittleBear/controllers:MainController"] = append(beego.GlobalControllerRouter["myLittleBear/controllers:MainController"],
        beego.ControllerComments{
            Method: "Test2",
            Router: `/test2`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
