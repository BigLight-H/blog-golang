package controllers

import (
	"axicoo/models"
	"axicoo/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	jsoniter "github.com/json-iterator/go"
	"regexp"
	"strings"
)

type baseController struct {
	beego.Controller
	o orm.Ormer
	controllerName string
	actionName     string
}

type Json struct {
	Msg string
	Status int
}

func (p *baseController) Prepare()  {
	controllerName, actionName := p.GetControllerAndAction()
	p.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	p.actionName = strings.ToLower(actionName)
	p.o = orm.NewOrm()
	if p.controllerName == "admin"{
		if p.GetSession("user") == nil {
			p.History("", "/error")
		}
		permissions := [] *models.Permissions{}
		p.o.QueryTable(new(models.Permissions).TableName()).Filter("status", 1).All(&permissions)
		p.Data["sidebar"] = &permissions
		p.Data["user"] = p.GetSession("user")
		p.Data["tag"] = "Admin"
	} else if p.controllerName == "home" {
		menu := [] *models.Type{}
		p.o.QueryTable(new(models.Type).TableName()).Filter("status", 1).All(&menu)
		p.Data["menu"] = menu
		p.hot()//热门文章
		p.tags()//全部标签
		p.footData()//footer数据
	} else if p.controllerName == "personal" {
		if p.GetSession("client") == nil {
			p.History("", "/home/error")
		}
	}
	if p.GetSession("client") != nil {
		client_id := p.GetSession("client_id").(int)
		p.Data["client_id"] = client_id
		p.Data["client"] = p.GetSession("client")
	}
}
//热门文章
func (p *baseController) hot() {
	article := []*models.Article{}
	qs := p.o.QueryTable(new(models.Article).TableName())
	qs = qs.Filter("status", 1)
	qs.OrderBy("-click_volume").Limit(3).All(&article)
	p.Data["hot"] = article//热门文章
}
//文章标签
func (p *baseController) tags() {
	tags := []*models.Tags{}
	p.o.QueryTable(new(models.Tags).TableName()).All(&tags)
	p.Data["tag_all"] = tags
}
//作者数量
//文章数量
//评论数量
//最近文章
func (p *baseController) footData() {
	//文章
	footer_article, _ := p.o.QueryTable(new(models.Article)).Filter("status", 1).Count()
	p.Data["footer_article"]= footer_article
	//作者
	footer_client, _ := p.o.QueryTable(new(models.Client).TableName()).Count()
	p.Data["footer_client"] = footer_client
	//评论数量
	footer_comment, _ := p.o.QueryTable(new(models.Comment).TableName()).Count()
	p.Data["footer_comment"] = footer_comment
	//最近文章
	article :=[]*models.Article{}
	p.o.QueryTable(new(models.Article).TableName()).Filter("status", 1).OrderBy("-id").Limit(3).All(&article)
	p.Data["footer_article_data"] = article

}

//文章类型数量加减
func (p *baseController) SumArticleNum(typeId int, status int) {
	types := models.Type{Id:typeId}
	if status > 0 {
		types.ArticleNum = types.ArticleNum + 1
	} else {
		if types.ArticleNum > 0 {
			types.ArticleNum = types.ArticleNum - 1
		}
	}
	p.o.Update(&types, "ArticleNum")
}

func (p *baseController) History(msg string, url string)  {
	if url == "" {
		p.Ctx.WriteString("<script>alert('"+msg+"');window.history.go(-1);</script>")
		p.StopRun()
	} else {
		p.Redirect(url,302)
	}
}

func (p *baseController) MsgBack(msg string, status int)  {
	data := &Json{
		msg,
		status,
	}
	p.Data["json"] = data
	p.ServeJSON()
}

//获取用户IP地址
func (p *baseController) getClientIp() string {
	s := strings.Split(p.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

//email verify
func (p *baseController) VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
//mobile verify
func (p *baseController) VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

//获取用户ip
func (p *baseController) GetUserIp() string {
	if p.GetSession("ip") == nil {
		res ,_ := util.DoHttpGetRequest("https://ip.seeip.org/geoip")
		params := p.AnalyzeJson(res)
		ip := params["ip"].(string)
		p.SetSession("ip", ip)
		return ip
	}
	return p.GetSession("ip").(string)
}

//获取用户所在城市天气
func (p *baseController) GetUserWeater() map[string]interface{} {
	if p.GetSession("weater") == nil {
		res, _ := util.DoHttpGetRequest(beego.AppConfig.String("hf_api"))
		params := p.AnalyzeJson(res)
		HeWeather6 := params["HeWeather6"].([]interface{})
		data := HeWeather6[0].(map[string]interface{})
		p.SetSession("weater", data)
		return data
	}
	return p.GetSession("weater").(map[string]interface{})
}


//解析json
func (p *baseController) AnalyzeJson(src string) map[string]interface{} {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	reader := strings.NewReader(src)
	decoder := json.NewDecoder(reader)
	params := make(map[string]interface{})
	err := decoder.Decode(&params)
	if err == nil {
		return params
	}
	return params
}

//增加点击量
func (p *baseController) AddClickVolume(aid int) {
	_, _ = p.o.QueryTable(new(models.Article).TableName()).Filter("id", aid).Update(orm.Params{
		"click_volume": orm.ColValue(orm.ColAdd, 1),
	})
}

//添加评论通知数量
func (p *baseController) AddCommentNum(aid int) {
	if p.GetSession("client_id") != nil  {
		cid := p.GetSession("client_id").(int)
		a := models.ArticleCommentNum{ArticleId:aid}
		p.o.Read(&a)
		if a.ClientId > 0 {
			_, _ = p.o.QueryTable(new(models.ArticleCommentNum).TableName()).Filter("article_id", aid).Update(orm.Params{
				"number": orm.ColValue(orm.ColAdd, 1),
			})
		} else {
			a.ArticleId = aid
			a.ClientId = cid
			a.Number = 1
			p.o.Insert(&a)
		}
	}
}

//已读某文章评论
func (p *baseController) ReadArticle(aid int) {
	if p.GetSession("client_id") != nil {
		cid := p.GetSession("client_id").(int)
		_,_ = p.o.QueryTable(new(models.ArticleCommentNum).TableName()).Filter("client_id", cid).Filter("article_id",aid).Delete()
	}
}

//记录文章浏览记录
func (p *baseController) AddLook(aid int) {
	if p.GetSession("client_id") != nil {
		cid := p.GetSession("client_id").(int)
		num, _ := p.o.QueryTable(new(models.Browse).TableName()).Filter("client_id", cid).Filter("article_id", aid).Count()
		if num < 1 {
			browse := models.Browse{}
			browse.Client = &models.Client{Id:cid}
			browse.Article = &models.Article{Id:aid}
			p.o.Insert(&browse)
		}
	}
}