package requests

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/bjvanbemmel/benkyou/internal/errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	en := en.New()
	uni = ut.New(en, en)

	trans, _ = uni.GetTranslator("en")
	validate = validator.New(validator.WithRequiredStructEnabled())
	en_translations.RegisterDefaultTranslations(validate, trans)

	// Use JSON property names instead of Go struct field names
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
}

type Request interface{}

func Decode[T Request](r *http.Request) (T, error) {
	var req T
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return req, errors.ErrInvalidRequest
	}

	return req, nil
}

// Assume that in case of an error, it will always be of type []FieldError.
func Validate[T Request](r *http.Request) (T, error) {
	var req T

	req, err := Decode[T](r)
	if err != nil {
		return req, err
	}

	if err := validate.Struct(req); err != nil {
		return req, errors.New(err.(validator.ValidationErrors)[0].Translate(trans))
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

// POST /auth/register
type AuthRegisterRequest struct {
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
