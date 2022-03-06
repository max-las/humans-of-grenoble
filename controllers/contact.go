package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
  "net/smtp"
  "fmt"
  "os"
  "crypto/tls"
  "html/template"
  "github.com/domodwyer/mailyak"
  "github.com/max-las/humans-of-grenoble/helpers"
  "github.com/beego/beego/v2/client/httplib"
)

type ContactController struct {
	beego.Controller
}

type ContactData struct {
	Name string
	Email string
  Subject string
  Message string
  RecaptchaToken string
}

type RecaptchaVerifyResponse struct {
  Success bool `json:"success"`
  Score float32 `json:"score"`
  Action string `json:"action"`
  ChallengeTs string `json:"challenge_ts"`
  Hostname string `json:"hostname"`
  ErrorCodes []string `json:"error-codes"`
}

func (c *ContactController) Get() {
	c.Data["PageTitle"] = "Contact | Humans of Grenoble"
	c.Layout = "layouts/main.tpl"
	c.TplName = "contact.tpl"

  etag := helpers.TplLastModifiedString(c.TplName)
  helpers.HandleEtag(&c.Controller, etag)
}

func (c *ContactController) Post() {
  var err error

  c.TplName = "dev/simpleMessage.tpl"

  data := ContactData{
    Name: c.GetString("name"),
  	Email: c.GetString("email"),
    Subject: c.GetString("subject"),
    Message: c.GetString("message"),
    RecaptchaToken: c.GetString("g-recaptcha-response"),
  }

  if(data.Name == "" || data.Email == "" || data.Subject == "" || data.Message == "" || data.RecaptchaToken == ""){
    c.Abort("400")
  }

  g_req := httplib.Post("https://www.google.com/recaptcha/api/siteverify")
  g_req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
  g_req.Param("secret", "6LeSGroeAAAAAK3ofYU3MRvYtYbdW_q-Tw8uo-1-")
  g_req.Param("response", data.RecaptchaToken)

  var g_res RecaptchaVerifyResponse
  g_req.ToJSON(&g_res)

  if(g_res.Score <= 0.5){
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

  err = mail.Send()
  if err != nil {
    fmt.Println(err)
    c.Abort("500")
  }
  c.Data["Message"] = "OK"
}
