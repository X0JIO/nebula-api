package vpn

import (
	"encoding/json"
	"net/http"

	"github.com/X0JIO/nebula-api/internal/platform/web"
	"github.com/X0JIO/nebula-api/internal/platform/web/middleware"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateConfig godoc
//
//	@Summary		Create VPN config
//	@Description	Generate VPN configuration
//	@Tags			VPN
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateConfigRequest	true	"Protocol"
//	@Success		200		{object}	ConfigResponse
//	@Failure		400		{string}	string
//	@Failure		401		{string}	string
//	@Router			/vpn/config [post]
func (h *Handler) CreateConfig(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req CreateConfigRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		web.Error(
			w,
			http.StatusBadRequest,
			"invalid request",
		)
		return
	}

	v := r.Context().Value(middleware.ContextUserID)

	userID, ok := v.(string)
	if !ok || userID == "" {
		web.Error(
			w,
			http.StatusUnauthorized,
			"unauthorized",
		)
		return
	}

	cfg, err := h.service.CreateConfig(
		r.Context(),
		userID,
		req.Protocol,
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
		cfg,
	)
}

// ListConfigs godoc
//
//	@Summary		Get VPN configs
//	@Description	Get all VPN configs for current user
//	@Tags			VPN
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	ConfigResponse
//	@Failure		401	{string}	string
//	@Router			/vpn/configs [get]
func (h *Handler) ListConfigs(
	w http.ResponseWriter,
	r *http.Request,
) {

	v := r.Context().Value(middleware.ContextUserID)

	userID, ok := v.(string)
	if !ok || userID == "" {
		web.Error(
			w,
			http.StatusUnauthorized,
			"unauthorized",
		)
		return
	}

	resp, err := h.service.ListConfigs(
		r.Context(),
		userID,
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
		resp,
	)
}

// GetAccount godoc
//
//	@Summary		Get VPN account
//	@Description	Get current user VPN account
//	@Tags			VPN
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	VPNAccountResponse
//	@Failure		401	{string}	string
//	@Router			/vpn/account [get]
func (h *Handler) GetAccount(
	w http.ResponseWriter,
	r *http.Request,
) {

	v := r.Context().Value(middleware.ContextUserID)

	userID, ok := v.(string)

	if !ok || userID == "" {
		web.Error(
			w,
			http.StatusUnauthorized,
			"unauthorized",
		)
		return
	}

	account, err := h.service.GetAccount(
		r.Context(),
		userID,
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
		account,
	)
}

// Subscription godoc
//
//	@Summary		Get subscription
//	@Tags			VPN
//	@Security		BearerAuth
//	@Produce		plain
//	@Success		200	{string}	string
//	@Router			/vpn/subscription [get]
func (h *Handler) Subscription(
	w http.ResponseWriter,
	r *http.Request,
) {

	v := r.Context().Value(middleware.ContextUserID)

	userID, ok := v.(string)
	if !ok || userID == "" {
		web.Error(
			w,
			http.StatusUnauthorized,
			"unauthorized",
		)
		return
	}

	sub, err := h.service.Subscription(
		r.Context(),
		userID,
	)

	if err != nil {
		web.WriteError(
			w,
			err,
		)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	_, _ = w.Write([]byte(sub))
}

// SubscriptionBase64 godoc
//
//	@Summary		Get subscription (Base64)
//	@Tags			VPN
//	@Security		BearerAuth
//	@Produce		plain
//	@Success		200	{string}	string
//	@Router			/vpn/subscription/base64 [get]
func (h *Handler) SubscriptionBase64(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID := r.Context().Value(middleware.ContextUserID).(string)

	sub, err := h.service.SubscriptionBase64(
		r.Context(),
		userID,
	)

	if err != nil {
		web.WriteError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(sub))
}
