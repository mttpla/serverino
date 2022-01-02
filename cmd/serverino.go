//Copyright 2021 Matteo Paoli - mttpla@gmail.com

package main

import (
	"github.com/mttpla/serverino/pkg"
	"github.com/mttpla/serverino/pkg/util"
)

func main() {
	util.LoadConfig()
	util.LogInit()
	util.Info.Println("Serverino: " + util.ServerVersion() + ".")
	pkg.Server()
}
