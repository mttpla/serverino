//Copyright 2021 Matteo Paoli - mttpla@gmail.com

package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/mttpla/serverino/pkg/utils"
	"net/http"
)

//ServerInfo info printed in / page
type ServerInfo struct {
	Name    string `json:",omitempty"`
	Surname string `json:",omitempty"`
}

var serverInfo = ServerInfo{Name: "Matteo", Surname: "Paoli"}

//HomeHandler main path
func HomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	serverInfoJSON, _ := json.Marshal(serverInfo)
	w.Write([]byte(string(serverInfoJSON)))
}

//NotFound manage not found path
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	util.Info.Println(fmt.Sprintf("Page Not Found: %s", r.URL))
	var msg = make(map[string]string)
	msg["Error"] = "Page not found"
	msg["URL"] = fmt.Sprintf("%s", r.URL)
	msgJSON, _ := json.Marshal(msg)
	w.Write([]byte(string(msgJSON)))
}