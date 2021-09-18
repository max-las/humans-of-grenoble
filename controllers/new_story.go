package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"fmt"
	"strconv"
	"github.com/max-las/humans-of-grenoble/models"

)

type NewStoryController struct {
	beego.Controller
}

func (c *NewStoryController) Get() {
  c.Data["PageTitle"] = "Nouvelle story | Humans of Grenoble"

  c.Layout = "layouts/main.tpl"
  c.TplName = "admin/new.tpl"
}

func (c *NewStoryController) Post() {
	text := c.GetString("text")
	photoUrl := c.GetString("photoUrl")
	photoPublicId := c.GetString("photoPublicId")
	if(text == "" || photoUrl == "" || photoPublicId == ""){
		c.Abort("400")
	}

	story := models.Story{
		PhotoUrl: photoUrl,
		PhotoPublicId: photoPublicId,
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
