package main

import (
	_ "beego-demo/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"beego-demo/models"
)


func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.Run()
}

