package controllers

import (
	"beego-demo/models"
	"beego-demo/util"
	"github.com/davecgh/go-spew/spew"
	"strconv"
	"strings"
)

type LoginController struct {
	baseController
}

//后台登录
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
			this.SetSession("user_id", data.Id)
			this.MsgBack("登录成功", 1)
		}
	} else {
		this.Data["tag"] = "Login"
		this.TplName = "admin/login.html"
	}

}

//前台登陆
func (this *LoginController) HomeLogin() {
	email := this.GetString("email")
	password := this.GetString("pwd")
	user := models.Client{Email:email}
	spew.Dump(user)
	this.o.Read(&user,"Email")
	if user.Password == "" {
		this.MsgBack("账号不存在",0)
	}

	if util.Md5(password) != strings.Trim(user.Password, " ") {
		this.MsgBack("密码错误", 0)
	} else {
		data := models.Client{}
		qs := this.o.QueryTable(new(models.Client).TableName())
		qs.Filter("id", user.Id).RelatedSel().All(&data)
		this.SetSession("client", data)
		this.SetSession("client_id", data.Id)
		this.MsgBack("登录成功", 1)
	}
}

//注册前台用户
func (this *LoginController) Register()  {
	email := this.GetString("email")
	username := this.GetString("uname")
	mobile := this.GetString("mobile")
	password := this.GetString("pwd")
	password1 := this.GetString("pwd_")
	id_ := this.GetString("id")
	id, _ := strconv.Atoi(id_)
	if password != password1 {
		this.MsgBack("两次密码不一致", 0)
	} else if len(username) < 2 {
		this.MsgBack("用户名长度不能小于2", 0)
	} else if len(password) < 6 {
		this.MsgBack("用户名长度不能小于6", 0)
	}else if this.VerifyEmailFormat(email) == false {
		this.MsgBack("邮箱号码格式不对", 0)
	}else if this.VerifyMobileFormat(mobile) == false {
		this.MsgBack("手机号码格式不对", 0)
	} else {
		data := models.Client{}
		data.Email = email
		data.Mobile = mobile
		data.Password = util.Md5(password)
		data.Username = username
		data.Sex = 1
		data.Pic = ""
		img := models.HeadImg{Id:id}
		data.HeadImg = &img
		_, err := this.o.Insert(&data)
		if err != nil {
			spew.Dump(err)
			this.MsgBack("注册失败", 0)
		}
		this.MsgBack("注册成功", 1)
	}
}

//找回密码
func (this *LoginController) Retrieve() {
	email := this.GetString("email")
	if this.VerifyEmailFormat(email) == false {
		this.MsgBack("邮箱号码格式不对", 0)
	} else {
		user := models.Client{Email:email}
		spew.Dump(user)
		this.o.Read(&user,"Email")
		if user.Password == "" {
			this.MsgBack("账号不存在",0)
		} else {
			//查看邮件验证是否存在
			//发送邮件验证码
			//对比邮件验证码
		}
	}
}