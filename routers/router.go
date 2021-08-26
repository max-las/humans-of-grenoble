package routers

import (
	"github.com/max-las/humans-of-grenoble/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
		beego.ErrorController(&controllers.ErrorController{})

    beego.Router("/", &controllers.HomeController{})
		beego.Router("/stories", &controllers.StoriesController{})
		beego.Router("/admin", &controllers.AdminHomeController{})
		beego.Router("/admin/login", &controllers.LoginController{})
		beego.Router("/admin/logout", &controllers.LogoutController{})
		beego.Router("/admin/new", &controllers.NewStoryController{})
		beego.Router("/admin/edit/:id", &controllers.EditStoryController{})
		beego.Router("/story/:id", &controllers.StoryController{})

		runmode, _ := beego.AppConfig.String("runmode")
		if(runmode == "dev"){
			beego.Router("/adduser", &controllers.AdduserController{})
		}
}
