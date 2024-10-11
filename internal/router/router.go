package router

import (
	"go-restapi/internal/handler"
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
	mux.Handle("/", middleware(handler.DefaultRouter()))
	mux.Handle("GET /employee", middleware(handler.ListEmployee(db)))
	mux.Handle("POST /employee", middleware(handler.CreateEmployee(db)))

	return mux
}

// Create middleware
func newMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lg.Info(r.Method, r.URL)
			next.ServeHTTP(w, r)
		})
	}
}
