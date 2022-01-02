//Copyright 2021 Matteo Paoli - mttpla@gmail.com

package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mttpla/serverino/pkg"
	"github.com/mttpla/serverino/pkg/util"
	"log"
	"net/http"
)

func main() {
	util.LoadConfig()
	util.LogInit()
	router := httprouter.New()
	router.GET("/", pkg.HomeHandler)
	router.NotFound = http.HandlerFunc(pkg.NotFound)
	util.Info.Println("Serverino: " + util.ServerVersion() + ".")
	util.Info.Println("Started on port " + util.Configuration.Port + ".")
	log.Fatal(http.ListenAndServe(":"+util.Configuration.Port, router))
}
