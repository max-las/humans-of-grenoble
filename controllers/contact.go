package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
  "net/smtp"
  "fmt"
)

type ContactController struct {
	beego.Controller
}

func (c *ContactController) Get() {
	c.Data["PageTitle"] = "Contact | Humans of Grenoble"
	c.Layout = "layouts/main.tpl"
	c.TplName = "contact.tpl"
}

func (c *ContactController) Post() {
  c.TplName = "dev/simpleMessage.tpl"

  name := c.GetString("name")
  email := c.GetString("email")
  subject := c.GetString("subject")
  message := c.GetString("message")
  if(name == "" || email == "" || subject == "" || message == ""){
    c.Abort("400")
  }

  username := ""
  password := ""
  smtpHost := ""
  smtpPort := ""

  from := ""
  to := []string{
    "",
  }

  actualMessage := []byte(message)

  auth := smtp.PlainAuth("", username, password, smtpHost)

  err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, actualMessage)
  if err != nil {
    fmt.Println(err)
    c.Abort("500")
  }
  c.Data["Message"] = "OK"
}
