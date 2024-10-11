package handler

import (
	"fmt"
	"go-restapi/internal/database"
	"go-restapi/internal/model"
	"go-restapi/internal/util"
	"net/http"

	"gorm.io/gorm"
)

var lg = util.NewLogger()

// Default route
func DefaultRouter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "RESTful API")
	}
}

// GET /employee
func ListEmployee(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Init repository
		repo := database.InitRepo(db)
		employee, err := repo.FindAll()
		if err != nil {
			lg.Fatal(err)
		}

		switch {
		case len(employee) == 0:
			lg.Info("Empty employee:", employee)
		default:
			lg.Info("Employee data", employee)
		}

		if err := util.Encode(w, r, http.StatusOK, employee); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

// POST /employee
func CreateEmployee(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode payload
		employee, err := util.Decode[model.Employee](r)
		switch {
		case err != nil:
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
		default:
			repo := database.InitRepo(db)
			err := repo.CreateEmployee(&employee)

			if err != nil {
				lg.Fatal("GORM Create employee failed:", err)
			}

			// Decode back to client
			if err := util.Encode(w, r, http.StatusOK, employee); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}

		}
	}
}
