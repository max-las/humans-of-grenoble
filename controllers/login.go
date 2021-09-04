package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
	"github.com/max-las/humans-of-grenoble/models"
	"github.com/max-las/humans-of-grenoble/helpers"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Prepare() {
	username := c.GetSession("username")
	if(username != nil && !c.Ctx.Input.Is("DELETE")){
		c.Redirect("/admin", 303)
	}else{
		c.Data["PageTitle"] = "Connexion | Humans of Grenoble"

		c.Layout = "layouts/main.tpl"
		c.TplName = "admin/login.tpl"
	}
}

func (c *LoginController) Get() {
	logout := c.GetSession("logout")
	if(logout != nil){
		c.Data["Success"] = "Déconnexion réussie."
		c.DestroySession()
	}
}

func (c *LoginController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")

	if(len(username) < 1 || len(password) < 1){
		c.Redirect("/admin", 302)
	}

	user, err := models.GetUserByUsername(username)
	if(err != nil){
		if(err == orm.ErrNoRows){
			c.Data["Error"] = "Identifiant ou mot de passe incorrect."
		}else{
			c.Abort("500")
		}
	}else{
		passOK := helpers.CheckPasswordHash(password, user.Password)
		if(!passOK){
			c.Data["Error"] = "Identifiant ou mot de passe incorrect."
		}else{
			c.SetSession("username", username)
			c.Redirect("/admin", 303)
		}
	}
}

func (c *LoginController) Delete() {
	c.DelSession("username")
	c.SetSession("logout", true)
	c.Ctx.Output.Body([]byte("OK"))
}
