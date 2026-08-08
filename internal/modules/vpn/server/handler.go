package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
)

type Handler struct {
	service *Service
}

func NewHandler(
	service *Service,
) *Handler {
	return &Handler{
		service: service,
	}
}

type createServerRequest struct {
	Name       string `json:"name"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Country    string `json:"country"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
	ShortID    string `json:"short_id"`
}

type updateServerRequest struct {
	Name       string `json:"name"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Country    string `json:"country"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
	ShortID    string `json:"short_id"`
}

type serverResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Country   string `json:"country"`
	PublicKey string `json:"public_key,omitempty"`
	Status    string `json:"status"`
	Capacity  int    `json:"capacity"`
	CreatedAt string `json:"created_at"`
}

func (h *Handler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req createServerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(
			w,
			"invalid request",
			http.StatusBadRequest,
		)
		return
	}

	server, err := h.service.Create(
		r.Context(),
		CreateRequest{
			Name:       req.Name,
			Host:       req.Host,
			Port:       req.Port,
			Country:    req.Country,
			PublicKey:  req.PublicKey,
			PrivateKey: req.PrivateKey,
			ShortID:    req.ShortID,
		},
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	writeJSON(
		w,
		http.StatusCreated,
		toResponse(server),
	)
}

func (h *Handler) List(
	w http.ResponseWriter,
	r *http.Request,
) {
	servers, err := h.service.List(
		r.Context(),
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	resp := make(
		[]serverResponse,
		0,
		len(servers),
	)

	for _, server := range servers {
		resp = append(
			resp,
			toResponse(server),
		)
	}

	writeJSON(
		w,
		http.StatusOK,
		resp,
	)
}

func (h *Handler) Get(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := chi.URLParam(r, "id")

	server, err := h.service.Get(
		r.Context(),
		id,
	)

	if err != nil {
		http.Error(
			w,
			"server not found",
			http.StatusNotFound,
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		toResponse(server),
	)
}

func (h *Handler) Update(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := chi.URLParam(r, "id")

	var req updateServerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(
			w,
			"invalid request",
			http.StatusBadRequest,
		)
		return
	}

	server, err := h.service.Update(
		r.Context(),
		id,
		UpdateRequest{
			Name:       req.Name,
			Host:       req.Host,
			Port:       req.Port,
			Country:    req.Country,
			PublicKey:  req.PublicKey,
			PrivateKey: req.PrivateKey,
			ShortID:    req.ShortID,
		},
	)

	if err != nil {
		if err.Error() == "invalid server id" {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)
			return
		}

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		toResponse(server),
	)
}

func (h *Handler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := chi.URLParam(r, "id")

	if err := h.service.Delete(
		r.Context(),
		id,
	); err != nil {
		if err.Error() == "invalid server id" {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)
			return
		}

		http.Error(
			w,
			"server not found",
			http.StatusNotFound,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) Activate(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := chi.URLParam(r, "id")

	server, err := h.service.Activate(
		r.Context(),
		id,
	)

	if err != nil {
		if err.Error() == "invalid server id" {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)
			return
		}

		http.Error(
			w,
			"server not found",
			http.StatusNotFound,
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		toResponse(server),
	)
}

func (h *Handler) Deactivate(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := chi.URLParam(r, "id")

	server, err := h.service.Deactivate(
		r.Context(),
		id,
	)

	if err != nil {
		if err.Error() == "invalid server id" {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)
			return
		}

		http.Error(
			w,
			"server not found",
			http.StatusNotFound,
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		toResponse(server),
	)
}

func (h *Handler) DeactivateAll(
	w http.ResponseWriter,
	r *http.Request,
) {
	if err := h.service.DeactivateAll(
		r.Context(),
	); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func toResponse(
	server db.VpnServer,
) serverResponse {
	response := serverResponse{
		ID:       server.ID.String(),
		Name:     server.Name,
		Host:     server.Host,
		Port:     int(server.Port),
		Country:  server.Country,
		Status:   server.Status,
		Capacity: int(server.Capacity),
		CreatedAt: server.CreatedAt.Time.Format(
			"2006-01-02T15:04:05Z07:00",
		),
	}

	if server.PublicKey.Valid {
		response.PublicKey = server.PublicKey.String
	}

	return response
}

func writeJSON(
	w http.ResponseWriter,
	status int,
	value any,
) {
	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(value)
}

func (h *Handler) Active(
	w http.ResponseWriter,
	r *http.Request,
) {
	server, err := h.service.GetActive(r.Context())
	if err != nil {
		http.Error(
			w,
			"active server not found",
			http.StatusNotFound,
		)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		toResponse(server),
	)
}
