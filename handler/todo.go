package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	todo, err := h.svc.CreateTODO(ctx, req.Subject, req.Description)
	if err != nil {
		log.Panicln(err)
	}
	return &model.CreateTODOResponse{TODO: todo}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	res, err := h.svc.ReadTODO(ctx, req.PrevID, req.Size)
	if err != nil {
		log.Panicln(err)
	}
	return &model.ReadTODOResponse{TODOs: res}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	todo, err := h.svc.UpdateTODO(ctx, int64(req.ID), req.Subject, req.Description)
	if err != nil {
		log.Panicln(err)
	}
	return &model.UpdateTODOResponse{TODO: todo}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}

// ServeHTTP implements http.Handler interface.
func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { // Create
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		req := &model.CreateTODORequest{}
		if err := json.Unmarshal(body, req); err != nil {
			log.Panicln(err)
		}
		if req.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		res, err := h.Create(r.Context(), req)
		if err != nil {
			log.Panicln(err)
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Panicln(err)
		}
	} else if r.Method == "PUT" { // Update
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		req := &model.UpdateTODORequest{}
		if err := json.Unmarshal(body, req); err != nil {
			log.Panicln(err)
		}
		if req.ID == 0 || req.Subject == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		res, err := h.Update(r.Context(), req)
		if err != nil {
			log.Panicln(err)
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Panicln(err)
		}
	} else if r.Method == "GET" { // Read
		query := r.URL.Query()
		prev_id, err := strconv.Atoi(query.Get("prev_id"))
		if err != nil {
			prev_id = 0
		}
		size, err := strconv.Atoi(query.Get("size"))
		if err != nil {
			size = 5
		}
		req := &model.ReadTODORequest{PrevID: int64(prev_id), Size: int64(size)}
		res, err := h.Read(r.Context(), req)
		if err != nil {
			log.Panicln(err)
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Panicln(err)
		}
	}
}
