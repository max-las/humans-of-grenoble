package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"fmt"
	"github.com/max-las/humans-of-grenoble/helpers"
)

type ProjetController struct {
	beego.Controller
}

func (c *ProjetController) Get() {
	c.Data["PageTitle"] = "Projet | Humans of Grenoble"
	c.Layout = "layouts/main.tpl"
	c.TplName = "projet.tpl"

	etag := helpers.TplLastModifiedString(c.TplName)
	c.Ctx.Output.Header("ETag", fmt.Sprintf("\"%s\"", etag))
}
