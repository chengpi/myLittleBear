package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type User struct {
	Id          int     `orm:"column(id);pk;auto"json:"id" description:"主键"`
	UserName    string  `orm:"column(user_name);"json:"user_name" description:"用户名"`
	Password    string  `orm:"column(password);"json:"password" description:"密码"`
	Status      string  `orm:"column(status);"json:"status" description:"状态"`
	CreateTime  string  `orm:"column(create_time);"json:"create_time" description:"创建注册时间"`
}

func init() {
	orm.RegisterModel(new(User))
}
func (r *User) TableName() string {
	return "users"
}
func InsertUser(user User)(string, error) {
	o := orm.NewOrm()
	err := o.Begin()
	//num,err := SelectUserById(user.Id)
	//if err == nil && num > 0 {
	//	return "已存在该用户", nil
	//}
	id, err := o.Insert(&user)
	if err != nil {
		err = o.Rollback()
		return "插入失败", err
	}
	fmt.Println("id:",id)

	err = o.Commit()
	return "插入成功", nil

}
func SelectUserById(id int)(int64, error) {
	o:=orm.NewOrm()
	num,err := o.Raw("select * from users where id=?", id).QueryRows()
	fmt.Println(num)
	if err != nil {
		fmt.Println("查询失败:",err)
		return 0, err
	}
	return num, nil
}
func SelectUserByUserName(user_name string)(int64, error) {
	o:=orm.NewOrm()
	var maps []orm.Params
	num,err := o.Raw("select * from users where user_name=?", user_name).Values(&maps)
	fmt.Println(num)
	if err != nil {
		fmt.Println("查询失败:",err)
		return 0, err
	}
	return num, nil
}
