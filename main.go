package main

import (
	_ "myLittleBear/routers"
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
)
func init() {
	path := beego.AppConfig.String("mysqluser") + ":" +
		beego.AppConfig.String("mysqlpass") + "@tcp(" + beego.AppConfig.String("mysqlhost") + ":" +
		beego.AppConfig.String("mysqlport") + ")/" + beego.AppConfig.String("mysqldb") + "?charset=utf8"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", path, maxIdle, maxConn)
	fmt.Println("****")
}

func main() {
	beego.SetStaticPath("/static", "static")
	beego.Run()
}

