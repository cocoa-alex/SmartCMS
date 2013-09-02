package models

import (
	//"websystem/utils"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

/*
*table stucts
*author:steve
 */
type Userinfo struct {
	Uid      int    `orm:"auto"`
	Username string `orm:"size(100)"`
	Password string `orm:"size(100)"`
	Created  time.Time
}

/*table end*/

func init() {
	orm.RegisterModel(new(Userinfo))
	orm.RegisterDataBase("default", "mysql", "root:mima123@/mywebsite?charset=utf8", 30)
}

/**
user table
**/
func InsertUsr(usr Userinfo) Userinfo {
	db := orm.NewOrm()
	db.Insert(&usr)
	return usr
}
func FindUserById(id int) Userinfo {
	var one Userinfo
	db := orm.NewOrm()
	db.QueryTable("userinfo").Filter("uid", id).One(&one)
	return one
}

func FindUsrByName(name string) Userinfo {
	var usr Userinfo
	db := orm.NewOrm()
	db.QueryTable("userinfo").Filter("username", name).One(&usr)
	return usr
}
func FindAllUsrs() []*Userinfo {
	var usrs []*Userinfo
	db := orm.NewOrm()
	db.QueryTable("userinfo").All(&usrs)
	return usrs
}
func UsrPageing(offset int64, limit int) []*Userinfo {
	var users []*Userinfo
	db := orm.NewOrm()
	db.QueryTable("userinfo").Limit(limit).Offset(offset).All(&users)
	return users
}
