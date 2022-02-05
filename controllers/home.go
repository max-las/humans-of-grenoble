package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/max-las/humans-of-grenoble/helpers"
	"fmt"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["PageTitle"] = "Humans of Grenoble"
	c.Layout = "layouts/home.tpl"
	c.TplName = "home.tpl"

	etag := helpers.TplLastModifiedString(c.TplName)
	c.Ctx.Output.Header("ETag", fmt.Sprintf("\"%s\"", etag))
}
