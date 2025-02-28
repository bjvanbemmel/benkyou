package controllers

import (
	"net/http"
	"strconv"

	"github.com/bjvanbemmel/benkyou/internal/data"
	"github.com/bjvanbemmel/benkyou/internal/repositories"
	"github.com/bjvanbemmel/benkyou/internal/requests"
	"github.com/bjvanbemmel/benkyou/internal/response"
	"github.com/google/uuid"
)

type FeatureController struct {
	repository repositories.FeatureRepository
}

func NewFeatureController(repo repositories.FeatureRepository) FeatureController {
	return FeatureController{
		repository: repo,
	}
}

func (f FeatureController) Index(w http.ResponseWriter, r *http.Request) {
	features, err := f.repository.Index()
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, features)
}

func (f FeatureController) Get(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)

	feature, err := f.repository.Get(id)
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, feature)
}

func (f FeatureController) Create(w http.ResponseWriter, r *http.Request) {
	req, err := requests.Validate[requests.FeatureCreateRequest](r)
	if err != nil {
		response.NewError(w, err)
		return
	}

	state, _ := strconv.Atoi(req.State)
	feature, err := f.repository.Create(data.CreateFeatureParams{
		UserID:      uuid.MustParse(req.UserID),
		SprintID:    uuid.MustParse(req.SprintID),
		State:       int32(state),
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, feature)
}

func (f FeatureController) Update(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)

	req, err := requests.Validate[requests.FeatureUpdateRequest](r)
	if err != nil {
		response.NewError(w, err)
		return
	}

	state, _ := strconv.Atoi(req.State)
	feature, err := f.repository.Update(data.UpdateFeatureParams{
		ID:          id,
		UserID:      uuid.MustParse(req.UserID),
		SprintID:    uuid.MustParse(req.SprintID),
		State:       int32(state),
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, feature)
}

func (f FeatureController) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)

	if err := f.repository.Delete(id); err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, "")
}
