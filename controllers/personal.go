/*
@Time : 2019-09-05 23:40
@Author : biglight
@File : personal
@Software: GoLand
*/
package controllers

import (
	"axicoo/models"
	"axicoo/util"
	"github.com/astaxie/beego/orm"
	"os"
	"strconv"
	"strings"
)

type PersonalController struct {
	baseController
}

//个人中心
func (p *PersonalController) PersonalCenter()  {
	p.getPerData()
	p.TplName = "personal/index.html"
}

//添加文章
func (p *PersonalController) AddArticle() {
	id := p.GetSession("client_id").(int)
	if p.Ctx.Request.Method == "POST"{
		aid_ := p.GetString("aid")
		aid, _ := strconv.Atoi(aid_)
		title := p.GetString("title")
		logo := p.GetString("logo")
		content := p.GetString("content")
		types := p.GetString("select")
		tags := p.GetString("tags")
		t_id, _ := strconv.Atoi(types)
		desc := p.GetString("desc")
		var str string
		var code int
		if aid_ != "" {
			article := models.Article{Id:aid}
			p.o.Read(&article)
			article.Title = title
			article.Picture = logo
			article.Content = content
			article.Description = desc
			article.Picture = logo
			article.Tags = tags
			tid := models.Type{Id:t_id}
			article.Type = &tid
			article.Title = title
			article.Created = util.TimeSet()
			_, err := p.o.Update(&article, "Title","Picture","Description","Type","Tags","Content")
			if err != nil {
				str = "文章编辑失败"
				code = 0
			} else {
				str = "文章编辑成功"
				code = 1
			}
		} else {
			article := models.Article{}
			article.Title = title
			article.Picture = logo
			article.Content = content
			article.Status = 0
			article.Review = 0
			article.CommentNum = 0
			article.ClickVolume = 0
			article.Description = desc
			article.Picture = logo
			article.Tags = tags
			c_id := models.Client{Id:id}
			tid := models.Type{Id:t_id}
			article.Client = &c_id
			article.Type = &tid
			article.Title = title
			article.Created = util.TimeSet()
			_, err := p.o.Insert(&article)
			if err != nil {
				str = "文章添加失败"
				code = 0
			} else {
				str = "文章添加成功"
				code = 1
			}
		}
		p.MsgBack(str, code)
	} else {
		aid_ := p.Ctx.Input.Param(":id")
		aid, _ := strconv.Atoi(aid_)
		types := []*models.Type{}
		tags := []*models.Tags{}
		qs := p.o.QueryTable(new(models.Type).TableName()).Filter("status",1).FilterRaw("pid", "!=''")
		qt := p.o.QueryTable(new(models.Tags).TableName())
		if aid_ != "" {
			as := models.Article{Id:aid}
			p.o.Read(&as)
			tid := strconv.Itoa(as.Type.Id)
			qs.FilterRaw("id", "!="+tid).All(&types)
			p.Data["a_types"] = types
			qs.Filter("id", tid).All(&types)
			p.Data["a_type"] = types
			//tag数组
			qt.FilterRaw("id","not in ("+as.Tags+")").All(&tags)
			p.Data["a_tags"] = tags
			qt.FilterRaw("id","in ("+as.Tags+")").All(&tags)
			p.Data["a_tag"] = tags
			p.Data["a_data"] = as
			p.TplName = "personal/article-edit.html"

		} else {
			qs.All(&types)
			p.Data["a_types"] = types
			qt.All(&tags)
			p.Data["a_tags"] = tags
			p.Data["a_data"] = make(map[int]string)
			p.TplName = "personal/article.html"
		}
	}
}

