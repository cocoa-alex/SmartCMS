package admin

import (
	"github.com/astaxie/beego"
)

type IndexControl struct {
	beego.Controller
}

func (u *IndexControl) Index() {
	u.Layout = "admin/layout.html"
	u.TplNames = "admin/index.tpl"
	u.Render()
}
