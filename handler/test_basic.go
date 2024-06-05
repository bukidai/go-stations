package handler

import "net/http"

type TestBasicHandler struct{}

func NewTestBasicHandler() *TestBasicHandler {
	return &TestBasicHandler{}
}

func (h *TestBasicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Authentication Success!\n"))
}
