package comments

import "github.com/go-chi/chi/v5"

func Routes(
	r chi.Router,
	handler *Handler,
) {

	r.Route("/tasks/{taskId}/comments", func(r chi.Router) {

		r.Post("/", handler.CreateComment)
		r.Get("/", handler.ListTaskComments)

	})

	r.Delete(
		"/comments/{id}",
		handler.DeleteComment,
	)
}
