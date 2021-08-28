package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/max-las/humans-of-grenoble/models"
	"github.com/max-las/humans-of-grenoble/helpers"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
)

type StoryController struct {
	beego.Controller
}

func (c *StoryController) Get() {
	c.Layout = "layouts/main.tpl"
  c.TplName = "story.tpl"

  id, err := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	if(err != nil){
		c.Abort("404")
	}

	storyJustAdded, ok := c.GetSession("storyJustAdded").(int64)
	if(ok && storyJustAdded == id){
		c.Data["JustAdded"] = true
		c.DelSession("storyJustAdded")
	}

	story, err := models.GetStoryById(id)
	if(err != nil){
		if(err == orm.ErrNoRows){
			c.Abort("404")
		}else{
			c.Abort("500")
		}
	}

	c.Data["PageTitle"] = helpers.FirstWords(story.Text, 3)
	c.Data["PhotoUrl"] = story.PhotoUrl
	c.Data["Text"] = story.Text
}