//添加图片
func (p *PersonalController) PushImg()  {
	f, h, err := p.GetFile("file")
	result := make(map[string]interface{})
	img := ""
	old := h.Filename
	if err == nil {
		exStrArr := strings.Split(h.Filename, ".")
		exStr := strings.ToLower(exStrArr[len(exStrArr)-1])
		if exStr != "jpg" && exStr!="png" && exStr != "gif" {
			result["code"] = 1
			result["message"] = "上传只能.jpg 或者png格式"
		}
		defer f.Close()
		img = "static/upload/" + util.UniqueId()+"."+exStr
		p.SaveToFile("file", img) // 保存位置在 static/upload, 没有文件夹要先创建
		result["code"] = 0
		result["message"] =img
		p.SetSession(old, img)
	}else{
		result["code"] = 2
		result["message"] = "上传异常"+err.Error()
	}
	p.Data["json"] = result
	p.ServeJSON()
}

//删除图片
func (p *PersonalController) DelImg() {
	img := p.GetString("img")
	img_name := p.GetSession(img).(string)
	err := os.Remove(img_name)
	if err != nil {
		p.MsgBack("删除失败!", 0)
	}
	cnt, err :=p.o.QueryTable(new(models.Article).TableName()).Filter("picture", img_name).Count()
	if cnt > 0{
		p.o.QueryTable(new(models.Article).TableName()).Filter("picture", img_name).Update(orm.Params{
			"picture" : "",
		})
	}
	p.MsgBack("删除成功!", 1)
}

//文章列表
func (p *PersonalController) List() {
	str := p.Ctx.Input.Param(":str")
	id := p.GetSession("client_id").(int)
	var article []*models.Article
	qs := p.o.QueryTable(new(models.Article).TableName())
	qs = qs.Filter("client_id", id)
	if str != "" {
		qs = qs.Filter("title__icontains", str)
	}
	qs.RelatedSel().All(&article)
	p.Data["article_list"] = article
	p.Data["str"] = str
	p.TplName = "personal/article-list.html"
}

//文章上下架操作
func (p *PersonalController) PushPull() {
	id_ := p.GetString("id")
	id, _ := strconv.Atoi(id_)
	article := models.Article{Id:id}
	p.o.Read(&article)
	if article.Review == 1 {
		if article.Status > 0 {
			article.Status = 0
		} else {
			article.Status = 1
		}
		p.SumArticleNum(article.Type.Id, article.Status)
		_, err := p.o.Update(&article, "Status")
		if err != nil {
			p.MsgBack("操作失败", 0)
		}
		p.MsgBack("操作成功", 1)
	}
	p.MsgBack("操作失败", 0)
}

//个人设置
func (p *PersonalController) Setting() {
	uid := p.GetSession("client_id").(int)
	if p.Ctx.Request.Method == "POST" {
		name := p.GetString("name")
		email := p.GetString("email")
		mobile := p.GetString("mobile")
		motto := p.GetString("motto")
		img := p.GetString("img")
		age, _ := strconv.Atoi(p.GetString("age"))
		sex, _ := strconv.Atoi(p.GetString("sex"))
		user := models.Client{Id:uid}
		user.Pic = img
		user.Email = email
		user.Username = name
		user.Mobile = mobile
		user.Motto = motto
		user.Age = age
		user.Sex = sex
		_,err := p.o.Update(&user,"Pic","Email","Username","Mobile","Motto","Age","Sex")
		if err != nil {
			p.MsgBack("修改失败", 0)
		}
		p.MsgBack("修改成功", 1)
	} else {
		user := models.Client{Id:uid}
		p.o.Read(&user)
		if user.Pic != "" {
			p.Data["u_img"] = user.Pic
		}
		p.TplName = "personal/personal-settings.html"
	}
}

//退出登录
func (p *PersonalController) Logout()  {
	p.DestroySession()
	p.History("退出登录", "/")
}

