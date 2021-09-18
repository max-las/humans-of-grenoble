package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"

	"fmt"
	"strconv"
	"strings"

	"github.com/max-las/humans-of-grenoble/models"
)

type EditStoryController struct {
	beego.Controller
}

func (c *EditStoryController) Prepare() {
	id, err := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	if(err != nil){
		c.Abort("404")
	}else{
		story, err := models.GetStoryById(id)
		if(err == orm.ErrNoRows){
			c.Abort("404")
		}else{
			c.Data["Story"] = story
			splitPhotoUrl := strings.Split(story.PhotoUrl, "/")
			c.Data["PreviousFileName"] = splitPhotoUrl[len(splitPhotoUrl)-1]
		}
	}
}

func (c *EditStoryController) Get() {
	c.Data["PageTitle"] = "Édition | Humans of Grenoble"

  c.Layout = "layouts/main.tpl"
  c.TplName = "admin/edit.tpl"
}

func (c *EditStoryController) Post() {
	text := c.GetString("text")
	photoUrl := c.GetString("photoUrl")
	photoPublicId := c.GetString("photoPublicId")
	if(text == ""){
		fmt.Println("Text missing")
		c.Abort("400")
	}

	keepPhoto := photoUrl == ""

	story := c.Data["Story"].(*models.Story)

	if(!keepPhoto){

		story.PhotoUrl = photoUrl
		story.PhotoPublicId = photoPublicId

	}

	story.Text = text

	err := models.UpdateStoryById(story)
	if(err != nil){
		fmt.Println(err.Error())
		c.Abort("500")
	}

	c.Data["Message"] = "/story/" + strconv.FormatInt(story.Id, 10)
	c.TplName = "dev/simpleMessage.tpl"

}

func (c *EditStoryController) Delete() {
	story := c.Data["Story"].(*models.Story)
	err := models.DeleteStory(story.Id)
	if(err == nil){
		c.Data["Message"] = "OK"
		c.TplName = "dev/simpleMessage.tpl"
	}else{
		c.Abort("500")
	}
}
