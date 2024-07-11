package middlewares

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/bjvanbemmel/benkyou/internal/cache"
	"github.com/bjvanbemmel/benkyou/internal/data"
	"github.com/bjvanbemmel/benkyou/internal/repositories"
	"github.com/bjvanbemmel/benkyou/internal/response"
	"github.com/bjvanbemmel/benkyou/internal/utils"
)

var (
	tokenCache cache.Cache[string, data.Token] = cache.New[string, data.Token]()
	tokenRepo  repositories.TokenRepository    = initTokenRepo()
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.NewError(w, http.StatusUnauthorized, errors.New("no bearer token"))
			return
		}

		bearerToken := strings.Split(authHeader, "Bearer ")[1]

		token, err := getToken(bearerToken)
		if err != nil {
			response.NewError(w, http.StatusUnauthorized, errors.New("incorrect token"))
			return
		}

		if err := expireToken(token); err != nil {
			response.NewError(w, http.StatusUnauthorized, errors.New("token has expired"))
			return
		}

		ctx := context.WithValue(r.Context(), "token", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Gets a token from either the cache or database
func getToken(bearerToken string) (data.Token, error) {
	hash := utils.Hash(bearerToken)

	token, err := tokenCache.Get(hash)
	if err == nil {
		return token, nil
	}

	token, err = tokenRepo.GetByValue(bearerToken)
	if err != nil {
		return data.Token{}, err
	}

	if err := tokenCache.Insert(hash, token); err != nil {
		return data.Token{}, err
	}

	return token, err
}

// Checks whether the given token has expired or not.
// If it did expire, deletes it from the database and cache.
// Throws an error when the token has expired.
// In case of a database error, it will ignore and log it.
func expireToken(token data.Token) error {
	if token.ExpiresAt.After(time.Now()) {
		return nil
	}
	defer tokenCache.Delete(utils.Hash(token.Value))

	if err := tokenRepo.Delete(token.ID); err != nil {
		slog.Error(err.Error())
	}

	return errors.New("token has expired")
}

// Method should only be called within the var() block and never after.
// Will throw a panic if it cannot create a database connection
func initTokenRepo() repositories.TokenRepository {
	repo, err := repositories.NewTokenRepository(context.Background())
	if err != nil {
		panic(err)
	}

	return repo
}
