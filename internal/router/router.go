package router

import (
	"go-restapi/internal/handler"
	"net/http"
)

func CreateRouter() *http.ServeMux {

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(handler.DefaultRouter))
	mux.Handle("POST /employee", http.HandlerFunc(handler.GetEmployeeById))
	mux.Handle("GET /employee", http.HandlerFunc(handler.ListEmployee))

	return mux

}
