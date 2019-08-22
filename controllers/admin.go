package controllers

type AdminController struct {
	baseController
}

func (p *AdminController) Home()  {
	if p.GetSession("user") == nil {
		p.Redirect("/login",302)
	}
}