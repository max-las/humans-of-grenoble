package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type LogoutController struct {
	beego.Controller
}

func (c *LogoutController) Get() {
  c.DelSession("username")
  c.SetSession("logout", true)
  c.Redirect("/admin/login", 303)
}
