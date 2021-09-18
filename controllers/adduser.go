package controllers

import (
	"os"

	beego "github.com/beego/beego/v2/server/web"
  "github.com/max-las/humans-of-grenoble/models"
  "github.com/max-las/humans-of-grenoble/helpers"
)

type AdduserController struct {
	beego.Controller
}

func (c *AdduserController) Get() {
  c.TplName = "dev/simpleMessage.tpl"

  username := os.Getenv("NEW_USER_USERNAME")
  password := os.Getenv("NEW_USER_PASSWORD")

	os.Unsetenv("NEW_USER_USERNAME")
	os.Unsetenv("NEW_USER_PASSWORD")

	if(username == "" || password == ""){
		c.Abort("404")
	}

	user, _ := models.GetUserByUsername(username)
	if(user != nil){

		c.Data["Message"] = "user already exists"

	}else{

		hash, err := helpers.HashPassword(password)
		if(err != nil){
			c.Data["Message"] = "bcrypt error"
		}else{
			_, err = models.AddUser(&models.User{
				Username: username,
				Password: hash,
			})

			if(err != nil){
				c.Data["Message"] = "orm error"
			}else{
				c.Data["Message"] = "success"
			}
		}

	}
}
