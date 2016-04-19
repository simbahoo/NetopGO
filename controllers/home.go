package controllers

import (
//"github.com/astaxie/beego"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	uid, uname, role := this.IsLogined()
	this.Data["Id"] = uid
	this.Data["Uname"] = uname
	this.Data["Role"] = role
	this.TplName = "index.html"
}
