package main

import (
	_ "github.com/max-las/humans-of-grenoble/routers"
	_ "github.com/lib/pq"
	_ "github.com/beego/beego/v2/server/web/session/postgres"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"

	"strings"
	"strconv"
	"os"
	"fmt"
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
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "sessionID"
	// beego.BConfig.WebConfig.Session.SessionProvider = "file"
	// beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	beego.BConfig.WebConfig.Session.SessionProvider = "postgresql"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = os.Getenv("DATABASE_URL")
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 5256000
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 5256000

	beego.BConfig.MaxUploadSize = 1000000 // 1 MB

	orm.RunSyncdb("default", false, true)

	o := orm.NewOrm()
	var to_regclass string
	err := o.Raw("SELECT to_regclass('session');").QueryRow(&to_regclass)
	if(err != nil){
		fmt.Println(err);
	}else{
		if(to_regclass == ""){
			fmt.Println("Table 'session' does not exist, creating");
			_, err := o.Raw("CREATE TABLE session ( session_key char(64) NOT NULL, session_data bytea, session_expiry timestamp NOT NULL, CONSTRAINT session_key PRIMARY KEY(session_key) );").Exec()
			if(err != nil){
				fmt.Println(err);
			}
		}else{
			fmt.Println("Table 'session' already exists");
		}
	}

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

	fileReader, err := os.Open("static/css/style.css")
	if err != nil {
		fmt.Println(err)
	}else{
		fileWriter, err := os.Create("static/css/style_minified.css")
		if err != nil {
			fmt.Println(err)
		}else{
			m := minify.New()
			m.AddFunc("text/css", css.Minify)
			err = m.Minify("text/css", fileWriter, fileReader)
			if err != nil {
				fmt.Println(err)
			}
		}
	}


	beego.Run()
}
