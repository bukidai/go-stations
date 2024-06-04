package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/handler/middleware"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB, authSetting *middleware.AuthSetting) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", handler.NewHealthzHandler().ServeHTTP)
	mux.HandleFunc("/todos", handler.NewTODOHandler(service.NewTODOService(todoDB)).ServeHTTP)
	mux.HandleFunc("/do-panic", middleware.Recovery(handler.NewDoPanicHandler()).ServeHTTP)
	mux.HandleFunc("/show-os", middleware.OSContextInjection(handler.NewShowOSHandler()).ServeHTTP)
	mux.HandleFunc("/duration", middleware.OSContextInjection(middleware.AccessLog(handler.NewDurationHandler())).ServeHTTP)
	mux.HandleFunc("/hello", handler.NewAuthTestHandler().ServeHTTP)

	return mux
}
