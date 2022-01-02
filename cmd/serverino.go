//Copyright 2021 Matteo Paoli - mttpla@gmail.com

package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mttpla/serverino/pkg"
	"github.com/mttpla/serverino/pkg/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	util.LogInit(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	config := util.LoadConfig()
	router := httprouter.New()
	router.GET("/", pkg.HomeHandler)
	router.NotFound = http.HandlerFunc(pkg.NotFound)
	util.Info.Println("Serverino: " + util.ServerVersion() + ".")
	util.Info.Println("Started on port " + config.PORT + ".")
	log.Fatal(http.ListenAndServe(":"+config.PORT, router))
}
