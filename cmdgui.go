package main

import (
	"github.com/wailovet/osmanthuswine"
	"github.com/wailovet/osmanthuswine/src/core"
	"log"
	"net/http"
)

func main() {
	core.SetConfig(&core.Config{
		Port:      "2880",
		ApiRouter: "/Api/*",
	})

	//go-bindata-assetfs static/...
	osmanthuswine.GetChiRouter().Handle("/*", http.FileServer(assetFS()))
	log.Println("Open the http://127.0.0.1:2880/")
	osmanthuswine.Run()
}
