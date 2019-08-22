package controllers

type AdminController struct {
	baseController
}

func (p *AdminController) Home()  {
	p.TplName = "admin/index.html"
}