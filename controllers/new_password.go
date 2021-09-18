package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/max-las/humans-of-grenoble/models"
	"github.com/max-las/humans-of-grenoble/helpers"
	"fmt"
)

type NewPasswordController struct {
	beego.Controller
}

func (c *NewPasswordController) Get() {
  c.Layout = "layouts/main.tpl"
  c.TplName = "admin/new_password.tpl"
}

func (c *NewPasswordController) Post() {
  c.TplName = "dev/simpleMessage.tpl"

	newPass := c.GetString("password")

	if(newPass == ""){
		c.Abort("400")
	}

	hash, err := helpers.HashPassword(newPass)
	if(err != nil){
		fmt.Println("bcrypt error")
		c.Abort("500")
	}

	username := c.GetSession("username").(string)
	fmt.Println(username)

	err = models.UpdateUserByUsername(&models.User{
		Username: username,
		Password: hash,
	})

	if(err != nil){
		fmt.Println(err)
		c.Abort("500")
	}else{
		c.Data["Message"] = "OK"
	}

}
