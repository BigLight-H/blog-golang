package routers

import (
	"axicoo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.BlogController{}, "*:Home")
}
