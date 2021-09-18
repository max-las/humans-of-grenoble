package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ProjetController struct {
	beego.Controller
}

func (c *ProjetController) Get() {
	c.Data["PageTitle"] = "Projet | Humans of Grenoble"
	c.Layout = "layouts/main.tpl"
	c.TplName = "projet.tpl"
}
