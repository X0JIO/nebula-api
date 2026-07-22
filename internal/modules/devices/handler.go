package devices

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/X0JIO/nebula-api/internal/platform/web/middleware"
	"github.com/go-chi/chi/v5"
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

// List godoc
//
//	@Summary		List devices
//	@Description	Returns current user's devices
//	@Tags			Devices
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{array}	DeviceResponse
//	@Router			/devices [get]
func (h *Handler) List(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID := middleware.UserID(r.Context())
	log.Printf("devices handler userID = %q", userID)
	devices, err := h.service.List(r.Context(), userID)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		ToResponses(devices),
	)
}

// Delete godoc
//
//	@Summary		Delete device
//	@Description	Deletes a specific device
//	@Tags			Devices
//	@Security		BearerAuth
//	@Success		204
//	@Router			/devices/{id} [delete]
func (h *Handler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {

	id := chi.URLParam(r, "id")

	if err := h.service.Delete(r.Context(), id); err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteAll godoc
//
//	@Summary		Delete all devices
//	@Description	Deletes all devices for the current user
//	@Tags			Devices
//	@Security		BearerAuth
//	@Success		204
//	@Router			/devices [delete]
func (h *Handler) DeleteAll(
	w http.ResponseWriter,
	r *http.Request,
) {

	userID := middleware.UserID(r.Context())

	if err := h.service.DeleteAll(
		r.Context(),
		userID,
	); err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
