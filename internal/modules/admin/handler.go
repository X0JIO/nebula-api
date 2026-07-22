package admin

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/X0JIO/nebula-api/internal/modules/users"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

type roleRequest struct {
	Role string `json:"role"`
}

type statusRequest struct {
	Status string `json:"status"`
}

// Dashboard godoc
//
//	@Summary		Admin dashboard
//	@Description	Get dashboard statistics
//	@Tags			Admin
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	Dashboard
//	@Failure		403	{string}	string
//	@Router			/admin/dashboard [get]
func (h *Handler) Dashboard(
	w http.ResponseWriter,
	r *http.Request,
) {
	stats, err := h.service.Dashboard(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, stats)
}

// ListUsers godoc
//
//	@Summary		List users
//	@Description	Get list of users
//	@Tags			Admin
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{array}		users.UserResponse
//	@Failure		403	{string}	string
//	@Router			/admin/users [get]
func (h *Handler) ListUsers(
	w http.ResponseWriter,
	r *http.Request,
) {
	list, err := h.service.ListUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := make([]users.UserResponse, 0, len(list))

	for _, u := range list {
		resp = append(resp, users.ToResponse(u))
	}

	writeJSON(w, http.StatusOK, resp)
}

// GetUser godoc
//
//	@Summary		Get user
//	@Description	Get user by ID
//	@Tags			Admin
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	users.UserResponse
//	@Failure		403	{string}	string
//	@Failure		404	{string}	string
//	@Router			/admin/users/{id} [get]
func (h *Handler) GetUser(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := chi.URLParam(r, "id")

	user, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		users.ToResponse(user),
	)
}

// ChangeRole godoc
//
//	@Summary		Change user role
//	@Description	Change the role of a user
//	@Tags			Admin
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	users.UserResponse
//	@Failure		403	{string}	string
//	@Failure		404	{string}	string
//	@Router			/admin/users/{id}/role [put]
func (h *Handler) ChangeRole(
	w http.ResponseWriter,
	r *http.Request,
) {

	id := chi.URLParam(r, "id")

	var req roleRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(
			w,
			"invalid request",
			http.StatusBadRequest,
		)
		return
	}

	user, err := h.service.ChangeRole(
		r.Context(),
		id,
		req.Role,
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
		http.StatusOK,
		users.ToResponse(user),
	)
}

// ChangeStatus godoc
//
//	@Summary		Change user status
//	@Description	Change the status of a user
//	@Tags			Admin
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	users.UserResponse
//	@Failure		403	{string}	string
//	@Failure		404	{string}	string
//	@Router			/admin/users/{id}/status [put]
func (h *Handler) ChangeStatus(
	w http.ResponseWriter,
	r *http.Request,
) {

	id := chi.URLParam(r, "id")

	var req statusRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(
			w,
			"invalid request",
			http.StatusBadRequest,
		)
		return
	}

	user, err := h.service.ChangeStatus(
		r.Context(),
		id,
		req.Status,
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
		http.StatusOK,
		users.ToResponse(user),
	)
}

// DeleteUser godoc
//
//	@Summary		Delete user
//	@Description	Delete a user by ID
//	@Tags			Admin
//	@Security		BearerAuth
//	@Produce		json
//	@Success		204	{string}	string
//	@Failure		403	{string}	string
//	@Failure		404	{string}	string
//	@Router			/admin/users/{id} [delete]
func (h *Handler) DeleteUser(
	w http.ResponseWriter,
	r *http.Request,
) {

	id := chi.URLParam(r, "id")

	if err := h.service.DeleteUser(
		r.Context(),
		id,
	); err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(
	w http.ResponseWriter,
	status int,
	v any,
) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(v)
}
