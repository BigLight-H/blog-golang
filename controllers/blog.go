package controllers

import (
	"axicoo/models"
	"github.com/astaxie/beego"
)

type BlogController struct {
	baseController
}

func (c *BlogController) Home()  {
	categorys := [] *models.Category{}
	c.o.QueryTable( new(models.Category).TableName()).All(&categorys)
	c.Data["cates"] = categorys
	beego.Info(c.Data)
	c.Ctx.WriteString("22222")
}
