package main

import (
	"flag"
	"go-restapi/internal/server"
	"go-restapi/internal/util"
)

var version = "v1.0.2"
var lg = util.NewLogger()

func main() {

	lg.Info("go-restapi:", version)

	configPath := flag.String("config", "../../config/config.yml", "full path to configuration file")
	flag.Parse()

	config := util.ConfigInit(*configPath)

	server := server.NewAPIServer(config)
	if err := server.Run(); err != nil {
		lg.Fatal("Fail to start server", err)
	}
}
