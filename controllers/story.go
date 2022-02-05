package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/max-las/humans-of-grenoble/models"
	"github.com/max-las/humans-of-grenoble/helpers"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"fmt"
)

type StoryController struct {
	beego.Controller
}

func (c *StoryController) Get() {
	c.Layout = "layouts/main.tpl"
  c.TplName = "story.tpl"

	etag := helpers.TplLastModifiedString(c.TplName)

  id, err := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	if(err != nil){
		c.Abort("404")
	}

	story, err := models.GetStoryById(id)
	if(err != nil){
		if(err == orm.ErrNoRows){
			c.Abort("404")
		}else{
			c.Abort("500")
		}
	}

	c.Data["PageTitle"] = helpers.FirstWords(story.Text, 3) + " | Humans of Grenoble"
	c.Data["PhotoUrl"] = story.PhotoUrl
	c.Data["Text"] = story.Text

	c.Ctx.Output.Header("ETag", fmt.Sprintf("\"%s.%d\"", etag, helpers.StructToCrc32(story)))
}
