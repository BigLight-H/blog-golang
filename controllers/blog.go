package controllers

import (
	"beego-demo/models"
)

type BlogController struct {
	baseController
}

func (c *BlogController) Home()  {
	post := models.Permissions{Id:1}
	c.o.Read(&post) // 返回 QuerySeter
	c.Data["json"] = &post
	c.ServeJSON()
}

func (c *BlogController) Error()  {
	c.Data["tag"] = "Error"
	c.TplName = "admin/500.html"
}
