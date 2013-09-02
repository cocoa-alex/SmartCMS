package libs

import (
	"github.com/astaxie/beego"
)

var (
	session_name string
)

type BaseController struct {
	beego.Controller
}

type RootController struct {
	BaseController
}

//相当于初始化 检测是否有session
func (this *BaseController) Prepare() {
	session_name, _ = this.GetSession("session_name").(string)
	if session_name == "" {
		this.Data["usrname"] = ""
	} else {
		this.Data["usrname"] = session_name
	}
}

func (this *RootController) Prepare() {
	this.BaseController.Prepare()
	if session_name == "" {
		this.Ctx.Redirect(302, "/root/login")
	}
}
