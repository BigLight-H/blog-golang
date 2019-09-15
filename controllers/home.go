package controllers

import (
	"beego-demo/models"
	"beego-demo/util"
	"github.com/davecgh/go-spew/spew"
	"math"
	"strconv"
	"strings"
)

type HomeController struct {
	baseController
}

const Number int = 3  //每页容量
const Number_ float64 = 3.0  //每页容量


//前台首页面
func (p *HomeController) Index() {
	page := p.Ctx.Input.Param(":page")
	p.article(0, page)
	p.TplName = "home/index.html"
}

//文章列表
func (p *HomeController) article(id int, page string) {
	pgs, _ := strconv.Atoi(page)
	//分页
	pages := ( pgs - 1 ) * Number
	if pages < 1 {
		pages = 1
	}
	article := []*models.Article{}
	qs := p.o.QueryTable(new(models.Article).TableName())
	qs = qs.Filter("status", 1)
	if id > 0 {
		qs = qs.Filter("client_id", id)
	}
	num, _ := qs.Count()
	//总条数
	p.Data["all_num"] = num
	//总页数
	num_str := strconv.FormatInt(num,10)
	n, _ := strconv.ParseFloat(num_str, 64)
	p.Data["total"] = math.Ceil(n / Number_)
	//上一页
	p.Data["previous_page"] = pgs - 1
	//下一页
	p.Data["next_page"] = pgs + 1
	//当前页
	p.Data["current_page"],_ = strconv.ParseFloat(page, 64)
	qs.OrderBy("-id").RelatedSel().Limit(Number, pages).All(&article)
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
	ment.RelatedSel().OrderBy("path").All(&comment)
	p.Data["comment"] = comment
	//遍历出文章标签
	tags := make(map[int]string)
	for _, v := range article {
		for _, tid := range strings.Split(v.Tags, ",") {
			tag_id, _ := strconv.Atoi(tid)
			tag := models.Tags{Id:tag_id}
			p.o.Read(&tag)
			tags[tag.Id] = tag.TagName
		}
	}
	p.starKeep(id)
	p.Data["tags"] = tags
	p.Data["c_id"] = id
	p.TplName = "home/detail.html"
}

//文章详情里面的喜欢收藏标签点亮
func (p *HomeController) starKeep(id int)  {
	client_id_ := p.GetSession("client_id")
	if client_id_ != nil {
		client_id := p.GetSession("client_id").(int)
		p.Data["keep_num"], _ = p.o.QueryTable(new(models.Collection).TableName()).Filter("client_id", client_id).Filter("article_id", id).Count()
		p.Data["star_num"], _ = p.o.QueryTable(new(models.Zan).TableName()).Filter("client_id", client_id).Filter("article_id", id).Count()
		spew.Dump(p.Data["keep_num"], p.Data["star_num"])
	}
}


//作者介绍
func (p *HomeController) Author() {
	cid := p.Ctx.Input.Param(":id")
	page := p.Ctx.Input.Param(":page")
	id, _ := strconv.Atoi(cid)
	user := []*models.Client{}
	p.o.QueryTable(new(models.Client).TableName()).Filter("id", id).RelatedSel().All(&user)
	p.Data["client_user"] = user
	p.article(id, page)
	p.Data["cid"] = id
	p.TplName = "home/author.html"
}

//添加评论
func (p *HomeController) AddComment()  {
	c_id := p.GetString("cid")
	path := p.GetString("path")
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
			if path != "" {
				ment.IsReply = 1
			}
			ment.Created = util.TimeSet()
			mid, err := p.o.Insert(&ment)
			if err != nil {
				p.MsgBack("评论失败", 0)
			}
			article := models.Article{Id:aid}
			p.o.Read(&article)
			article.CommentNum = article.CommentNum + 1
			p.o.Update(&article,"CommentNum")
			//存入评论path
			comm := models.Comment{Id:mid}
			if path != "" {
				comm.Path = path+","+ strconv.FormatInt(mid,10)
			} else {
				comm.Path = strconv.FormatInt(mid,10)
			}
			p.o.Update(&comm,"Path")
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
	p.Data["abouts"] = about
	p.TplName = "home/about.html"
}

//错误页面
func (p *HomeController) HomeError() {
	p.TplName = "home/404.html"
}


//tag搜索
func (p *HomeController) SearchTag()  {
	str := p.Ctx.Input.Param(":str")
	tid, _ :=strconv.Atoi(str)
	tag_model := models.Tags{Id:tid}
	p.o.Read(&tag_model)
	p.Data["search"] = tag_model.TagName
	p.sou(str,"tags")
	p.Data["title"] = "标签搜索"
	p.TplName = "home/search.html"
}

//文章搜索
func (p *HomeController) Search()  {
	str := p.Ctx.Input.Param(":str")
	p.sou(str,"title")
	p.Data["title"] = "文章搜索"
	p.Data["search"] = str
	p.TplName = "home/search.html"
}

//搜索文章方法
func (p *HomeController) sou(str string, types string) {
	article := []*models.Article{}
	qs := p.o.QueryTable(new(models.Article).TableName())
	qs = qs.Filter(types+"__icontains", str)
	qs.RelatedSel().All(&article)
	p.Data["article_search"] = article
}


//收藏的文章
func (p *HomeController) Keep() {
	client_id := p.GetSession("client_id").(int)
	collect := []*models.Collection{}
	p.o.QueryTable(new(models.Collection).TableName()).Filter("client_id", client_id).RelatedSel().All(&collect)
	p.Data["article_search"] = collect
	p.Data["title"] = "收藏的文章"
	p.Data["search"] = "收藏文章"
	p.TplName = "home/personal.html"
}

//点赞的文章
func (p *HomeController) Zan() {
	client_id := p.GetSession("client_id").(int)
	zan := []*models.Zan{}
	p.o.QueryTable(new(models.Zan).TableName()).Filter("client_id", client_id).RelatedSel().All(&zan)
	p.Data["article_search"] = zan
	p.Data["title"] = "喜欢的文章"
	p.Data["search"] = "喜欢文章"
	p.TplName = "home/personal.html"
}