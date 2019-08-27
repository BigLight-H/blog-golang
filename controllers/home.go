package controllers

import (
	"beego-demo/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type HomeController struct {
	beego.Controller
	o orm.Ormer
}

//前台首页面
func (p *HomeController) Index() {

}

//文章列表
func (p *HomeController) article() {
	article := []*models.Article{}
	p.o.QueryTable(new(models.Article).TableName()).OrderBy("-id").All(&article)
	p.Data["articles"] = article
}