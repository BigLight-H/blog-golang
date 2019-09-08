package routers

import (
	"beego-demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//后台页面
	beego.Router("/error", &controllers.BlogController{}, "*:Error")
	beego.Router("/login", &controllers.LoginController{}, "get:Login")
	beego.Router("/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/admin/home", &controllers.AdminController{}, "*:Home")
	beego.Router("/admin/logout", &controllers.AdminController{}, "*:Logout")
	beego.Router("/admin/classify/add/?:id", &controllers.AdminController{}, "get:ClassifyAdd")
	beego.Router("/admin/classify/add", &controllers.AdminController{}, "post:ClassifyAdd")
	beego.Router("/admin/classify/list", &controllers.AdminController{}, "*:ClassifyList")
	beego.Router("/admin/article/list", &controllers.AdminController{}, "*:ArticleList")
	beego.Router("/admin/article/review", &controllers.AdminController{}, "post:Review")
	beego.Router("/admin/article/updown", &controllers.AdminController{}, "post:UpDown")
	beego.Router("/admin/users", &controllers.AdminController{}, "*:UserList")
	beego.Router("/admin/users/user_status", &controllers.AdminController{}, "post:UserStatus")
	beego.Router("/admin/add_user", &controllers.AdminController{}, "post:AddUser")
	beego.Router("/admin/user/detail_user", &controllers.AdminController{}, "*:UserMessge")
	beego.Router("/admin/user/detail_user", &controllers.AdminController{}, "post:UserMessge")
	beego.Router("/admin/user/feed_back", &controllers.AdminController{}, "*:FeedBack")
	beego.Router("/admin/user/push_email", &controllers.AdminController{}, "post:PushEmail")

	//前台页面
	beego.Router("/", &controllers.HomeController{}, "*:Index")
	beego.Router("/home_login", &controllers.LoginController{}, "post:HomeLogin")
	beego.Router("/home_register", &controllers.LoginController{}, "post:Register")
	beego.Router("/detail/?:id", &controllers.HomeController{}, "get:Detail")
	beego.Router("/author/?:id", &controllers.HomeController{}, "get:Author")
	beego.Router("/comment/add", &controllers.HomeController{}, "post:AddComment")
	beego.Router("/leave/comments", &controllers.HomeController{}, "*:SetMessage")
	beego.Router("/leave/comments", &controllers.HomeController{}, "post:SetMessage")
	beego.Router("/about", &controllers.HomeController{}, "*:About")
	beego.Router("/home/error", &controllers.HomeController{}, "*:HomeError")

	//个人中心路由
	beego.Router("/personal", &controllers.PersonalController{}, "*:PersonalCenter")
	beego.Router("/push_img", &controllers.PersonalController{}, "post:PushImg")
	beego.Router("/del_img", &controllers.PersonalController{}, "post:DelImg")
	beego.Router("/add/article/?:id", &controllers.PersonalController{}, "*:AddArticle")
	beego.Router("/add/article", &controllers.PersonalController{}, "post:AddArticle")
	beego.Router("/article/list", &controllers.PersonalController{}, "*:List")
	beego.Router("/article/push_pull", &controllers.PersonalController{}, "post:PushPull")
	beego.Router("/personal/settings", &controllers.PersonalController{}, "*:Setting")
	beego.Router("/personal/settings", &controllers.PersonalController{}, "post:Setting")
}
