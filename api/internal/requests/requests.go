package requests

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Request interface{}

func Decode[T Request](r *http.Request) (T, error) {
	var req T
	err := json.NewDecoder(r.Body).Decode(&req)

	return req, err
}

// Assume that in case of an error, it will always be of type []FieldError.
func Validate[T Request](r *http.Request) (T, error) {
	var req T

	req, err := Decode[T](r)
	if err != nil {
		return req, err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(req); err != nil {
		return req, err.(validator.ValidationErrors)[0]
	}

	return req, nil
}

// PUT /users/{id}
type UserUpdateRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Username        string `json:"username" validate:"required,min=3,max=30"`
	Password        string `json:"password" validate:"required,min=8,max=255"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

// POST /users/{id}
type UserCreateRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Username        string `json:"username" validate:"required,min=3,max=30"`
	Password        string `json:"password" validate:"required,min=8,max=255"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

// POST /auth/login
type AuthLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
