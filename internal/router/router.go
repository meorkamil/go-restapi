package router

import (
	"go-restapi/internal/handler"
	"go-restapi/internal/security"
	"go-restapi/internal/util"
	"net/http"

	"gorm.io/gorm"
)

var lg = util.NewLogger()

// Create router
func CreateRouter(db *gorm.DB) *http.ServeMux {
	// Create middleware return http.Handler
	middleware := newMiddleware()
	mux := http.NewServeMux()
	mux.Handle("GET /employee", middleware(handler.ListEmployee(db)))
	mux.Handle("POST /employee", middleware(handler.CreateEmployee(db)))
	mux.Handle("POST /login", handler.Login(db))

	return mux
}

// Create middleware
func newMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			switch {
			case token == "":
				lg.Info(http.StatusUnauthorized, r.Method, r.URL, "unauthorized")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			case token != "":
				token := token[len("Bearer "):]
				if err := security.ValidateToken(token); err != nil {
					lg.Info(http.StatusUnauthorized, r.Method, r.URL, "unauthorized")
					lg.Fatal(err)
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}

				next.ServeHTTP(w, r)
			default:
				lg.Info(http.StatusUnauthorized, r.Method, r.URL, "unauthorized")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		})
	}
}
