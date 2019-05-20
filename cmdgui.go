package main

import (
	"github.com/wailovet/cmdgui/app"
	"github.com/wailovet/osmanthuswine"
	"github.com/wailovet/osmanthuswine/src/core"
	"log"
	"net/http"
)

func main() {
	core.SetConfig(&core.Config{
		Port:         "2880",
		Host:         "127.0.0.1",
		StaticRouter: "/T/*",
	})

	//go-bindata-assetfs static/...
	osmanthuswine.GetChiRouter().Handle("/*", http.FileServer(assetFS()))
	//osmanthuswine.GetChiRouter().Handle("/*", http.FileServer(http.Dir("../static/")))
	log.Println("Open the http://127.0.0.1:2880/")

	core.GetInstanceRouterManage().Registered(&app.Cmd{})
	osmanthuswine.Run()
}
