package server

import (
	"go-restapi/internal/database"
	"go-restapi/internal/model"
	"go-restapi/internal/router"
	"go-restapi/internal/util"
	"net"
	"net/http"
	"strconv"
)

var lg = util.NewLogger()

type APIServer struct {
	listenAddr string
	dbDsn      string
}

func NewAPIServer(cfg *model.Config) *APIServer {
	dbDsn := "postgresql://" + cfg.Database.User + ":" + cfg.Database.Pass + "@" + cfg.Database.Host + "/" + cfg.Database.DBName + cfg.Database.DBFlags

	return &APIServer{
		listenAddr: net.JoinHostPort(cfg.Server.Addr, strconv.Itoa(cfg.Server.Port)),
		dbDsn:      dbDsn,
	}
}

func (s *APIServer) Run() error {
	mux := router.CreateRouter()

	if err := database.InitDB(s.dbDsn); err != nil {

		lg.Fatal("Init DB failed !!!", err)
	}

	lg.Info("Starting server:", s.listenAddr)

	if err := http.ListenAndServe(s.listenAddr, mux); err != nil {
		lg.Fatal("Failed to start server", err)
	}

	return nil
}
