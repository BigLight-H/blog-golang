package controllers

import (
	"axicoo/models"
	"axicoo/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/davecgh/go-spew/spew"
	"os"
	"strconv"
	"strings"
)

type AdminController struct {
	baseController
}

//后台首页
func (p *AdminController) Home()  {
	p.TplName = "admin/index.html"
}

//退出登录
func (p *AdminController) Logout()  {
	p.DestroySession()
	p.History("退出登录", "/login")
}

//文章类型添加
func (p *AdminController) ClassifyAdd() {
	if p.Ctx.Request.Method == "POST" {
		types := models.Type{}
		pid_ := p.Input().Get("pid")
		pid, _ := strconv.Atoi(pid_)
		name := p.Input().Get("name")
		id  := p.Input().Get("id")
		url  := p.Input().Get("url")
		icon  := p.Input().Get("icon")
		if id != "" {
			id, _ := strconv.Atoi(id)
			types.Id = id
			types.Pid = pid
			types.TName = name
			types.Url = url
			types.Icon = icon
			types.ArticleNum = 0
			if pid == 0 {
				types.Dir = 1
			} else {
				types.Dir = 0
			}
			if _, err := p.o.Update(&types); err != nil {
				p.History("", "/admin/classify/add/"+strconv.Itoa(id))
			} else {
				p.History("", "/admin/classify/list")
			}
		} else {
			types.Pid = pid
			types.TName = name
			types.Status = 1
			types.Icon = icon
			if pid == 0 {
				types.Dir = 1
			} else {
				types.Dir = 0
			}
			types.Url = url
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
			id_, _ := strconv.Atoi(id)
			t := models.Type{Id:id_}
			p.o.Read(&t)
			pid := strconv.Itoa(t.Pid)
			qs.Filter("status",1).FilterRaw("id", "!="+pid).All(&class)
			p.Data["class_type"] = class
			qs.Filter("id", id).One(&class)
			p.Data["type_data"] = class
			qs.FilterRaw("id", "="+pid).One(&class)
			p.Data["type_title"] = class
		} else {
			qs.Filter("status",1).All(&class)
			p.Data["class_type"] = class
		}
		p.TplName = "admin/classify_add.html"
	}
}

//文章类型列表
func (p *AdminController) ClassifyList() {
	if p.Ctx.Request.Method == "POST"{
		class := models.Type{}
		p.o.QueryTable(new(models.Type).TableName())
		ids := p.Input().Get("id")
		id, _ := strconv.Atoi(ids)
		class.Status = 0
		class.Id = id
		pid_ := p.Input().Get("pid")
		pid, _ := strconv.Atoi(pid_)
		class.Pid = pid
		class.TName = p.Input().Get("name")
		if _, err := p.o.Update(&class); err == nil {
			p.MsgBack("删除成功", 1)
		} else {
			p.MsgBack("删除失败", 0)
		}
	} else {
		class := []*models.Type{}
		qs := p.o.QueryTable(new(models.Type).TableName())
		qs.Filter("status",1).OrderBy("pid").All(&class)
		p.Data["types"] = class
		p.TplName = "admin/classify_list.html"
	}
}

//文章列表渲染
func (p *AdminController) ArticleList() {
	p.list()
	p.TplName = "admin/article_list.html"
}

//文章列表查询
func (p *AdminController) list() {
	article := []*models.Article{}
	p.o.QueryTable(new(models.Article).TableName()).RelatedSel().All(&article)
	p.Data["article"] = article
}

//文章审核确认
func (p *AdminController) Review()  {
	id_ := p.GetString("id")
	id, _ := strconv.Atoi(id_)
	revs := p.GetString("rev")
	rev, _ := strconv.Atoi(revs)
	txt := p.GetString("txt")
	article := models.Article{}
	article.Id = id
	article.Review = rev
	if txt != "" {
		article.Cause = txt
	}
	_, erro := p.o.Update(&article, "Review", "Cause"); if erro != nil {
		p.MsgBack("审核失败", 0)
	}
	p.MsgBack("审核成功", 1)
}

//文章上下架操作
func (p *AdminController) UpDown()  {
	id_ := p.GetString("id")
	id, _ := strconv.Atoi(id_)
	status_ := p.GetString("status")
	status, _ := strconv.Atoi(status_)
	article := models.Article{}
	article.Id = id
	article.Status = status
	p.SumArticleNum(article.Type.Id, status)
	_, error := p.o.Update(&article, "Status")
	if error != nil {
		p.MsgBack("操作失败", 0)
	}
	p.MsgBack("操作成功", 1)
}

//后台管理员列表
func (p *AdminController) UserList() {
	p.users()
	p.TplName = "admin/user_list.html"
}

