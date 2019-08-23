package controllers

import (
	"beego-demo/models"
)

type BlogController struct {
	baseController
}

func (c *BlogController) Home()  {
	//var (
	//	list       []*models.Post
	//)
	//categorys := [] *models.Category{}
	//c.o.QueryTable( new(models.Category).TableName()).All(&categorys)
	//c.Data["cates"] = categorys
	//
	//query := c.o.QueryTable(new(models.Post).TableName())
	//query.Filter("id", 1).All(&list)
	//categorys := [] *models.Category{}
	//c.o.QueryTable( new(models.Category).TableName()).All(&categorys)
	//p := [] *models.Post{}
	//post := c.o.QueryTable( new(models.Post).TableName())
	//post = post.Filter("id", 1)
	//post.All(&p)
	//c.Data["cates"] = categorys


	post := models.Permissions{Id:1}
	c.o.Read(&post) // 返回 QuerySeter
	c.Data["json"] = &post
	c.ServeJSON()
}

func (c *BlogController) Error()  {
	c.TplName = "admin/500.html"
}
