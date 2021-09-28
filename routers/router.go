package routers

import (
	"github.com/max-las/humans-of-grenoble/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
		beego.ErrorController(&controllers.ErrorController{})

    beego.Router("/", &controllers.HomeController{})
		beego.Router("/stories", &controllers.StoriesController{})
		beego.Router("/story/:id", &controllers.StoryController{})
		beego.Router("/projet", &controllers.ProjetController{})
		beego.Router("/contact", &controllers.ContactController{})

		beego.Router("/admin", &controllers.AdminHomeController{})
		beego.Router("/admin/cloudinary", &controllers.CloudinaryController{})
		beego.Router("/admin/login", &controllers.LoginController{})
		beego.Router("/admin/new", &controllers.NewStoryController{})
		beego.Router("/admin/edit/:id", &controllers.EditStoryController{})
		beego.Router("/admin/new-password/", &controllers.NewPasswordController{})

		beego.Router("/adduser", &controllers.AdduserController{})

		beego.ErrorController(&controllers.ErrorController{})
}
