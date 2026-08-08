package vpn

import (
	"encoding/json"
	"net/http"

	"github.com/X0JIO/nebula-api/internal/platform/web"
	"github.com/X0JIO/nebula-api/internal/platform/web/middleware"
	"github.com/go-chi/chi/v5"
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
		req.DeviceID,
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

func (h *Handler) DeleteAccount(
	w http.ResponseWriter,
	r *http.Request,
) {

	v := r.Context().Value(
		middleware.ContextUserID,
	)

	userID, ok := v.(string)

	if !ok || userID == "" {
		web.Error(
			w,
			http.StatusUnauthorized,
			"unauthorized",
		)
		return
	}

	err := h.service.DeleteAccount(
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
		DeleteVPNResponse{
			Status: "deleted",
		},
	)
}

// PublicSubscription godoc
//
//	@Summary		Get public VPN subscription
//	@Description	Get subscription by token
//	@Tags			VPN
//	@Produce		plain
//	@Param			token	path	string	true	"Subscription token"
//	@Success		200		{string}	string
//	@Router			/subscription/{token} [get]
func (h *Handler) PublicSubscription(
	w http.ResponseWriter,
	r *http.Request,
) {

	token := chi.URLParam(r, "token")

	if token == "" {
		web.Error(
			w,
			http.StatusBadRequest,
			"token required",
		)
		return
	}

	sub, err := h.service.PublicSubscription(
		r.Context(),
		token,
	)

	if err != nil {
		web.WriteError(
			w,
			err,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"text/plain",
	)

	_, _ = w.Write([]byte(sub))
}

// PublicSubscriptionBase64 godoc
//
//	@Summary		Get public VPN subscription Base64
//	@Tags			VPN
//	@Produce		plain
//	@Param			token	path	string	true	"Subscription token"
//	@Success		200		{string}	string
//	@Router			/subscription/{token}/base64 [get]
func (h *Handler) PublicSubscriptionBase64(
	w http.ResponseWriter,
	r *http.Request,
) {

	token := chi.URLParam(r, "token")

	if token == "" {
		web.Error(
			w,
			http.StatusBadRequest,
			"token required",
		)
		return
	}

	sub, err := h.service.PublicSubscriptionBase64(
		r.Context(),
		token,
	)

	if err != nil {
		web.WriteError(
			w,
			err,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"text/plain",
	)

	_, _ = w.Write([]byte(sub))
}

func (h *Handler) CreateDevice(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req CreateDeviceRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		web.Error(
			w,
			http.StatusBadRequest,
			"invalid request",
		)
		return
	}

	userID, ok := r.Context().Value(
		middleware.ContextUserID,
	).(string)

	if !ok || userID == "" {
		web.Error(
			w,
			http.StatusUnauthorized,
			"unauthorized",
		)
		return
	}

	device, err := h.service.CreateDevice(
		r.Context(),
		userID,
		req,
	)

	if err != nil {
		web.WriteError(w, err)
		return
	}

	web.OK(
		w,
		device,
	)
}

func (h *Handler) ListDevices(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID, ok := r.Context().Value(
		middleware.ContextUserID,
	).(string)

	if !ok || userID == "" {
		web.Error(
			w,
			http.StatusUnauthorized,
			"unauthorized",
		)
		return
	}

	devices, err := h.service.ListDevices(
		r.Context(),
		userID,
	)

	if err != nil {
		web.WriteError(w, err)
		return
	}

	web.OK(
		w,
		devices,
	)
}

func (h *Handler) DeleteDevice(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID, ok := r.Context().Value(
		middleware.ContextUserID,
	).(string)

	if !ok || userID == "" {
		web.Error(
			w,
			http.StatusUnauthorized,
			"unauthorized",
		)
		return
	}

	id := chi.URLParam(
		r,
		"id",
	)

	err := h.service.DeleteDevice(
		r.Context(),
		userID,
		id,
	)

	if err != nil {
		web.WriteError(w, err)
		return
	}

	web.OK(
		w,
		map[string]string{
			"status": "deleted",
		},
	)
}

func (h *Handler) RevokeDevice(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID, ok := r.Context().Value(
		middleware.ContextUserID,
	).(string)

	if !ok || userID == "" {
		web.Error(
			w,
			http.StatusUnauthorized,
			"unauthorized",
		)
		return
	}

	id := chi.URLParam(
		r,
		"id",
	)

	err := h.service.RevokeDevice(
		r.Context(),
		userID,
		id,
	)

	if err != nil {
		web.WriteError(w, err)
		return
	}

	web.OK(
		w,
		map[string]string{
			"status": "revoked",
		},
	)
}
