package util

import (
	"encoding/json"
	"fmt"
)

//ServerInfo info printed in / page
type ServerInfo struct {
	Name           string `json:",omitempty"`
	Author         string `json:",omitempty"`
	CommitHash     string `json:",omitempty"`
	BuildTimestamp string `json:",omitempty"`
	GitDescribe    string `json:",omitempty"`
}

var (
	serverInfo ServerInfo
)

func init() {
	serverInfo = ServerInfo{
		Name:           "Serverino",
		Author:         "Matteo Paoli",
		CommitHash:     "n/a",
		BuildTimestamp: "n/a",
		GitDescribe:    "n/a",
	}
}

func ServerInfoJson() string {
	serverInfoJson, _ := json.Marshal(serverInfo)
	return string(serverInfoJson)
}

func ServerVersion() string {
	return fmt.Sprintf("%s (%s)", serverInfo.GitDescribe, serverInfo.BuildTimestamp)
}
