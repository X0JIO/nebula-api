package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net"
	"net/http"
	"time"

	"github.com/X0JIO/nebula-api/internal/platform/web"
	"github.com/X0JIO/nebula-api/internal/platform/web/middleware"
)

type Handler struct {
	service *Service
}

type logoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type registerResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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
//	@Success		200		{object}	registerResponse
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

	resp := registerResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		Status:    user.Status,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Time.Format(time.RFC3339),
	}

	web.Created(w, resp)
}

// @Summary		Login
// @Description	Authenticate user
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			X-Device-Fingerprint	header		string			true	"Unique device fingerprint"
// @Param			X-Device-Name			header		string			false	"Device name"
// @Param			request					body		loginRequest	true	"Login request"
// @Success		200						{object}	Tokens
// @Failure		401						{string}	string
// @Router			/auth/login [post]
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

	device := r.Header.Get("X-Device-Name")

	if device == "" {
		device = "Unknown Device"
	}

	fingerprint := r.Header.Get("X-Device-Fingerprint")

	if fingerprint == "" {
		web.Error(
			w,
			http.StatusBadRequest,
			"missing X-Device-Fingerprint header",
		)
		return
	}

	ip := ""

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		ip = host
	} else {
		ip = r.RemoteAddr
	}

	tokens, err := h.service.LoginTokens(
		r.Context(),
		req.Email,
		req.Password,
		device,
		ip,
		r.UserAgent(),
		fingerprint,
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

// Logout godoc
//
//	@Summary		Logout
//	@Description	Logout current session
//	@Tags			Auth
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		logoutRequest	true	"Logout request"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{string}	string
//	@Failure		401		{string}	string
//	@Router			/auth/logout [post]
func (h *Handler) Logout(
	w http.ResponseWriter,
	r *http.Request,
) {
	sessionID := r.Context().Value(middleware.ContextSessionID).(string)

	var req logoutRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		web.Error(
			w,
			http.StatusBadRequest,
			"invalid request",
		)
		return
	}

	if req.RefreshToken == "" {
		web.Error(
			w,
			http.StatusBadRequest,
			"missing refresh token",
		)
		return
	}

	hash := sha256.Sum256([]byte(req.RefreshToken))
	refreshHash := hex.EncodeToString(hash[:])

	if err := h.service.Logout(
		r.Context(),
		sessionID,
		refreshHash,
	); err != nil {
		web.WriteError(w, err)
		return
	}

	web.OK(w, map[string]string{
		"status": "ok",
	})
}

// LogoutAll godoc
//
//	@Summary		Logout from all devices
//	@Description	Revoke all user sessions and refresh tokens
//	@Tags			Auth
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Failure		401	{string}	string
//	@Router			/auth/logout-all [post]
func (h *Handler) LogoutAll(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID := r.Context().Value(middleware.ContextUserID).(string)

	if err := h.service.LogoutAll(
		r.Context(),
		userID,
	); err != nil {
		web.WriteError(w, err)
		return
	}

	web.OK(w, map[string]string{
		"status": "ok",
	})
}
