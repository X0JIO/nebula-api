package web

import (
	"errors"
	"net/http"

	"github.com/X0JIO/nebula-api/internal/shared/apperrors"
)

func WriteError(
	w http.ResponseWriter,
	err error,
) {

	switch {

	case errors.Is(err, apperrors.ErrUnauthorized):

		http.Error(
			w,
			err.Error(),
			http.StatusUnauthorized,
		)

	case errors.Is(err, apperrors.ErrForbidden):

		http.Error(
			w,
			err.Error(),
			http.StatusForbidden,
		)

	case errors.Is(err, apperrors.ErrNotFound):

		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)

	case errors.Is(err, apperrors.ErrUserNotFound):

		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)

	case errors.Is(err, apperrors.ErrProjectNotFound):

		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)

	case errors.Is(err, apperrors.ErrTaskNotFound):

		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)

	case errors.Is(err, apperrors.ErrInvalidStatus):

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

	case errors.Is(err, apperrors.ErrInvalidPriority):

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

	case errors.Is(err, apperrors.ErrTitleRequired):

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

	default:

		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)

	}

}

func Error(
	w http.ResponseWriter,
	status int,
	message string,
) {

	http.Error(
		w,
		message,
		status,
	)

}
