package handler

import (
	"net/http"
)

type AuthTestHandler struct{}

func NewAuthTestHandler() *AuthTestHandler {
	return &AuthTestHandler{}
}

func (h *AuthTestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Authentication Success!\n"))
}
