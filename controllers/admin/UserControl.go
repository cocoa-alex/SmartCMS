package admin

import (
	"SmartCMS/models"
	"github.com/astaxie/beego"
)

type UserControl struct {
	beego.Controller
}

func (u *UserControl) ListUsers() {
	u.Layout = "admin/layout.html"
	u.TplNames = "admin/user.tpl"
	users := models.FindAllUsrs()
	u.Data["items"] = users
	u.Render()
}

func (u *UserControl) CreateUsr() {
	u.Layout = "admin/layout.html"
	u.TplNames = "admin/add_user.tpl"
	u.Render()
}
