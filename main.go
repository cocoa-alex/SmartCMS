package main

import (
	"github.com/astaxie/beego"
	//"websystem/api"
	"SmartCMS/controllers/admin"
	"SmartCMS/utils"
)

func main() {
	beego.Router("/admin", &admin.UserControl{}, "*:ListUsers")

	//temeplate fucation
	beego.AddFuncMap("date2string", utils.Format2string)
	beego.SessionOn = true
	beego.AutoRender = false
	beego.SetStaticPath("/upload", "upload")
	beego.Run()
}
