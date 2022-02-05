package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/max-las/humans-of-grenoble/helpers"
	"fmt"
)

type LegalController struct {
	beego.Controller
}

func (c *LegalController) Get() {
	c.Data["PageTitle"] = "Mentions légales | Humans of Grenoble"
	c.Layout = "layouts/main.tpl"
	c.TplName = "legal.tpl"

	etag := helpers.TplLastModifiedString(c.TplName)
	c.Ctx.Output.Header("ETag", fmt.Sprintf("\"%s\"", etag))
}
