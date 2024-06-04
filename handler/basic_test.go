package handler

import "net/http"

type BasicTestHandler struct{}

func NewBasicTestHandler() *BasicTestHandler {
	return &BasicTestHandler{}
}

func (h *BasicTestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Authentication Success!"))
}
