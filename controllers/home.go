package controllers

import (
	"beego-demo/models"
	"github.com/davecgh/go-spew/spew"
	"strconv"
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
	qs := p.o.QueryTable(new(models.Article).TableName())
	qs = qs.Filter("status", 1)
	qs.OrderBy("-id").RelatedSel().All(&article)
	p.Data["articles"] = article
	qs.OrderBy("-click_volume").Limit(3).All(&article)
	p.Data["hot"] = article//热门文章
}

//文章详情
func (p *HomeController) Detail() {
	id_ := p.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(id_)
	article := []*models.Article{}
	qs := p.o.QueryTable(new(models.Article).TableName())
	//文章详情
	qs.Filter("id", id).RelatedSel().One(&article)
	p.Data["article"] = article
	spew.Dump(article)
	//热门文章3篇
	qs.Filter("status", 1).OrderBy("-click_volume").Limit(3).All(&article)
	p.Data["hot"] = article
	//文章评论
	comment := []*models.Comment{}
	ment := p.o.QueryTable(new(models.Comment).TableName())
	ment = ment.Filter("article_id", id)
	count, _ := ment.Count()
	//文章数量
	p.Data["num"] = count
	ment.RelatedSel().All(&comment)
	p.Data["comment"] = comment
	p.TplName = "home/detail.html"
}

//tag搜索
func (p *HomeController) SearchTag()  {

}

//文章搜索
func (p *HomeController) Search()  {

}