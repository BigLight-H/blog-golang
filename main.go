package main

import (
	_ "axicoo/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"axicoo/models"
)


func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.Run()
}

