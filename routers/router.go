package routers

import (
	"beego-demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.BlogController{}, "*:Home")
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
}
