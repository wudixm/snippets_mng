package main

import (
	"fmt"
	"main/src/models"
	"main/src/router"
	"net/http"
)

func main() {
	r := router.InitRouter()
	r.Static("/assets", "./assets")
	r.StaticFile("/index", "./index.html")
	r.StaticFile("/favicon.ico", "./favicon.ico")


	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: r,
		//ReadTimeout:    settings.ReadTimeout,
		//WriteTimeout:   settings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	models.ConnectES()

	s.ListenAndServe()
}
