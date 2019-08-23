package controllers

import (
	"beego-demo/models"
)

type AdminController struct {
	baseController
}

func (p *AdminController) Home()  {
	p.TplName = "admin/index.html"
}

func (p *AdminController) Sidebar()  {
	permissions := [] *models.Permissions{}
	p.o.QueryTable( new(models.Permissions).TableName() ).All(&permissions)
	p.Data["json"] = permissions
	p.ServeJSON()
}

func (p *AdminController) Logout()  {
	p.DestroySession()
	p.History("退出登录", "/login")
}
