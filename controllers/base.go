package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

type baseController struct {
	beego.Controller
	o orm.Ormer
	controllerName string
	actionName     string
}

type Json struct {
	Msg string
	Status int
}

func (p *baseController) Prepare()  {
	controllerName, actionName := p.GetControllerAndAction()
	p.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	p.actionName = strings.ToLower(actionName)
	p.o = orm.NewOrm()
	if p.controllerName == "admin"{
		if p.GetSession("user") == nil {
			p.TplName = "admin/500.html"
		}
	}
}

func (p *baseController) History(msg string, url string)  {
	if url == "" {
		p.Ctx.WriteString("<script>alert('"+msg+"');window.history.go(-1);</script>")
		p.StopRun()
	} else {
		p.Redirect(url,302)
	}
}

func (p *baseController) MsgBack(msg string, status int)  {
	data := &Json{
		msg,
		status,
	}
	p.Data["json"] = data
	p.ServeJSON()
}
