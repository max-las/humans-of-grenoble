package controllers

import (
    "github.com/beego/beego/v2/server/web"
)

type ErrorController struct {
    web.Controller
}

// func (c *ErrorController) Error404() {
//     c.Data["Message"] = "michel"
//     c.TplName = "dev/simpleMessage.tpl"
// }
//
// func (c *ErrorController) Error413() {
//     c.Data["Message"] = "too large motherfucker"
//     c.TplName = "dev/simpleMessage.tpl"
// }