//管理员列表查询
func (p *AdminController) users() {
	user := []*models.User{}
	p.o.QueryTable( new(models.User).TableName() ).RelatedSel().All(&user)
	p.Data["users_list"] = user
}

//冻结解冻管理员
func (p *AdminController) UserStatus() {
	id_ := p.GetString("id")
	id, _ := strconv.Atoi(id_)
	status_ := p.GetString("status")
	status, _ := strconv.Atoi(status_)
	user := models.User{}
	user.Id = id
	if id == 1 {
		p.MsgBack("操作失败", 0)
	}
	user.Status = status
	_, err := p.o.Update(&user, "Status");
	if err != nil {
		p.MsgBack("操作失败", 0)
	}
	p.MsgBack("操作成功", 1)
}

//添加管理员
func (p *AdminController) AddUser() {
	username, password, mobile, email := p.GetString("username"), util.Md5(p.GetString("password")), p.GetString("mobile"), p.GetString("email")
	if len(username) <= 2 {
		p.MsgBack("用户名长度不能小于2", 0)
	} else if len(password) <= 6 {
		p.MsgBack("密码长度不能小于6", 0)
	} else if p.VerifyEmailFormat(email) == false {
		p.MsgBack("邮箱格式不对", 0)
	} else if p.VerifyMobileFormat(mobile) == false {
		p.MsgBack("手机号码格式不对", 0)
	} else {
		u := models.User{Username: username}
		p.o.Read(&u, "username")
		if u.Password != "" {
			p.MsgBack("用户名已存在", 0)
		} else {
			status := 1
			user := models.User{}
			user.Status = status
			user.Email = email
			user.Username = username
			user.Password = password
			user.Mobile = mobile
			user.LastIp = p.getClientIp()
			user.LoginCount = 0
			role := models.Role{Id:2}
			user.Role = &role
			img := models.HeadImg{Id:1}
			user.HeadImg = &img
			_, error := p.o.Insert(&user)
			if error != nil {
				p.MsgBack("添加管理员失败", 0)
			}
			p.MsgBack("管理员添加成功", 1)
		}
	}
}

//个人信息修改
func (p *AdminController) UserMessge() {
	if p.Ctx.Request.Method == "POST" {
		uid := p.GetSession("user_id")
		mobile := p.GetString("mobile")
		password := p.GetString("password")
		email := p.GetString("email")
		spew.Dump(uid)
		user := models.User{}
		user.Id = uid.(int)
		user.Mobile = mobile
		user.Email = email
		if password != "" {
			user.Password = util.Md5(password)
			_, err := p.o.Update(&user, "Mobile", "Email", "Password");if err != nil {
				p.MsgBack("信息修改失败", 0)
			}
		} else {
			_, err := p.o.Update(&user, "Mobile", "Email");if err != nil {
				p.MsgBack("信息修改失败", 0)
			}
		}
		p.MsgBack("信息修改成功", 1)
	}
	p.TplName = "admin/admin_message.html"
}

//意见反馈列表
func (p *AdminController) FeedBack() {
	feedback := []*models.FeedBack{}
	p.o.QueryTable(new(models.FeedBack).TableName()).OrderBy("-id").All(&feedback)
	p.Data["feedback"] = feedback
	p.TplName = "admin/feed_back.html"
}

//文章评论列表
func (p *AdminController) Comment() {
//	todo
}

//发送意见反馈回复
func (p *AdminController) PushEmail() {
	txt := p.GetString("txt")
	id_ := p.GetString("id")
	email := p.GetString("email")
	if txt == "" {
		p.MsgBack("内容不能为空", 0)
	} else if p.VerifyEmailFormat(email) == false  {
		p.MsgBack("邮箱格式不对", 0)
	} else {
		feed := models.FeedBack{}
		id, _ := strconv.Atoi(id_)
		feed.Id = id
		feed.Reply = txt
		_, err := p.o.Update(&feed, "Reply");if err != nil {
			p.MsgBack("回复返回失败", 0)
		}
		mailTo := []string{
			email,
		}
		subject := beego.AppConfig.String("set_title")
		body := txt
		_ = util.SendEmail(mailTo, subject, body)
		p.MsgBack("回复成功", 1)
	}
}

//关于我板块
func (p *AdminController) About() {
	if p.Ctx.Request.Method == "POST" {
		img := p.GetString("logo")
		content := p.GetString("content")
		about := models.About{Id: 1}
		about.Content = content
		about.Img = img
		if _, err := p.o.Update(&about); err == nil {
			p.MsgBack("添加成功", 1)
		}
		p.MsgBack("添加失败", 0)
	} else {
		about := models.About{Id:1}
		p.o.Read(&about)
		p.Data["about"] = about
		p.TplName = "admin/about.html"
	}
}

//添加图片
func (p *AdminController) PushImg()  {
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
func (p *AdminController) DelImg() {
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





