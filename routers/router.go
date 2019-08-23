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
}
