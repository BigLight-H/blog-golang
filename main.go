package main

import (
	_ "beego-demo/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"beego-demo/models"
)


func init() {
	orm.Debug = true
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.CopyRequestBody = true
	beego.SetStaticPath("/static","static/")
}

func main() {
	beego.Run()
}

