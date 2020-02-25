package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

type ExtendedController struct {
	beego.Controller
	Session session.Store
}

func (this *ExtendedController) Prepare() {
	this.Session = this.StartSession()
}
