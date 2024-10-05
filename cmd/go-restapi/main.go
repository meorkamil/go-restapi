package main

import (
	"go-restapi/internal/server"
)

func main() {

	server := server.NewAPIServer("0.0.0.0:5001")
	server.Run()
}
