package server

import (
	"go-restapi/internal/database"
	"go-restapi/internal/model"
	"go-restapi/internal/router"
	"go-restapi/internal/util"
	"net"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

var lg = util.NewLogger()

type APIServer struct {
	listenAddr string
	config     *model.Config
}

func NewAPIServer(cfg *model.Config) *APIServer {
	// Return APIServer sturct
	return &APIServer{
		listenAddr: net.JoinHostPort(
			cfg.Server.Addr,
			strconv.Itoa(cfg.Server.Port),
		),
		config: cfg,
	}
}

func (s *APIServer) Run() error {
	// Initialize database
	db := database.InitDB(s.config)
	con := db.Con()

	// Run AutoMigrate based on model
	if err := db.Migration(); err != nil {
		lg.Fatal("GORM Migration failed", err)
	}
	switch {
	case s.config.AppAdmin.Enable:
		if err := appUser(s, con); err != nil {
			lg.Fatal("App User creation failed", err)

		}
	default:
		lg.Info("Skip application user creation")
	}

	// Start server
	lg.Info("Starting server:", s.listenAddr)

	mux := router.CreateRouter(con)
	if err := http.ListenAndServe(s.listenAddr, mux); err != nil {
		lg.Fatal("Failed to start server", err)
	}

	return nil
}

// Set Admin User based on config.yml
func appUser(s *APIServer, db *gorm.DB) error {
	appUser := model.Employee{
		Empid:    "APP1",
		Name:     s.config.AppAdmin.User,
		Password: s.config.AppAdmin.Pass,
		Dept:     "Int-APP",
	}

	repo := database.InitRepo(db)
	if err := repo.CreateAppUser(&appUser); err != nil {
		lg.Fatal("Failed to create user")
	}

	return nil
}
