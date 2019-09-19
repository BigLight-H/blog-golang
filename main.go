package main

import (
	_ "axicoo/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"axicoo/models"
)


func init() {
	orm.Debug = true
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.CopyRequestBody = true
	beego.SetStaticPath("/static","static/")
	beego.SetStaticPath("/home/","static/home/")
}

func main() {
	beego.Run()
}

