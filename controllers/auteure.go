package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type AuteureController struct {
	beego.Controller
}

func (c *AuteureController) Get() {
	c.Data["PageTitle"] = "Auteure | Humans of Grenoble"
	c.Layout = "layouts/main.tpl"
	c.TplName = "auteure.tpl"
}
