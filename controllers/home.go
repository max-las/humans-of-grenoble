package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["PageTitle"] = "Humans of Grenoble"
	c.Layout = "layouts/home.tpl"
	c.TplName = "home.tpl"
}
