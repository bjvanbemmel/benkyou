package controllers

import (
	"net/http"

	"github.com/bjvanbemmel/benkyou/internal/data"
	"github.com/bjvanbemmel/benkyou/internal/repositories"
	"github.com/bjvanbemmel/benkyou/internal/requests"
	"github.com/bjvanbemmel/benkyou/internal/response"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type FeatureController struct {
	featureRepository repositories.FeatureRepository
	userRepository    repositories.UserRepository
	sprintRepository  repositories.SprintRepository
}

func NewFeatureController(
	featureRepository repositories.FeatureRepository,
	userRepository repositories.UserRepository,
	sprintRepository repositories.SprintRepository,
) FeatureController {
	return FeatureController{
		featureRepository: featureRepository,
		userRepository:    userRepository,
		sprintRepository:  sprintRepository,
	}
}

func (f FeatureController) Index(w http.ResponseWriter, r *http.Request) {
	features, err := f.featureRepository.Index()
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, features)
}

func (f FeatureController) Get(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)

	feature, err := f.featureRepository.Get(id)
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

	if req.UserID != "" {
		if _, err := f.userRepository.Get(uuid.MustParse(req.UserID)); err != nil {
			response.NewError(w, err)
			return
		}
	}

	if req.SprintID != "" {
		if _, err := f.sprintRepository.Get(uuid.MustParse(req.SprintID)); err != nil {
			response.NewError(w, err)
			return
		}
	}

	var sprintId pgtype.UUID
	sprintId.Scan(req.SprintID)

	var userId pgtype.UUID
	userId.Scan(req.UserID)

	feature, err := f.featureRepository.Create(data.CreateFeatureParams{
		UserID:      userId,
		SprintID:    sprintId,
		State:       req.State,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusCreated, feature)
}

func (f FeatureController) Update(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)

	req, err := requests.Validate[requests.FeatureUpdateRequest](r)
	if err != nil {
		response.NewError(w, err)
		return
	}

	if req.UserID != "" {
		if _, err := f.userRepository.Get(uuid.MustParse(req.UserID)); err != nil {
			response.NewError(w, err)
			return
		}
	}

	if req.SprintID != "" {
		if _, err := f.sprintRepository.Get(uuid.MustParse(req.SprintID)); err != nil {
			response.NewError(w, err)
			return
		}
	}

	var userId pgtype.UUID
	userId.Scan(req.UserID)

	var sprintId pgtype.UUID
	sprintId.Scan(req.SprintID)

	feature, err := f.featureRepository.Get(id)
	if err != nil {
		response.NewError(w, err)
		return
	}

	if req.Position < feature.Position {
		if err := f.featureRepository.IncrementPositionForAllFeaturesBetweenDescending(feature.Position, req.Position); err != nil {
			response.NewError(w, err)
			return
		}
	} else if req.Position > feature.Position {
		if err := f.featureRepository.IncrementPositionForAllFeaturesBetweenAscending(feature.Position, req.Position); err != nil {
			response.NewError(w, err)
			return
		}
	}

	feature, err = f.featureRepository.Update(data.UpdateFeatureParams{
		ID:          id,
		UserID:      userId,
		SprintID:    sprintId,
		State:       req.State,
		Title:       req.Title,
		Description: req.Description,
		Position:    req.Position,
	})
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, feature)
}

func (f FeatureController) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)

	if err := f.featureRepository.Delete(id); err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, "")
}
