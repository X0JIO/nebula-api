package comments

import (
	"encoding/json"
	"net/http"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
	"github.com/X0JIO/nebula-api/internal/platform/web/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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

// CreateComment godoc
//
//	@Summary		Create comment
//	@Tags			Comments
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			taskId	path		string					true	"Task ID"
//	@Param			request	body		CreateCommentRequest	true	"Comment"
//	@Success		200		{object}	CommentResponse
//	@Failure		400		{string}	string
//	@Failure		401		{string}	string
//	@Router			/tasks/{taskId}/comments [post]
func (h *Handler) CreateComment(
	w http.ResponseWriter,
	r *http.Request,
) {

	taskIDStr := chi.URLParam(
		r,
		"taskId",
	)

	taskUUID, err := uuid.Parse(taskIDStr)
	if err != nil {

		http.Error(
			w,
			"invalid task id",
			http.StatusBadRequest,
		)

		return
	}

	authorIDStr := middleware.UserID(
		r.Context(),
	)

	authorUUID, err := uuid.Parse(authorIDStr)
	if err != nil {

		http.Error(
			w,
			"invalid user id",
			http.StatusUnauthorized,
		)

		return
	}

	var req CreateCommentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		http.Error(
			w,
			"invalid request",
			http.StatusBadRequest,
		)

		return
	}

	comment, err := h.service.Create(
		r.Context(),
		db.CreateCommentParams{
			TaskID: pgtype.UUID{
				Bytes: taskUUID,
				Valid: true,
			},
			AuthorID: pgtype.UUID{
				Bytes: authorUUID,
				Valid: true,
			},
			Body: req.Body,
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

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		ToResponse(comment),
	)
}

// ListTaskComments godoc
//
//	@Summary		List task comments
//	@Tags			Comments
//	@Security		BearerAuth
//	@Produce		json
//	@Param			taskId	path		string	true	"Task ID"
//	@Success		200		{array}		CommentResponse
//	@Router			/tasks/{taskId}/comments [get]
func (h *Handler) ListTaskComments(
	w http.ResponseWriter,
	r *http.Request,
) {

	taskIDStr := chi.URLParam(
		r,
		"taskId",
	)

	taskUUID, err := uuid.Parse(taskIDStr)
	if err != nil {

		http.Error(
			w,
			"invalid task id",
			http.StatusBadRequest,
		)

		return
	}

	comments, err := h.service.ListTask(
		r.Context(),
		pgtype.UUID{
			Bytes: taskUUID,
			Valid: true,
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

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		ToResponses(comments),
	)
}

// DeleteComment godoc
//
//	@Summary		Delete comment
//	@Tags			Comments
//	@Security		BearerAuth
//	@Param			id	path	string	true	"Comment ID"
//	@Success		204
//	@Router			/comments/{id} [delete]
func (h *Handler) DeleteComment(
	w http.ResponseWriter,
	r *http.Request,
) {

	idStr := chi.URLParam(
		r,
		"id",
	)

	commentUUID, err := uuid.Parse(idStr)
	if err != nil {

		http.Error(
			w,
			"invalid comment id",
			http.StatusBadRequest,
		)

		return
	}

	err = h.service.Delete(
		r.Context(),
		pgtype.UUID{
			Bytes: commentUUID,
			Valid: true,
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

	w.WriteHeader(
		http.StatusNoContent,
	)
}
