package controllers

import (
	"beego-demo/models"
	"strconv"
	"strings"
)

type HomeController struct {
	baseController
}

//前台首页面
func (p *HomeController) Index() {
	p.article(0)
	p.TplName = "home/index.html"
}

//文章列表
func (p *HomeController) article(id int) {
	article := []*models.Article{}
	qs := p.o.QueryTable(new(models.Article).TableName())
	qs = qs.Filter("status", 1)
	if id > 0 {
		qs = qs.Filter("client_id", id)
	}
	qs.OrderBy("-id").RelatedSel().All(&article)
	p.Data["articles"] = article
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
	//文章评论
	comment := []*models.Comment{}
	ment := p.o.QueryTable(new(models.Comment).TableName())
	ment = ment.Filter("article_id", id)
	count, _ := ment.Count()
	//文章数量
	p.Data["num"] = count
	ment.RelatedSel().All(&comment)
	p.Data["comment"] = comment
	//遍历出文章标签
	tags := make(map[int]string)
	for _, v := range article {
		for k, tid := range strings.Split(v.Tags, ",") {
			tag_id, _ := strconv.Atoi(tid)
			tag := models.Tags{Id:tag_id}
			p.o.Read(&tag)
			tags[k] = tag.TagName
		}
	}
	p.Data["tags"] = tags
	p.TplName = "home/detail.html"
}

//作者介绍
func (p *HomeController) Author() {
	cid := p.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(cid)
	user := []*models.Client{}
	p.o.QueryTable(new(models.Client).TableName()).Filter("id", id).RelatedSel().All(&user)
	p.Data["client_user"] = user
	p.article(id)
	p.TplName = "home/author.html"
}


//tag搜索
func (p *HomeController) SearchTag()  {

}

//文章搜索
func (p *HomeController) Search()  {

}