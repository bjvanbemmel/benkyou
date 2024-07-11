package controllers

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"

	"github.com/bjvanbemmel/benkyou/internal/data"
	"github.com/bjvanbemmel/benkyou/internal/repositories"
	"github.com/bjvanbemmel/benkyou/internal/requests"
	"github.com/bjvanbemmel/benkyou/internal/response"
	"github.com/google/uuid"
)

type UserController struct {
	users repositories.UserRepository
}

func NewUserController(repo repositories.UserRepository) UserController {
	return UserController{
		users: repo,
	}
}

func (u UserController) Index(w http.ResponseWriter, r *http.Request) {
	users, err := u.users.Index()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	response.New(w, http.StatusOK, users)
}

func (u UserController) Get(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)

	user, err := u.users.Get(id)
	if err != nil {
		response.NewError(w, http.StatusNotFound, err)
		return
	}

	response.New(w, http.StatusOK, user)
}

func (u UserController) Create(w http.ResponseWriter, r *http.Request) {
	req, err := requests.Validate[requests.UserCreateRequest](r)
	if err != nil {
		response.NewError(w, http.StatusBadRequest, err)
		return
	}

	user, err := u.users.Create(data.CreateUserParams{
		Email:    req.Email,
		Username: req.Username,
		Password: fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password))),
	})
	if err != nil {
		response.NewError(w, http.StatusInternalServerError, err)
		return
	}

	response.New(w, http.StatusCreated, user)
}

func (u UserController) Update(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)
	token := r.Context().Value("token").(data.Token)

	if token.UserID != id {
		response.NewError(w, http.StatusUnauthorized, errors.New("you are not authorized to edit this user"))
		return
	}

	req, err := requests.Validate[requests.UserUpdateRequest](r)
	if err != nil {
		response.NewError(w, http.StatusBadRequest, err)
		return
	}

	user, err := u.users.Update(data.UpdateUserParams{
		ID:       id,
		Email:    req.Email,
		Username: req.Username,
		Password: fmt.Sprintf("%x", sha256.Sum256([]byte(req.Password))),
	})
	if err != nil {
		response.NewError(w, http.StatusInternalServerError, err)
		return
	}

	response.New(w, http.StatusOK, user)
}

func (u UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)
	token := r.Context().Value("token").(data.Token)

	if token.UserID != id {
		response.NewError(w, http.StatusUnauthorized, errors.New("you are not authorized to delete this user"))
		return
	}

	if err := u.users.Delete(id); err != nil {
		response.NewError(w, http.StatusInternalServerError, err)
	}

	response.New(w, http.StatusOK, "")
}
