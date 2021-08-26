package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"fmt"
	"time"
	"net/http"
	"github.com/max-las/humans-of-grenoble/models"
)

type NewStoryController struct {
	beego.Controller
}

func (c *NewStoryController) Get() {
  c.Data["PageTitle"] = "Nouvelle Story"
  c.Data["AdditionnalScripts"] = [1]string{"/static/private/scripts/new.js"}

  c.Layout = "layouts/main.tpl"
  c.TplName = "admin/new.tpl"
}

func (c *NewStoryController) Post() {
	text := c.GetString("text")
  _, header, err := c.GetFile("imageFile")
	if(text == ""){
		c.Abort("403")
	}
	if(err != nil){
		if(err == http.ErrMissingFile){
			fmt.Println("missing file")
			c.Abort("403")
		}else{
			fmt.Println(err.Error())
			c.Abort("500")
		}
	}else{
		mimeType := header.Header["Content-Type"][0]
		if(mimeType != "image/jpeg" && mimeType != "image/png" && mimeType != "image/gif"){
			c.Abort("403")
		}else{
			savePath := "static/photos/" + time.Now().Format("02012006150405") + "-" + header.Filename
			err := c.SaveToFile("imageFile", savePath);
			if(err != nil){
				fmt.Println(err.Error())
				c.Abort("500")
			}else{
				story := models.Story{
					PhotoUrl: "/" + savePath,
					Text: text,
				}

				_, err = models.AddStory(&story)
				if(err != nil){
					fmt.Println(err.Error())
					c.Abort("500")
				}

				c.Data["Message"] = "OK"
				c.TplName = "dev/simpleMessage.tpl"
			}
		}
	}
}
