package controllers

import (
	"beego-demo/models"
	"github.com/davecgh/go-spew/spew"
)

type HomeController struct {
	baseController
}

//前台首页面
func (p *HomeController) Index() {
	p.article()
	p.TplName = "home/index.html"
}

//文章列表
func (p *HomeController) article() {
	article := []*models.Article{}
	p.o.QueryTable(new(models.Article).TableName()).OrderBy("-id").RelatedSel().All(&article)
	p.Data["articles"] = article
	spew.Dump(article)
}