package main

import (
	"SmartCMS/controllers/admin"
	"SmartCMS/utils"
	"github.com/astaxie/beego"
)

func main() {
	//main layout
	beego.Router("/", &admin.IndexControl{}, "get:Index")
	//users
	beego.Router("/admin/usr/list", &admin.UserControl{}, "*:ListUsers")
	beego.Router("/admin/usr/add", &admin.UserControl{}, "*:CreateUsr")

	//temeplate fucation
	beego.AddFuncMap("date2string", utils.Format2string)
	beego.SessionOn = true
	beego.AutoRender = false
	beego.SetStaticPath("/upload", "upload")
	beego.Run()
}
