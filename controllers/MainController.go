package controllers

import (
	_"github.com/astaxie/beego"
	"fmt"
	"myLittleBear/models"
	"time"
	"myLittleBear/utils"
	//"encoding/json"
)

type MainController struct {
	//beego.Controller
	BaseController
}
type User struct {
	Id       string    `json:"id"`
	Username string    `json:"username"`
}

func (c *MainController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
	c.TplName = "index.html"
}
func (c *MainController) LoginGet() {
	c.TplName = "login.html"
}
func (c *MainController) LoginPost() {
	//获取表单信息
	username := c.GetString("user_name")
	password := c.GetString("password")
	repassword := c.GetString("repassword")
	fmt.Println(username, password, repassword)
	//注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误
	num,err := models.SelectUserByUserName(username)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code":0,"message":"查询错误"}
		c.ServeJSON()
		return
	}
	fmt.Println("num:",num)
	if num > 0 {
		c.Data["json"] = map[string]interface{}{"code":0,"message":"用户名已经存在"}
		c.ServeJSON()
		return
	}

	password = utils.MD5(password)
	fmt.Println("md5后：",password)

	user := models.User{
		UserName:username,
		Password:password,
		Status:"1",
		CreateTime:time.Now().Format("2006-01-02 15:04:05"),
		}
	_,err =models.InsertUser(user)
	if err != nil{
		c.Data["json"] = map[string]interface{}{"code":0,"message":"注册失败"}
	}else{
		c.Data["json"] = map[string]interface{}{"code":1,"message":"注册成功"}
	}
	c.ServeJSON()
}

// @Title test1
// @Description test1
// @Param	id		path 	string		true	"id"
// @Success 200 {object} model.RespResult {"id":id}
// @Failure 403 body is empty
// @router /test1 [get]
func (c *MainController) Test1()  {
	var id = c.Ctx.Input.Param(":id")
	c.Data["json"] = map[string]interface{}{"id":id}
	c.ServeJSON()

}

// @Title test2
// @Description test2
// @Param	id		       query 	string		true	"id"
// @Param	username	   query 	string		true	"username"
// @Success 200 {object} model.RespResult {"id":id,"username":username}
// @Failure 403 body is empty
// @router /test2 [get]
func (c *MainController) Test2()  {
	var params = c.Ctx.Input.Params()
	id := params["id"]
	username := params["username"]
	fmt.Println("id:", id, "username:", username)

	//var user User = User{Id:"0",Username:"yang"}
	//err := json.Unmarshal(c.Ctx.Input.RequestBody,&user)
	//if err != nil {
	//	c.Data["json"] = map[string]interface{}{"err":err}
	//	c.ServeJSON()
	//	return
	//}

	//c.Data["json"] = map[string]interface{}{"id":user.Id,"username":user.Username}

	//下面这样写是可以的
	//id = c.GetString("id")
	//username = c.GetString("username")

	id = c.Ctx.Input.Query("id")
	username = c.Ctx.Input.Query("username")
	fmt.Println(id,username)

	c.Data["json"] = map[string]interface{}{"id":id,"username":username}

	c.ServeJSON()
}

// @Title test3
// @Description test3
// @Param	id		       query 	string		true	"id"
// @Param	username	   query 	string		true	"username"
// @Success 200 {object} model.RespResult {"id":id,"username":username}
// @Failure 403 body is empty
// @router /test3 [get]
func (c *MainController) Test3() {
	id := c.Ctx.Input.Query("id")
	username := c.Ctx.Input.Query("username")
	fmt.Println(id,username)

	c.Data["json"] = map[string]interface{}{"id":id,"username":username}

	c.ServeJSON()
}