package server

import (
	"go-restapi/internal/model"
	"go-restapi/internal/router"
	"go-restapi/internal/store"
	"log"
	"net/http"
	"strconv"
)

type APIServer struct {
	listenAddr string
	dbDsn      string
}

func NewAPIServer(cfg *model.Config) *APIServer {

	listenAddr := cfg.Server.Addr + ":" + strconv.Itoa(cfg.Server.Port)
	dbDsn := "postgresql://" + cfg.Database.User + ":" + cfg.Database.Pass + "@" + cfg.Database.Host + "/" + cfg.Database.DBName + cfg.Database.DBFlags

	return &APIServer{
		listenAddr: listenAddr,
		dbDsn:      dbDsn,
	}

}

func (s *APIServer) Run() {

	mux := router.CreateRouter()

	if err := store.InitDB(s.dbDsn); err != nil {

		log.Fatal("Init DB failed !!!")
	}

	log.Println("Starting server:", s.listenAddr)

	if err := http.ListenAndServe(s.listenAddr, mux); err != nil {

		log.Fatal("Failed to start server", err)

	}

}
