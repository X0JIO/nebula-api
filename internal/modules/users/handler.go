package users

import (
	"encoding/json"
	"net/http"

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

type createUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req createUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		http.Error(
			w,
			"invalid request",
			http.StatusBadRequest,
		)

		return
	}

	user, err := h.service.Create(
		r.Context(),
		req.Email,
		req.Password,
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
		ToResponse(user),
	)
}


func (h *Handler) Me(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID, ok := r.Context().Value("user_id").(string)

	if !ok {
		http.Error(
			w,
			"unauthorized",
			http.StatusUnauthorized,
		)
		return
	}

	user, err := h.service.GetByID(
		r.Context(),
		userID,
	)

	if err != nil {
		http.Error(
			w,
			"user not found",
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		ToResponse(user),
	)
}


func ToResponse(
	user db.User,
) UserResponse {

	return UserResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}
}