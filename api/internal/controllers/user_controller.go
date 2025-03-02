package controllers

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/bjvanbemmel/benkyou/internal/data"
	"github.com/bjvanbemmel/benkyou/internal/errors"
	"github.com/bjvanbemmel/benkyou/internal/repositories"
	"github.com/bjvanbemmel/benkyou/internal/requests"
	"github.com/bjvanbemmel/benkyou/internal/response"
	"github.com/google/uuid"
)

type UserController struct {
	repository repositories.UserRepository
}

func NewUserController(repo repositories.UserRepository) UserController {
	return UserController{
		repository: repo,
	}
}

func (u UserController) Index(w http.ResponseWriter, r *http.Request) {
	users, err := u.repository.Index()
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, users)
}

func (u UserController) Get(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)

	user, err := u.repository.Get(id)
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, user)
}

func (u UserController) Update(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)
	token := r.Context().Value("token").(data.Token)

	if token.UserID != id {
		response.NewError(w, errors.ErrUnauthorized)
		return
	}

	req, err := requests.Validate[requests.UserUpdateRequest](r)
	if err != nil {
		response.NewError(w, err)
		return
	}

	user, err := u.repository.Update(data.UpdateUserParams{
		ID:        id,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password))),
	})
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, user)
}

func (u UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)
	token := r.Context().Value("token").(data.Token)

	if token.UserID != id {
		response.NewError(w, errors.ErrUnauthorized)
		return
	}

	if err := u.repository.Delete(id); err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, "")
}
