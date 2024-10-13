package handler

import (
	"go-restapi/internal/model"
	"go-restapi/internal/security"
	"go-restapi/internal/util"
	"net/http"

	"gorm.io/gorm"
)

// POST /login
func Login(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check user authenctication
		data, err := util.Decode[model.Employee](r)
		if err != nil {
			lg.Fatal(err)
		}

		auth, err := security.Auth(db, &data)
		if err != nil {
			http.Error(w, "Authentication failed", http.StatusInternalServerError)
			return
		}
		// Generate tokne
		token, err := auth.CreateToken()
		switch {
		case err != nil:
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		case token != "":
			data := &model.Token{
				Raw: token,
			}
			// Response back to client
			if err := util.Encode(w, r, http.StatusOK, data); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		default:
			data := &model.Token{
				Raw: "JWT Creation failed",
			}
			// Response back to client
			if err := util.Encode(w, r, http.StatusInternalServerError, data); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}
	}
}
