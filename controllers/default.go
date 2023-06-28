package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}
type IndexController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["test"] = "小甜甜"
	c.TplName = "index.html"
}
func (this *IndexController) Post() {
	this.Data["test"] = "小甜甜我的最爱"
	this.TplName = "index.html"
}
