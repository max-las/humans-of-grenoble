package controllers

import (
    "github.com/beego/beego/v2/server/web"
    "strconv"
)

type ErrorController struct {
    web.Controller
}

func (c *ErrorController) Prepare() {
    c.TplName = "dev/simpleMessage.tpl"
    c.Data["PageTitle"] = strconv.Itoa(c.Ctx.Output.Status) + " | Humans of Grenoble"
}

func (c *ErrorController) Error403() {
    c.Data["Message"] = "403"
}

func (c *ErrorController) Error404() {
    c.TplName = "errors/404.tpl"
    c.Layout = "layouts/main.tpl"
}

func (c *ErrorController) Error405() {
    c.Data["Message"] = "405"
}

func (c *ErrorController) Error413() {
    c.Data["Message"] = "413"
}

func (c *ErrorController) Error417() {
    c.Data["Message"] = "417"
}

func (c *ErrorController) Error422() {
    c.Data["Message"] = "422"
}

func (c *ErrorController) Error500() {
    c.Data["Message"] = "500"
}

func (c *ErrorController) Error501() {
    c.Data["Message"] = "501"
}

func (c *ErrorController) Error502() {
    c.Data["Message"] = "502"
}

func (c *ErrorController) Error503() {
    c.Data["Message"] = "503"
}

func (c *ErrorController) Error504() {
    c.Data["Message"] = "504"
}
