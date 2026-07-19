package projects

import (
	"encoding/json"
	"net/http"

	"github.com/X0JIO/nebula-api/internal/platform/web/middleware"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service     *Service
	permissions *PermissionService
}

func NewHandler(
	service *Service,
	permissions *PermissionService,
) *Handler {

	return &Handler{
		service:     service,
		permissions: permissions,
	}
}

// CreateProject godoc
//
//	@Summary		Create project
//	@Tags			Projects
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateProjectRequest	true	"Project"
//	@Success		200		{object}	ProjectResponse
//	@Failure		400		{string}	string
//	@Failure		401		{string}	string
//	@Router			/projects [post]
func (h *Handler) CreateProject(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req CreateProjectRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		http.Error(
			w,
			"invalid request",
			http.StatusBadRequest,
		)

		return
	}

	userID := middleware.UserID(r.Context())

	project, err := h.service.CreateProject(
		r.Context(),
		req.Name,
		req.Description,
		req.Visibility,
		userID,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		ToResponse(project),
	)
}

// ListProjects godoc
//
//	@Summary		My projects
//	@Tags			Projects
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200		{array}		ProjectResponse
//	@Router			/projects [get]
func (h *Handler) ListProjects(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID := middleware.UserID(
		r.Context(),
	)

	projects, err := h.service.ListProjects(
		r.Context(),
		userID,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		ToResponses(projects),
	)
}

// GetProject godoc
//
//	@Summary		Get project
//	@Tags			Projects
//	@Security		BearerAuth
//	@Produce		json
//	@Param			id	path		string	true	"Project ID"
//	@Success		200	{object}	ProjectResponse
//	@Failure		404	{string}	string
//	@Router			/projects/{id} [get]
func (h *Handler) GetProject(
	w http.ResponseWriter,
	r *http.Request,
) {

	projectID := chi.URLParam(
		r,
		"id",
	)

	project, err := h.service.GetProject(
		r.Context(),
		projectID,
	)

	if err != nil {

		http.Error(
			w,
			"project not found",
			http.StatusNotFound,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		ToResponse(project),
	)
}

// UpdateProject godoc
//
//	@Summary		Update project
//	@Tags			Projects
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string					true	"Project ID"
//	@Param			request	body		UpdateProjectRequest	true	"Project"
//	@Success		200		{object}	ProjectResponse
//	@Failure		400		{string}	string
//	@Failure		404		{string}	string
//	@Router			/projects/{id} [put]
func (h *Handler) UpdateProject(
	w http.ResponseWriter,
	r *http.Request,
) {

	projectID := chi.URLParam(
		r,
		"id",
	)

	projectUUID, err := ParseUUID(projectID)
	if err != nil {
		http.Error(w, "invalid project id", http.StatusBadRequest)
		return
	}

	userUUID, err := ParseUUID(
		middleware.UserID(r.Context()),
	)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusUnauthorized)
		return
	}

	if err := h.permissions.CanUpdateProject(
		r.Context(),
		projectUUID,
		userUUID,
	); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	var req UpdateProjectRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		http.Error(
			w,
			"invalid request",
			http.StatusBadRequest,
		)

		return
	}

	project, err := h.service.UpdateProject(
		r.Context(),
		projectID,
		req.Name,
		req.Description,
		req.Visibility,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		ToResponse(project),
	)
}

// DeleteProject godoc
//
//	@Summary		Delete project
//	@Tags			Projects
//	@Security		BearerAuth
//	@Param			id	path	string	true	"Project ID"
//	@Success		204
//	@Failure		404	{string}	string
//	@Router			/projects/{id} [delete]
func (h *Handler) DeleteProject(
	w http.ResponseWriter,
	r *http.Request,
) {

	projectID := chi.URLParam(
		r,
		"id",
	)

	projectUUID, err := ParseUUID(projectID)
	if err != nil {
		http.Error(w, "invalid project id", http.StatusBadRequest)
		return
	}

	userUUID, err := ParseUUID(
		middleware.UserID(r.Context()),
	)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusUnauthorized)
		return
	}

	if err := h.permissions.CanDeleteProject(
		r.Context(),
		projectUUID,
		userUUID,
	); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	if err := h.service.DeleteProject(
		r.Context(),
		projectID,
	); err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)

		return
	}

	w.WriteHeader(
		http.StatusNoContent,
	)
}

// AddMember godoc
//
//	@Summary		Add member
//	@Tags			Projects
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string					true	"Project ID"
//	@Param			request	body	AddMemberRequest		true	"Member"
//	@Success		204
//	@Failure		400	{string}	string
//	@Router			/projects/{id}/members [post]
func (h *Handler) AddMember(
	w http.ResponseWriter,
	r *http.Request,
) {

	projectID := chi.URLParam(
		r,
		"id",
	)

	projectUUID, err := ParseUUID(projectID)
	if err != nil {
		http.Error(w, "invalid project id", http.StatusBadRequest)
		return
	}

	userUUID, err := ParseUUID(
		middleware.UserID(r.Context()),
	)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusUnauthorized)
		return
	}

	if err := h.permissions.CanManageMembers(
		r.Context(),
		projectUUID,
		userUUID,
	); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	var req AddMemberRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		http.Error(
			w,
			"invalid request",
			http.StatusBadRequest,
		)

		return
	}

	if err := h.service.AddMember(
		r.Context(),
		projectID,
		req.UserID,
		req.Role,
	); err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.WriteHeader(
		http.StatusNoContent,
	)
}

// RemoveMember godoc
//
//	@Summary		Remove member
//	@Tags			Projects
//	@Security		BearerAuth
//	@Param			id		path	string	true	"Project ID"
//	@Param			userId	path	string	true	"User ID"
//	@Success		204
//	@Failure		400	{string}	string
//	@Router			/projects/{id}/members/{userId} [delete]
func (h *Handler) RemoveMember(
	w http.ResponseWriter,
	r *http.Request,
) {

	projectID := chi.URLParam(
		r,
		"id",
	)

	projectUUID, err := ParseUUID(projectID)
	if err != nil {
		http.Error(w, "invalid project id", http.StatusBadRequest)
		return
	}

	userUUID, err := ParseUUID(
		middleware.UserID(r.Context()),
	)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusUnauthorized)
		return
	}

	if err := h.permissions.CanManageMembers(
		r.Context(),
		projectUUID,
		userUUID,
	); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	userID := chi.URLParam(
		r,
		"userId",
	)

	if err := h.service.RemoveMember(
		r.Context(),
		projectID,
		userID,
	); err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.WriteHeader(
		http.StatusNoContent,
	)
}
