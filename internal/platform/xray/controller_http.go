package xray

import (
	"encoding/json"
	"net/http"
)

type HTTPController struct {
	service *Service
}

func NewHTTPHandler(
	service *Service,
) *HTTPController {

	return &HTTPController{
		service: service,
	}
}

func (h *HTTPController) Status(
	w http.ResponseWriter,
	r *http.Request,
) {

	json.NewEncoder(w).Encode(
		h.service.Status(),
	)
}

func (h *HTTPController) Start(
	w http.ResponseWriter,
	r *http.Request,
) {

	if err := h.service.Start(r.Context()); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *HTTPController) Stop(
	w http.ResponseWriter,
	r *http.Request,
) {

	if err := h.service.Stop(r.Context()); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *HTTPController) Restart(
	w http.ResponseWriter,
	r *http.Request,
) {

	if err := h.service.Restart(r.Context()); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *HTTPController) Reload(
	w http.ResponseWriter,
	r *http.Request,
) {

	if err := h.service.Reload(r.Context()); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *HTTPController) Version(
	w http.ResponseWriter,
	r *http.Request,
) {

	version, err := h.service.Version(
		r.Context(),
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"version": version,
		},
	)
}

func (h *HTTPController) Validate(
	w http.ResponseWriter,
	r *http.Request,
) {

	if err := h.service.Validate(r.Context()); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
}
