package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/bjvanbemmel/benkyou/internal/data"
	"github.com/bjvanbemmel/benkyou/internal/repositories"
	"github.com/bjvanbemmel/benkyou/internal/requests"
	"github.com/bjvanbemmel/benkyou/internal/response"
	"github.com/bjvanbemmel/benkyou/internal/utils"
)

type AuthController struct {
	users  repositories.UserRepository
	tokens repositories.TokenRepository
}

func NewAuthController(users repositories.UserRepository, tokens repositories.TokenRepository) AuthController {
	return AuthController{
		users:  users,
		tokens: tokens,
	}
}

func (a AuthController) Login(w http.ResponseWriter, r *http.Request) {
	req, err := requests.Validate[requests.AuthLoginRequest](r)
	if err != nil {
		response.NewError(w, http.StatusBadRequest, err)
		return
	}

	user, err := a.users.GetUserWithPasswordByUsername(req.Username)
	if err != nil {
		response.NewError(w, http.StatusNotFound, errors.New("given user does not exist"))
		return
	}

	hashedPassword := utils.Hash(req.Password)

	if user.Password != hashedPassword {
		response.NewError(w, http.StatusUnauthorized, errors.New("incorrect password"))
		return
	}

	// Don't create a new token when a valid one already exists
	token, err := a.tokens.GetByUserID(user.ID)
	if err == nil && token.ExpiresAt.After(time.Now()) {
		response.New(w, http.StatusOK, token)
		return
	}

	token, err = a.tokens.Create(data.CreateTokenParams{
		UserID:    user.ID,
		Value:     utils.Hash(fmt.Sprintf("%v%d", user.ID, time.Now().UnixMicro())),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
	})
	if err != nil {
		response.NewError(w, http.StatusInternalServerError, err)
		return
	}

	response.New(w, http.StatusOK, token)
}

func (a AuthController) Identity(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("token").(data.Token)

	user, err := a.users.Get(token.UserID)
	if err != nil {
		response.NewError(w, http.StatusInternalServerError, err)
		return
	}

	response.New(w, http.StatusOK, user)
}

func (a AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("token").(data.Token)

	err := a.tokens.Delete(token.ID)
	if err != nil {
		response.NewError(w, http.StatusInternalServerError, err)
		return
	}

	response.New(w, http.StatusOK, "")
}
