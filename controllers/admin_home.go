package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
	"github.com/max-las/humans-of-grenoble/models"
)

type AdminHomeController struct {
	beego.Controller
}

func (c *AdminHomeController) Get() {
  username := c.GetSession("username")
  if(username == nil){
    c.Redirect("/admin/login", 302)
  }else{
    c.Data["PageTitle"] = "Administration"
		c.Data["AdditionnalScripts"] = [1]string{"/static/private/js/home.js"}
    c.Layout = "layouts/main.tpl"
    c.TplName = "admin/home.tpl"

		stories, err := models.GetAllStory(nil, nil, nil, nil, 0, 100)
		if(err != nil){
			if(err != orm.ErrNoRows){
				fmt.Println(err.Error())
				c.Abort("500")
			}
		}else{
			c.Data["Stories"] = stories
		}
  }
}
