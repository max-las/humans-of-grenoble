package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
  "github.com/max-las/humans-of-grenoble/models"
  "github.com/max-las/humans-of-grenoble/helpers"
)

type AdduserController struct {
	beego.Controller
}

func (c *AdduserController) Get() {
  on := false;

  c.Data["Message"] = "Connexion"
  c.TplName = "dev/simpleMessage.tpl"

  if(on){

    username := ""
    password := ""

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
  }else{
    c.Data["Message"] = "off"
  }
}
