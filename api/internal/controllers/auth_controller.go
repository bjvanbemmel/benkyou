package controllers

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"strings"
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

const (
	HASH_ALGORITHM_VERSION = 1
)

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

	user, err := a.userRepository.GetWithPasswordByEmail(req.Email)
	if err != nil {
		response.NewError(w, errors.ErrInvalidCredentials)
		return
	}

	splitUserPassword := strings.Split(user.Password, "$")
	salt := splitUserPassword[2]
	hashedPassword := utils.Hash(req.Password + salt)

	if splitUserPassword[1] != hashedPassword {
		response.NewError(w, errors.ErrInvalidCredentials)
		return
	}

	// Don't create a new token when a valid one already exists
	token, err := a.tokenRepository.GetByUserID(user.ID)
	if err == nil && token.ExpiresAt.After(time.Now()) {
		response.New(w, http.StatusOK, token)
		return
		// Remove token if expired
	} else if err == nil {
		a.tokenRepository.Delete(token.ID)
	}

	token, err = a.tokenRepository.Create(data.CreateTokenParams{
		UserID:    user.ID,
		Value:     utils.Hash(fmt.Sprintf("%v%d%s", user.ID, time.Now().UnixMicro(), utils.RandomString())),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
	})
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, token)
}

func (a AuthController) Register(w http.ResponseWriter, r *http.Request) {
	req, err := requests.Validate[requests.AuthRegisterRequest](r)
	if err != nil {
		response.NewError(w, err)
		return
	}

	if req.AccessToken != utils.AccessToken {
		response.NewError(w, errors.ErrIncorrectAccessToken)
		return
	}

	salt := utils.RandomString()
	saltedPassword := req.Password + salt

	user, err := a.userRepository.Create(data.CreateUserParams{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  fmt.Sprintf("%d$%x$%s", HASH_ALGORITHM_VERSION, sha256.Sum256([]byte(saltedPassword)), salt),
	})
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusCreated, user)
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
