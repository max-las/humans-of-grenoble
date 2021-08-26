package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
	"github.com/max-las/humans-of-grenoble/models"
	"fmt"
)

type StoriesController struct {
	beego.Controller
}

func (c *StoriesController) Get() {
	c.Layout = "layouts/main.tpl"
	c.TplName = "stories.tpl"

	const nbColumns = 4

	var columns [nbColumns][]models.Story

	stories, err := models.GetAllStory(nil, nil, nil, nil, 0, 100)
	if(err != nil){
		if(err != orm.ErrNoRows){
			fmt.Println(err.Error())
			c.Abort("500")
		}
	}else{
		for i := 0; i < nbColumns; i++ {
			for j := i; j < len(stories); j = j + nbColumns {
				columns[i] = append(columns[i], stories[j].(models.Story))
			}
		}

		c.Data["Columns"] = columns
	}
}
