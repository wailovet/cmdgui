package main

import (
	"github.com/wailovet/cmdgui/app"
	"github.com/wailovet/osmanthuswine"
	"github.com/wailovet/osmanthuswine/src/core"
	"log"
	"net/http"
	"os"
)

func main() {
	core.SetConfig(&core.Config{
		Port:         "2880",
		Host:         "127.0.0.1",
		StaticRouter: "/T/*",
	})

	//go-bindata-assetfs static/...
	if os.Args[1] == "dev" {
		osmanthuswine.GetChiRouter().Handle("/*", http.FileServer(http.Dir("../static/")))
	} else {
		osmanthuswine.GetChiRouter().Handle("/*", http.FileServer(assetFS()))
	}
	log.Println("Open the http://127.0.0.1:2880/")

	core.GetInstanceRouterManage().Registered(&app.Cmd{})
	osmanthuswine.Run()
}