//文章收藏/喜欢
func (p *PersonalController) Like() {
	client_id := p.GetSession("client_id").(int)
	id, _ := strconv.Atoi(p.GetString("id"))
	status, _ := strconv.Atoi(p.GetString("status"))
	if status == 1 {
		qs := p.o.QueryTable(new(models.Zan).TableName())
		qs = qs.Filter("article_id", id).Filter("client_id",client_id)
		num, _ :=qs.Count()
		if num > 0 {
			qs.Delete()
			p.setZan(id, 0)
		} else {
			zan := models.Zan{}
			zan.Client = &models.Client{Id:client_id}
			zan.Article = &models.Article{Id:id}
			_, err := p.o.Insert(&zan)
			if err != nil {
				p.MsgBack("操作失败", 0)
			}
			p.setZan(id, 1)
		}
	} else {
		qs := p.o.QueryTable(new(models.Collection).TableName())
		qs = qs.Filter("article_id", id).Filter("client_id",client_id)
		num, _ := qs.Count()
		if num > 0  {
			qs.Delete()
			p.setKeep(id, 0)
		} else {
			colle := models.Collection{}
			colle.Article = &models.Article{Id:id}
			colle.Client = &models.Client{Id:client_id}
			_, err := p.o.Insert(&colle)
			if err != nil {
				p.MsgBack("操作失败", 0)
			}
			p.setKeep(id, 1)
		}
	}
	p.MsgBack("操作成功", 1)
}

//添加,减少文章赞的数量
func (p *PersonalController) setZan(id int, status int) {
	article := models.Article{Id:id}
	p.o.Read(&article)
	if status == 1 {
		article.ZanNum = article.ZanNum + 1
	} else {
		article.ZanNum = article.ZanNum - 1
	}
	p.o.Update(&article, "ZanNum")
}

//添加,减少文章收藏的数量
func (p *PersonalController) setKeep(id int, status int) {
	article := models.Article{Id:id}
	p.o.Read(&article)
	if status == 1 {
		article.CollectNum = article.CollectNum + 1
	} else {
		article.CollectNum = article.CollectNum - 1
	}
	p.o.Update(&article, "CollectNum")
}

//个人中心首页数据
func (p *PersonalController) getPerData() {
	client_id := p.GetSession("client_id").(int)
	//conn := myredis.Conn()
	//defer conn.Close()
	//conn.Send("PING", client_id)
	article := []*models.Article{}
	qs := p.o.QueryTable(new(models.Article).TableName()).Filter("client_id", client_id)
	//总文章数
	a_num, _ := qs.Count()
	p.Data["a_num"] = a_num
	//被喜欢数,被收藏数,被评论数,被点击数(总量)
	p.Data["be_like"], p.Data["be_keep"], p.Data["be_comment"], p.Data["be_click"] = 0,0,0,0
	if a_num > 0 {
		qs.All(&article)
		be_like,be_keep,be_comment,be_click := 0,0,0,0
		for _,v:= range article {
			be_like = be_like + v.ZanNum
			be_keep = be_keep + v.CollectNum
			be_comment = be_comment + v.CommentNum
			be_click = be_click + v.ClickVolume
		}
		p.Data["be_like"], p.Data["be_keep"], p.Data["be_comment"], p.Data["be_click"] = be_like, be_keep, be_comment, be_click
	}
	//自己喜欢
	p.Data["me_like"], _ = p.o.QueryTable(new(models.Zan).TableName()).Filter("client_id",client_id).Count()
	//自己收藏
	p.Data["me_keep"], _ = p.o.QueryTable(new(models.Collection).TableName()).Filter("client_id",client_id).Count()
	//已评论文章
	p.Data["me_comment"], _ = p.o.QueryTable(new(models.Comment).TableName()).Filter("client_id",client_id).Count()
	//根据城市获取天气
	p.Data["weater"] = p.GetUserWeater()
	//根据现在ip获取城市
	p.Data["ip"] = p.GetUserIp()
	client := models.Client{Id:client_id}
	client.Ip = p.Data["ip"].(string)
	client.LoginTime = util.TimeSet()
	p.o.Update(&client, "Ip","LoginTime")
	//文章浏览数
	p.Data["look_num"], _ = p.o.QueryTable(new(models.Browse).TableName()).Filter("client_id",client_id).Count()
}
