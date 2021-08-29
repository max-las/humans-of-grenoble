package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"fmt"
	"time"
	"net/http"
	"strconv"
	"github.com/max-las/humans-of-grenoble/models"
)

type NewStoryController struct {
	beego.Controller
}

func (c *NewStoryController) Get() {
  c.Data["PageTitle"] = "Nouvelle Story"
  c.Data["AdditionnalScripts"] = [1]string{"/static/private/js/publish.js"}

  c.Layout = "layouts/main.tpl"
  c.TplName = "admin/new.tpl"
}

func (c *NewStoryController) Post() {
	text := c.GetString("text")
	if(text == ""){
		c.Abort("400")
	}

  _, header, err := c.GetFile("imageFile")

	if(err != nil){
		if(err == http.ErrMissingFile){
			c.Abort("400")
		}else{
			fmt.Println(err.Error())
			c.Abort("500")
		}
	}else{
		mimeType := header.Header["Content-Type"][0]
		if(mimeType != "image/jpeg"){
			c.Abort("400")
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

				id, err := models.AddStory(&story)
				if(err != nil){
					fmt.Println(err.Error())
					c.Abort("500")
				}

				c.SetSession("storyJustAdded", id)

				c.Data["Message"] = "/story/" + strconv.FormatInt(id, 10)
				c.TplName = "dev/simpleMessage.tpl"
			}
		}
	}
}
