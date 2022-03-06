package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/pagination"
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

	etag := helpers.TplLastModifiedString(c.TplName)

	const nbColumns = 4
	const storiesPerPage = 8

	var columns [nbColumns][]models.Story

	cnt, err := models.CountStory()
	if(err != nil){
		fmt.Println(err.Error())
		c.Abort("500")
	}

	if(cnt == 0){

		c.Data["NoStory"] = true

		etag = fmt.Sprintf("%s.%d", etag, 0)

	}else{

		paginator := pagination.SetPaginator(c.Ctx, storiesPerPage, int64(cnt))

		stories, err := models.GetAllStory(nil, nil, []string{"id"}, []string{"desc"}, int64(paginator.Offset()), storiesPerPage)
		if(err != nil){
			if(err != orm.ErrNoRows){
				fmt.Println(err.Error())
				c.Abort("500")
			}
		}

		etag = fmt.Sprintf("%s.%d", etag, helpers.StructsToCrc32(stories))

		for i := 0; i < nbColumns; i++ {
			for j := i; j < len(stories); j = j + nbColumns {
				columns[i] = append(columns[i], stories[j].(models.Story))
			}
		}

		c.Data["Columns"] = columns

	}

	helpers.HandleEtag(&c.Controller, etag)
}