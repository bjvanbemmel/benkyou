package controllers

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/bjvanbemmel/benkyou/internal/data"
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
		http.Error(w, http.StatusText(500), 500)
		return
	}

	response.New(w, http.StatusOK, users)
}

func (u UserController) Get(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)

	user, err := u.repository.Get(id)
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

	user, err := u.repository.Create(data.CreateUserParams{
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

	req, err := requests.Validate[requests.UserUpdateRequest](r)
	if err != nil {
		response.NewError(w, http.StatusBadRequest, err)
		return
	}

	user, err := u.repository.Update(data.UpdateUserParams{
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

	if err := u.repository.Delete(id); err != nil {
		response.NewError(w, http.StatusInternalServerError, err)
	}

	response.New(w, http.StatusOK, "")
}
