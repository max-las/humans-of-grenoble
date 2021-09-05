package main

import (
	_ "github.com/max-las/humans-of-grenoble/routers"
	_ "github.com/lib/pq"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"

	"strings"
	"strconv"
	"os"
)

func init() {
	var err error

	beego.BConfig.Listen.HTTPPort, err = strconv.Atoi(os.Getenv("PORT"))
	if(err != nil){
		beego.BConfig.Listen.HTTPPort = 8080
	}

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", os.Getenv("DATABASE_URL"))

}

func main() {
	orm.RunSyncdb("default", false, true)

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "sessionID"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 15768000
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 15768000

	beego.BConfig.MaxUploadSize = 1000000 // 1 MB

	var FilterStatic = func(ctx *context.Context) {
		sess, _ := beego.GlobalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
    defer sess.SessionRelease(nil, ctx.ResponseWriter)
		username := sess.Get(nil, "username")
		if(username == nil){
			ctx.Abort(404, "404")
		}
	}

	var FilterAuth = func(ctx *context.Context) {
		url := strings.TrimSuffix(ctx.Input.URL(), "/")
		if(url != "/admin/login" && url != "/admin"){
			username := ctx.Input.Session("username")
			if(username == nil){
				ctx.Abort(404, "404")
			}
		}
	}

	beego.InsertFilter("/static/private/*", beego.BeforeStatic, FilterStatic)
	beego.InsertFilter("/admin/*", beego.BeforeRouter, FilterAuth)

	beego.Run()
}
