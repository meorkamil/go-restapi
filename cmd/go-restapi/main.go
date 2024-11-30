package main

import (
	"flag"
	"go-restapi/internal/server"
	"go-restapi/internal/util"
	"os"
)

var version = "v1.0.5"
var lg = util.NewLogger()

func main() {

	lg.Info("go-restapi:", version)

	// Load configuration
	configPath := flag.String("config", "../../config/config.yml", "full path to configuration file")
	flag.Parse()

	// Parse config to struct
	config := util.ConfigInit(*configPath)

	// Star Server
	server := server.NewAPIServer(config)
	if err := server.Run(); err != nil {
		lg.Fatal(err)
		os.Exit(1)
	}
}
