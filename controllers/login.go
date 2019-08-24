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
			data := models.User{}
			qs := this.o.QueryTable(new(models.User).TableName())
			qs.Filter("id", user.Id).RelatedSel().All(&data)
			this.SetSession("user", data)
			this.MsgBack("登录成功", 1)
		}
	} else {
		this.Data["tag"] = "Login"
		this.TplName = "admin/login.html"
	}

}
