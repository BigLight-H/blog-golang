package controllers

import (
	"beego-demo/models"
	"beego-demo/util"
	"github.com/davecgh/go-spew/spew"
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
	p.Data["c_id"] = id
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

//添加评论
func (p *HomeController) AddComment()  {
	c_id := p.GetString("cid")
	aid, _ :=strconv.Atoi(c_id)
	uid := p.GetSession("client_id").(int)
	if uid > 0 {
		comment := p.GetString("comments")
		if comment == "" {
			p.MsgBack("评论不能为空", 0)
		} else {
			ment := models.Comment{}
			user_id := models.Client{Id:uid}
			a_id := models.Article{Id:aid}
			ment.Client = &user_id
			ment.Article = &a_id
			ment.Content = comment
			ment.Created = util.TimeSet()
			_, err := p.o.Insert(&ment)
			if err != nil {
				p.MsgBack("评论失败", 0)
			}
			article := models.Article{Id:aid}
			p.o.Read(&article)
			spew.Dump(article.CommentNum)
			article.CommentNum = article.CommentNum + 1
			p.o.Update(&article,"CommentNum")
			p.MsgBack("评论成功", 1)
		}
	} else {
		p.MsgBack("登录后可评论", 0)
	}


}

//给我留言
func (p *HomeController) SetMessage() {
	if p.Ctx.Request.Method == "POST"{
		name := p.GetString("name")
		email := p.GetString("email")
		content := p.GetString("content")
		if name == "" {
			p.MsgBack("名字不能为空", 0)
		} else if p.VerifyEmailFormat(email) != true {
			p.MsgBack("邮箱格式不对", 0)
		} else if len(content) < 6 {
			p.MsgBack("内容太短", 0)
		} else {
			feed := models.FeedBack{}
			feed.Name = name
			feed.Content = content
			feed.Email = email
			feed.Created = util.TimeSet()
			_, err :=p.o.Insert(&feed)
			if err != nil {
				p.MsgBack("留言失败", 0)
			}
			p.MsgBack("留言成功,看到后第一时间回复您", 1)
		}
	}else {
		p.TplName = "home/contact.html"
	}
}

//关于我自己
func (p *HomeController) About() {
	about := models.About{Id:1}
	p.o.Read(&about)
	spew.Dump(about)
	p.Data["abouts"] = about
	p.TplName = "home/about.html"
}

//添加文章
func (p *HomeController) AddArticle() {

	p.TplName = "home/add-article.html"
}


//tag搜索
func (p *HomeController) SearchTag()  {

}

//文章搜索
func (p *HomeController) Search()  {

}