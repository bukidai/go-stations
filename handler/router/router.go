package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/handler/middleware"
	"github.com/TechBowl-japan/go-stations/service"
	"github.com/mileusna/useragent"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", handler.NewHealthzHandler().ServeHTTP)
	mux.HandleFunc("/todos", handler.NewTODOHandler(service.NewTODOService(todoDB)).ServeHTTP)
	mux.HandleFunc("/do-panic", middleware.Recovery(handler.NewDoPanicHandler()).ServeHTTP)
	mux.Handle("/show-header", middleware.Recovery(showHeader()))

	return mux
}

func showHeader() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ua := useragent.Parse(r.UserAgent())
		w.Write([]byte(ua.OS + "\n"))
	}
	return http.HandlerFunc(fn)
}
