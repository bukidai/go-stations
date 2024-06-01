package handler

import (
	"log"
	"net/http"
)

// A PanicHandler implements panic handler.
type DoPanicHandler struct{}

// NewDoPanicHandler returns DoPanicHandler based http.Handler.
func NewDoPanicHandler() *DoPanicHandler {
	return &DoPanicHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *DoPanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Panicln("Panic!")
}
