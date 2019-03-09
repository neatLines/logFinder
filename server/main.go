package main

import (
	"net/http"
	"strings"

	"github.com/astaxie/beego/context"

	_ "github.com/neatLines/logFinder/server/routers"

	"github.com/astaxie/beego"
)

func main() {
	var TransparentStatic = func(ctx *context.Context) {
		if strings.Index(ctx.Request.URL.Path, "v1/") >= 0 {
			return
		}
		http.ServeFile(ctx.ResponseWriter, ctx.Request, "web-page/dist/"+ctx.Request.URL.Path)
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
	beego.Run()
}
