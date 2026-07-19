package web

import (
	"net/http"
)

func OK(
	w http.ResponseWriter,
	data any,
) {
	WriteJSON(
		w,
		http.StatusOK,
		data,
	)
}

func Created(
	w http.ResponseWriter,
	data any,
) {
	WriteJSON(
		w,
		http.StatusCreated,
		data,
	)
}

func NoContent(
	w http.ResponseWriter,
) {
	w.WriteHeader(http.StatusNoContent)
}
