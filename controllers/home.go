package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/max-las/humans-of-grenoble/helpers"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["PageTitle"] = "Humans of Grenoble"
	c.Layout = "layouts/home.tpl"
	c.TplName = "home.tpl"

	etag := helpers.TplLastModifiedString(c.TplName)
	helpers.HandleEtag(&c.Controller, etag)
}
