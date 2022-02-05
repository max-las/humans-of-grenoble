package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/utils/pagination"
	"github.com/max-las/humans-of-grenoble/models"
	"github.com/max-las/humans-of-grenoble/helpers"
	"fmt"
)

type StoriesController struct {
	beego.Controller
}

func (c *StoriesController) Get() {
	c.Data["PageTitle"] = "Stories | Humans of Grenoble"
	c.Layout = "layouts/main.tpl"
	c.TplName = "stories.tpl"

	const nbColumns = 4
	const storiesPerPage = 8

	var columns [nbColumns][]models.Story

	stories, err := models.GetAllStory(nil, nil, []string{"id"}, []string{"desc"}, 0, 100)
	if(err != nil){
		if(err != orm.ErrNoRows){
			fmt.Println(err.Error())
			c.Abort("500")
		}else{
			c.Data["NoStory"] = true
		}
	}

	if(len(stories) == 0){

		c.Data["NoStory"] = true

	}else{

		paginator := pagination.NewPaginator(c.Ctx.Request, storiesPerPage, len(stories))
		c.Data["paginator"] = paginator

		for i := 0; i < nbColumns; i++ {
			for j := i; j < helpers.MinInt(storiesPerPage, len(stories)-paginator.Offset()); j = j + nbColumns {
				columns[i] = append(columns[i], stories[helpers.MinInt(len(stories)-1, paginator.Offset()+j)].(models.Story))
			}
		}

		c.Data["Columns"] = columns

	}

}
