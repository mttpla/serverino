//Copyright 2021 Matteo Paoli - mttpla@gmail.com

package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/mttpla/serverino/pkg/util"
	"net/http"
	"log"	
)


//HomeHandler main path
func HomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	serverInfoJSON := util.ServerInfoJson()
	w.Write([]byte(string(serverInfoJSON)))
}

//NotFound manage not found path
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	util.Info.Println(fmt.Sprintf("Page Not Found: %s", r.URL))
	var msg = make(map[string]string)
	msg["Error"] = "Page not found"
	msg["URL"] = (r.URL).String()
	msgJSON, _ := json.Marshal(msg)
	w.Write([]byte(string(msgJSON)))
}

func setupRouter() (router *httprouter.Router){
	router = httprouter.New()
	router.GET("/", HomeHandler)
	router.NotFound = http.HandlerFunc(NotFound)
	return router
}

func Server(){
	router := setupRouter()
	util.Info.Println("Started on port " + util.Configuration.Port + ".")
	log.Fatal(http.ListenAndServe(":"+util.Configuration.Port, router))

}


