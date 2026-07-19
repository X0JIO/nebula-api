package auth

import (
	"encoding/json"
	"net/http"

	"github.com/X0JIO/nebula-api/internal/platform/web"
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

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type refreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// Register godoc
//
//	@Summary		Register
//	@Description	Create new user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		registerRequest	true	"Register request"
//	@Success		200		{object}	users.UserResponse
//	@Failure		400		{string}	string
//	@Router			/auth/register [post]
func (h *Handler) Register(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req registerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		web.Error(
			w,
			http.StatusBadRequest,
			"invalid request",
		)

		return
	}

	user, err := h.service.Register(
		r.Context(),
		req.Email,
		req.Password,
	)

	if err != nil {

		web.WriteError(
			w,
			err,
		)

		return
	}

	web.Created(
		w,
		user,
	)
}

// Login godoc
//
//	@Summary		Login
//	@Description	Authenticate user and return JWT tokens
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		loginRequest	true	"Login request"
//	@Success		200		{object}	Tokens
//	@Failure		401		{string}	string
//	@Router			/auth/login [post]
func (h *Handler) Login(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req loginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		web.Error(
			w,
			http.StatusBadRequest,
			"invalid request",
		)
		return
	}

	tokens, err := h.service.LoginTokens(
		r.Context(),
		req.Email,
		req.Password,
	)

	if err != nil {
		web.WriteError(
			w,
			err,
		)
		return
	}

	web.OK(
		w,
		tokens,
	)
}

// Refresh godoc
//
//	@Summary		Refresh tokens
//	@Description	Get new access token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		refreshRequest	true	"Refresh request"
//	@Success		200		{object}	Tokens
//	@Failure		401		{string}	string
//	@Router			/auth/refresh [post]
func (h *Handler) Refresh(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req refreshRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		web.Error(
			w,
			http.StatusBadRequest,
			"invalid request",
		)
		return
	}

	tokens, err := h.service.Refresh(
		r.Context(),
		req.RefreshToken,
	)

	if err != nil {
		web.WriteError(
			w,
			err,
		)
		return
	}

	web.OK(
		w,
		tokens,
	)
}
