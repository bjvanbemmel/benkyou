package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

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

	// Validate whether a given field is using the RFC3339 datetime format
	validate.RegisterValidation("rfc3339", IsValidDateTimeFormat)
}

type Request interface{}

func Decode[T Request](r *http.Request) (T, error) {
	var req T
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
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

func IsValidDateTimeFormat(fl validator.FieldLevel) bool {
	_, err := time.Parse(time.RFC3339, fl.Field().String())
	return err == nil
}

// PUT /users/{id}
type UserUpdateRequest struct {
	Email           string `json:"email" validate:"required,email,max=255"`
	FirstName       string `json:"first_name" validate:"required,min=1,max=50"`
	LastName        string `json:"last_name" validate:"required,min=1,max=100"`
	Password        string `json:"password" validate:"required,min=8,max=255"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

// POST /auth/register
type AuthRegisterRequest struct {
	Email           string `json:"email" validate:"required,email,max=255"`
	FirstName       string `json:"first_name" validate:"required,min=1,max=50"`
	LastName        string `json:"last_name" validate:"required,min=1,max=100"`
	Password        string `json:"password" validate:"required,min=8,max=255"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	AccessToken     string `json:"access_token" validate:"required,min=1,max=255"`
}

// POST /auth/login
type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,max=255"`
}

// POST /sprints
type SprintCreateRequest struct {
	Title     string `json:"title" validate:"required,max=80"`
	StartDate string `json:"start_date" validate:"required,rfc3339"`
	EndDate   string `json:"end_date" validate:"required,rfc3339"`
}

// PUT /sprints/{id}
type SprintUpdateRequest struct {
	Title     string `json:"title" validate:"required,max=80"`
	StartDate string `json:"start_date" validate:"required,rfc3339"`
	EndDate   string `json:"end_date" validate:"required,rfc3339"`
}

// POST /features
type FeatureCreateRequest struct {
	UserID      string `json:"user_id" validate:"omitempty,uuid"`
	SprintID    string `json:"sprint_id" validate:"omitempty,uuid"`
	State       int32  `json:"state" validate:"number"`
	Title       string `json:"title" validate:"required,max=255"`
	Description string `json:"description"`
}

// PUT /features/{id}
type FeatureUpdateRequest struct {
	UserID      string `json:"user_id" validate:"omitempty,uuid"`
	SprintID    string `json:"sprint_id" validate:"omitempty,uuid"`
	State       int32  `json:"state" validate:"number"`
	Title       string `json:"title" validate:"required,max=255"`
	Description string `json:"description"`
	Position    int32  `json:"position" validate:"number"`
}

// POST /requirements
type RequirementCreateRequest struct {
	UserID      string `json:"user_id" validate:"omitempty,uuid"`
	SprintID    string `json:"sprint_id" validate:"omitempty,uuid"`
	FeatureID   string `json:"feature_id" validate:"omitempty,uuid"`
	State       int32  `json:"state" validate:"number"`
	Title       string `json:"title" validate:"required,max=255"`
	Estimate    int64  `json:"estimate" validate:"omitempty,number,min=1,max=13"`
	Description string `json:"description"`
}

// PUT /requirements/{id}
type RequirementUpdateRequest struct {
	UserID      string `json:"user_id" validate:"omitempty,uuid"`
	SprintID    string `json:"sprint_id" validate:"omitempty,uuid"`
	FeatureID   string `json:"feature_id" validate:"omitempty,uuid"`
	State       int32  `json:"state" validate:"number"`
	Title       string `json:"title" validate:"required,max=255"`
	Estimate    int64  `json:"estimate" validate:"omitempty,number,min=1,max=13"`
	Description string `json:"description"`
}

// POST /action-points
type ActionPointCreateRequest struct {
	Title string `json:"title" validate:"required"`
}

// PUT /action-points/{id}
type ActionPointUpdateRequest struct {
	Title string `json:"title" validate:"required"`
}
