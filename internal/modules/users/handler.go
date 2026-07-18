package users

import (
	"encoding/json"
	"net/http"
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

	response := User{
		ID:        user.ID.String(),
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt.Time,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(response)
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

	_ = json.NewEncoder(w).Encode(user)
}