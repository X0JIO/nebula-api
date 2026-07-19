package tasks

import (
	"encoding/json"
	"net/http"

	"github.com/X0JIO/nebula-api/internal/platform/web/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
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

// CreateTask godoc
func (h *Handler) CreateTask(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(
			w,
			"invalid request",
			http.StatusBadRequest,
		)
		return
	}

	projectID := middleware.ProjectID(
		r.Context(),
	)

	creatorID := middleware.UserID(
		r.Context(),
	)

	var projectUUID pgtype.UUID
	if err := projectUUID.Scan(projectID); err != nil {
		http.Error(
			w,
			"invalid project id",
			http.StatusBadRequest,
		)
		return
	}

	var creatorUUID pgtype.UUID
	if err := creatorUUID.Scan(creatorID); err != nil {
		http.Error(
			w,
			"invalid user id",
			http.StatusBadRequest,
		)
		return
	}

	task, err := h.service.Create(
		r.Context(),
		projectUUID,
		creatorUUID,
		req,
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

	json.NewEncoder(w).Encode(task)
}

// GetTask godoc
func (h *Handler) GetTask(
	w http.ResponseWriter,
	r *http.Request,
) {

	id := chi.URLParam(
		r,
		"id",
	)

	var taskID pgtype.UUID

	if err := taskID.Scan(id); err != nil {
		http.Error(
			w,
			"invalid id",
			http.StatusBadRequest,
		)
		return
	}

	task, err := h.service.Get(
		r.Context(),
		taskID,
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(task)
}

// ListProjectTasks godoc
func (h *Handler) ListProjectTasks(
	w http.ResponseWriter,
	r *http.Request,
) {

	projectID := chi.URLParam(
		r,
		"projectId",
	)

	var id pgtype.UUID

	if err := id.Scan(projectID); err != nil {
		http.Error(
			w,
			"invalid project id",
			http.StatusBadRequest,
		)
		return
	}

	tasks, err := h.service.ListProject(
		r.Context(),
		id,
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

	json.NewEncoder(w).Encode(tasks)
}

// UpdateTask
func (h *Handler) UpdateTask(
	w http.ResponseWriter,
	r *http.Request,
) {

	id := chi.URLParam(
		r,
		"id",
	)

	var taskID pgtype.UUID

	if err := taskID.Scan(id); err != nil {
		http.Error(
			w,
			"invalid id",
			http.StatusBadRequest,
		)
		return
	}

	var req UpdateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(
			w,
			"invalid request",
			http.StatusBadRequest,
		)
		return
	}

	task, err := h.service.Update(
		r.Context(),
		taskID,
		req,
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	json.NewEncoder(w).Encode(task)
}

// DeleteTask
func (h *Handler) DeleteTask(
	w http.ResponseWriter,
	r *http.Request,
) {

	id := chi.URLParam(
		r,
		"id",
	)

	var taskID pgtype.UUID

	if err := taskID.Scan(id); err != nil {
		http.Error(
			w,
			"invalid id",
			http.StatusBadRequest,
		)
		return
	}

	if err := h.service.Delete(
		r.Context(),
		taskID,
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

func (h *Handler) ListAssigneeTasks(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID := chi.URLParam(
		r,
		"userId",
	)

	var id pgtype.UUID

	if err := id.Scan(userID); err != nil {

		http.Error(
			w,
			"invalid user id",
			http.StatusBadRequest,
		)

		return
	}

	tasks, err := h.service.ListAssignee(
		r.Context(),
		id,
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

	json.NewEncoder(w).Encode(tasks)
}

func (h *Handler) ListStatusTasks(
	w http.ResponseWriter,
	r *http.Request,
) {

	projectID := chi.URLParam(
		r,
		"projectId",
	)

	status := chi.URLParam(
		r,
		"status",
	)

	var id pgtype.UUID

	if err := id.Scan(projectID); err != nil {

		http.Error(
			w,
			"invalid project id",
			http.StatusBadRequest,
		)

		return
	}

	tasks, err := h.service.ListStatus(
		r.Context(),
		id,
		status,
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

	json.NewEncoder(w).Encode(tasks)
}
