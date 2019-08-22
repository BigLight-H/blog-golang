package controllers

import (
	"beego-demo/models"
	"beego-demo/util"
	"strings"
)

type LoginController struct {
	baseController
}

func (this *LoginController) Login()	 {
	if this.Ctx.Request.Method == "POST" {
		username := this.GetString("username")
		password := this.GetString("password")
		user := models.User{Username:username}
		this.o.Read(&user,"username")
		if user.Password == "" {
			this.MsgBack("账号不存在",0)
		}

		if util.Md5(password) != strings.Trim(user.Password, " ") {
			this.MsgBack("密码错误", 0)
		} else {
			this.Redirect("/admin/home",302)
		}
		this.SetSession("user", user)
	} else {
		this.TplName = "admin/login.html"
	}

}
