package handler

import (
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler/middleware"
)

type ShowOSHandler struct{}

func NewShowOSHandler() *ShowOSHandler {
	return &ShowOSHandler{}
}

func (h *ShowOSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	os := middleware.GetOS(r.Context())
	w.Write([]byte("OS: " + os + "\n"))
}
