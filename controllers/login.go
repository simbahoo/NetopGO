package controllers

import (
	"NetopGO/models"
	//"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {

	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	uname := this.Input().Get("uname")
	passwd := this.Input().Get("passwd")
	encodePasswd := string(models.Base64Encode([]byte(passwd)))
	user, err := models.Login(uname)
	if err != nil || encodePasswd != user.Passwd {
		beego.Error(err)
		this.Data["Error"] = true
		this.TplName = "login.html"
		return
	}

	this.SetSession("id", user.Id)
	this.SetSession("uname", user.Name)
	this.SetSession("passwd", user.Passwd)
	this.SetSession("auth", user.Auth)

	this.Redirect("/", 302)
}

func (this *LoginController) Logout() {
	this.DelSession("uname")
	this.DelSession("passwd")
	this.DelSession("auth")
	this.TplName = "login.html"
	return
}
