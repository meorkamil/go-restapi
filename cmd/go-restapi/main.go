package main

import (
	"flag"
	"go-restapi/internal/server"
	"go-restapi/internal/util"
)

func main() {

	configPath := flag.String("config", "../../config/config.yml", "full path to configuration file")
	flag.Parse()
	config := util.ConfigInit(*configPath)

	server := server.NewAPIServer(config)
	server.Run()
}
