package controllers

import (
	"beego-demo/models"
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
			p.History("", "/error")
		}
	}

	permissions := [] *models.Permissions{}
	p.o.QueryTable(new(models.Permissions).TableName()).All(&permissions)
	p.Data["sidebar"] = &permissions
	p.Data["user"] = p.GetSession("user")
	p.Data["tag"] = "Admin"
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
