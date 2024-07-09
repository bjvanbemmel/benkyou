package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/bjvanbemmel/benkyou/internal/repositories"
	"github.com/bjvanbemmel/benkyou/internal/response"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.NewError(w, http.StatusUnauthorized, errors.New("no bearer token"))
			return
		}

		bearerToken := strings.Split(authHeader, "Bearer ")[1]

		tokenRepo, err := repositories.NewTokenRepository(context.Background())
		if err != nil {
			response.NewError(w, http.StatusInternalServerError, err)
			return
		}

		token, err := tokenRepo.GetByValue(bearerToken)
		if err != nil {
			response.NewError(w, http.StatusUnauthorized, errors.New("incorrect token"))
			return
		}

		if !token.ExpiresAt.After(time.Now()) {
			// Delete expired token, expose error upon failure
			if err := tokenRepo.Delete(token.ID); err != nil {
				response.NewError(w, http.StatusInternalServerError, err)
				return
			}
			response.NewError(w, http.StatusUnauthorized, errors.New("token has expired"))
			return
		}

		ctx := context.WithValue(r.Context(), "token", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
