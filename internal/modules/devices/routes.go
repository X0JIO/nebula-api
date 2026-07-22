package devices

import "github.com/go-chi/chi/v5"

func RegisterRoutes(
	r chi.Router,
	h *Handler,
) {

	r.Route("/devices", func(r chi.Router) {

		r.Get("/", h.List)

		r.Delete("/", h.DeleteAll)

		r.Delete("/{id}", h.Delete)

	})
}
