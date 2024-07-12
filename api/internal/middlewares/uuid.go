package middlewares

import (
	"context"
	"net/http"

	"github.com/bjvanbemmel/benkyou/internal/errors"
	"github.com/bjvanbemmel/benkyou/internal/response"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// Parses the UUID from a {id} query param and adds it to the context.
func Uuid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			response.NewError(w, errors.ErrInvalidID)
			return
		}

		parsed, err := uuid.Parse(id)
		if err != nil {
			response.NewError(w, errors.ErrInvalidID)
			return
		}

		ctx := context.WithValue(r.Context(), "uuid", parsed)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
