package apperror

import (
	"errors"
	"net/http"
)

type appHandler = func(w http.ResponseWriter, r *http.Request) error

func Middleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var appErr *AppError
		err := h(w, r)
		if err != nil {
			if errors.As(err, &appErr) {

				if errors.Is(err, ErrNotFound) {
					w.WriteHeader(http.StatusNotFound)
					_ = ErrNotFound.JSON(w)
					return
				}

				err := err.(*AppError)
				w.WriteHeader(http.StatusBadRequest)
				_ = err.JSON(w)
				return
			}
			w.WriteHeader(418)
			_ = SystemError(err.Error()).JSON(w)
		}
	}
}
