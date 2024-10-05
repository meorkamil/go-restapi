package server

import (
	"go-restapi/internal/router"
	"go-restapi/internal/store"
	"log"
	"net/http"
)

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}

}

func (s *APIServer) Run() {

	mux := router.CreateRouter()

	if err := store.InitDB("postgresql://postgres:xx@192.168.0.40/xx?sslmode=disable"); err != nil {

		log.Fatal("Init DB failed !!!")
	}

	log.Println("Starting server:", s.listenAddr)

	if err := http.ListenAndServe(s.listenAddr, mux); err != nil {

		log.Fatal("Failed to start server", err)

	}

}
