package controllers

type AdminController struct {
	baseController
}

func (p *AdminController) Home()  {
	p.TplName = "admin/index.html"
}

func (p *AdminController) Logout()  {
	p.DestroySession()
	p.History("退出登录", "/login")
}
