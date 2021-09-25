package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
  "net/smtp"
  "fmt"
  "os"
  "html/template"
  "github.com/domodwyer/mailyak"
)

type ContactController struct {
	beego.Controller
}

type ContactData struct {
	Name string
	Email string
  Subject string
  Message string
}

func (c *ContactController) Get() {
	c.Data["PageTitle"] = "Contact | Humans of Grenoble"
	c.Layout = "layouts/main.tpl"
	c.TplName = "contact.tpl"
}

func (c *ContactController) Post() {
  c.TplName = "dev/simpleMessage.tpl"

  data := ContactData{
    Name: c.GetString("name"),
  	Email: c.GetString("email"),
    Subject: c.GetString("subject"),
    Message: c.GetString("message"),
  }

  if(data.Name == "" || data.Email == "" || data.Subject == "" || data.Message == ""){
    c.Abort("400")
  }

  username := os.Getenv("SMTP_USERNAME")
  password := os.Getenv("SMTP_PASSWORD")
  smtpHost := os.Getenv("SMTP_HOST")
  smtpPort := os.Getenv("SMTP_PORT")
  recipient := os.Getenv("CONTACT_RECIPIENT")

  mail := mailyak.New(smtpHost+":"+smtpPort, smtp.PlainAuth("", username, password, smtpHost))
  mail.To(recipient)
  mail.From(username)
  mail.FromName("HumansOfGrenoble")
  mail.Subject(data.Subject)

  t := template.Must(template.ParseFiles("views/mail/contact.tpl"))
  t.Execute(mail.HTML(), data)

  err := mail.Send()
  if err != nil {
    fmt.Println(err)
    c.Abort("500")
  }
  c.Data["Message"] = "OK"
}
