package controllers

import (
	"beego-demo/models"
	"beego-demo/util"
	"github.com/davecgh/go-spew/spew"
	"strconv"
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

//文章类型列表
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
	_, error := p.o.Update(&article, "Status");
	if error != nil {
		p.MsgBack("审核失败", 0)
	}
	p.MsgBack("审核成功", 1)
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
	spew.Dump(p.Data["user"])
	//if p.Ctx.Request.Method == "POST" {
	//	spew.Dump()
	//}
	p.Ctx.WriteString("个人信息修改")
}

//意见反馈列表
func (p *AdminController) FeedBack() {
	feedback := []*models.FeedBack{}
	p.o.QueryTable(new(models.FeedBack).TableName()).OrderBy("-created").All(&feedback)
	p.Data["feedback"] = feedback
	spew.Dump(feedback)
	p.TplName = ""
}




