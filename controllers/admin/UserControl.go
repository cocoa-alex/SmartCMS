package controllers

import (
	"github.com/astaxie/beego"
)

type UserControl struct {
	beego.Controller
}

func (u *UserControl) ListUsers() {
	u.Layout = "layout.html"
	u.TplNames = "index.tpl"
}
