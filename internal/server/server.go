package server

import (
	"fmt"
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

	con, err := db.Con()
	if err != nil {
		return fmt.Errorf("GORM Connection Failed: %s", err)
	}

	// Run AutoMigrate based on model
	if err := db.Migration(); err != nil {
		return fmt.Errorf("GORM Migration failed %s", err)
	}
	switch {
	case s.config.AppAdmin.Enable:
		if err := appUser(s, con); err != nil {
			return fmt.Errorf("App User creation failed %s", err)
		}
	default:
		lg.Info("Skip application user creation")
	}

	// Start server
	lg.Info("Starting server:", s.listenAddr)

	mux := router.CreateRouter(con)
	if err := http.ListenAndServe(s.listenAddr, mux); err != nil {
		return fmt.Errorf("Failed to start server %s", err)
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
		return fmt.Errorf("Failed to create user")
	}

	return nil
}
