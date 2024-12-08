package server

import (
	"context"
	"errors"
	"fmt"
	"go-restapi/internal/database"
	"go-restapi/internal/model"
	"go-restapi/internal/router"
	"go-restapi/internal/util"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

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
		return fmt.Errorf("Server %s", err)
	}

	// Run AutoMigrate based on model
	if err := db.Migration(); err != nil {
		return fmt.Errorf("Server %s", err)
	}
	switch {
	case s.config.AppAdmin.Enable:
		if err := appUser(s, con); err != nil {
			return fmt.Errorf("Server %s", err)
		}
	default:
		lg.Info("Skip application user creation")
	}

	// Setup server with graceful shutdown
	mux := router.CreateRouter(con)

	server := &http.Server{
		Addr:    s.listenAddr,
		Handler: mux,
	}

	// Graceful shutdown
	go func() {
		lg.Info("Starting Server", server.Addr)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			lg.Fatal("Server HTTP shutdown error:", err)
		}
		lg.Info("Server stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("Server HTTP shutdown error %s", err)
	}

	lg.Info("Server Shutdown Completed.")

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
		return fmt.Errorf("Server create user %s", err)
	}

	return nil
}
