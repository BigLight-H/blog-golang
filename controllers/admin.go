package controllers

import (
	"beego-demo/models"
	"strconv"
)

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

func (p *AdminController) ClassifyAdd() {
	if p.Ctx.Request.Method == "POST" {
		types := models.Type{}
		pid := p.Input().Get("pid")
		name := p.Input().Get("name")
		id  := p.Input().Get("id")
		if id != "" {
			id, _ := strconv.Atoi(id)
			types.Id = id
			types.Pid = pid
			types.TName = name
			types.Status = 1
			if _, err := p.o.Update(&types); err != nil {
				p.History("", "/admin/classify/add/"+strconv.Itoa(id))
			} else {
				p.History("", "/admin/classify/list")
			}
		} else {
			types.Pid = pid
			types.TName = name
			types.Status = 1
			_, err := p.o.Insert(&types)
			if err == nil {
				p.History("", "/admin/classify/list")
			} else {
				p.TplName = "admin/classify_add.html"
			}
		}
	} else {
		id := p.Ctx.Input.Param(":id")
		class := []*models.Type{}
		qs := p.o.QueryTable(new(models.Type).TableName())
		if id != "" {
			qs.Filter("status",1).FilterRaw("id", "!="+id).All(&class)
			p.Data["class_type"] = class
			qs.Filter("id", id).One(&class)
			p.Data["type_data"] = class
		} else {
			qs.Filter("status",1).All(&class)
			p.Data["class_type"] = class
		}
		p.TplName = "admin/classify_add.html"
	}
}

func (p *AdminController) ClassifyList() {
	if p.Ctx.Request.Method == "POST"{
		class := models.Type{}
		p.o.QueryTable(new(models.Type).TableName())
		ids := p.Input().Get("id")
		id, _ := strconv.Atoi(ids)
		class.Status = 0
		class.Id = id
		class.Pid = p.Input().Get("pid")
		class.TName = p.Input().Get("name")
		if _, err := p.o.Update(&class); err == nil {
			p.MsgBack("删除成功", 1)
		} else {
			p.MsgBack("删除失败", 0)
		}
	} else {
		class := []*models.Type{}
		qs := p.o.QueryTable(new(models.Type).TableName())
		qs.Filter("status",1).All(&class)
		p.Data["types"] = class
		p.TplName = "admin/classify_list.html"
	}
}
