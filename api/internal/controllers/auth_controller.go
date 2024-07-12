package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bjvanbemmel/benkyou/internal/data"
	"github.com/bjvanbemmel/benkyou/internal/errors"
	"github.com/bjvanbemmel/benkyou/internal/repositories"
	"github.com/bjvanbemmel/benkyou/internal/requests"
	"github.com/bjvanbemmel/benkyou/internal/response"
	"github.com/bjvanbemmel/benkyou/internal/utils"
)

type AuthController struct {
	userRepository  repositories.UserRepository
	tokenRepository repositories.TokenRepository
}

func NewAuthController(userRepo repositories.UserRepository, tokenRepo repositories.TokenRepository) AuthController {
	return AuthController{
		userRepository:  userRepo,
		tokenRepository: tokenRepo,
	}
}

func (a AuthController) Login(w http.ResponseWriter, r *http.Request) {
	req, err := requests.Validate[requests.AuthLoginRequest](r)
	if err != nil {
		response.NewError(w, err)
		return
	}

	user, err := a.userRepository.GetUserWithPasswordByUsername(req.Username)
	if err != nil {
		response.NewError(w, errors.ErrInvalidCredentials)
		return
	}

	hashedPassword := utils.Hash(req.Password)

	if user.Password != hashedPassword {
		response.NewError(w, errors.ErrInvalidCredentials)
		return
	}

	// Don't create a new token when a valid one already exists
	token, err := a.tokenRepository.GetByUserID(user.ID)
	if err == nil && token.ExpiresAt.After(time.Now()) {
		response.New(w, http.StatusOK, token)
		return
	}

	token, err = a.tokenRepository.Create(data.CreateTokenParams{
		UserID:    user.ID,
		Value:     utils.Hash(fmt.Sprintf("%v%d", user.ID, time.Now().UnixMicro())),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
	})
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, token)
}

func (a AuthController) Identity(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("token").(data.Token)

	user, err := a.userRepository.Get(token.UserID)
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, user)
}

func (a AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("token").(data.Token)

	err := a.tokenRepository.Delete(token.ID)
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, "")
}
