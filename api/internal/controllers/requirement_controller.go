package controllers

import (
	"net/http"
	"strconv"

	"github.com/bjvanbemmel/benkyou/internal/data"
	"github.com/bjvanbemmel/benkyou/internal/repositories"
	"github.com/bjvanbemmel/benkyou/internal/requests"
	"github.com/bjvanbemmel/benkyou/internal/response"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type RequirementController struct {
	requirementRepository repositories.RequirementRepository
	featureRepository     repositories.FeatureRepository
	sprintRepository      repositories.SprintRepository
	userRepository        repositories.UserRepository
}

func NewRequirementController(
	requirementRepository repositories.RequirementRepository,
	featureRepository repositories.FeatureRepository,
	sprintRepository repositories.SprintRepository,
	userRepository repositories.UserRepository,
) RequirementController {
	return RequirementController{
		requirementRepository: requirementRepository,
		featureRepository:     featureRepository,
		sprintRepository:      sprintRepository,
		userRepository:        userRepository,
	}
}

func (c RequirementController) Index(w http.ResponseWriter, r *http.Request) {
	requirements, err := c.requirementRepository.Index()
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, requirements)
}

func (c RequirementController) Get(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)

	requirement, err := c.requirementRepository.Get(id)
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, requirement)
}

func (c RequirementController) Create(w http.ResponseWriter, r *http.Request) {
	req, err := requests.Validate[requests.RequirementCreateRequest](r)
	if err != nil {
		response.NewError(w, err)
		return
	}

	if req.UserID != "" {
		if _, err := c.userRepository.Get(uuid.MustParse(req.UserID)); err != nil {
			response.NewError(w, err)
			return
		}
	}

	if req.FeatureID != "" {
		if _, err := c.featureRepository.Get(uuid.MustParse(req.FeatureID)); err != nil {
			response.NewError(w, err)
			return
		}
	}

	if req.SprintID != "" {
		if _, err := c.sprintRepository.Get(uuid.MustParse(req.SprintID)); err != nil {
			response.NewError(w, err)
			return
		}
	}

	var userId pgtype.UUID
	userId.Scan(req.UserID)

	var sprintId pgtype.UUID
	sprintId.Scan(req.SprintID)

	var featureId pgtype.UUID
	featureId.Scan(req.FeatureID)

	state, _ := strconv.Atoi(req.State)
	requirement, err := c.requirementRepository.Create(data.CreateRequirementParams{
		UserID:      userId,
		SprintID:    sprintId,
		FeatureID:   featureId,
		State:       int32(state),
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusCreated, requirement)
}

func (c RequirementController) Update(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)

	req, err := requests.Validate[requests.RequirementUpdateRequest](r)
	if err != nil {
		response.NewError(w, err)
		return
	}

	if req.UserID != "" {
		if _, err := c.userRepository.Get(uuid.MustParse(req.UserID)); err != nil {
			response.NewError(w, err)
			return
		}
	}

	if req.FeatureID != "" {
		if _, err := c.featureRepository.Get(uuid.MustParse(req.FeatureID)); err != nil {
			response.NewError(w, err)
			return
		}
	}

	if req.SprintID != "" {
		if _, err := c.sprintRepository.Get(uuid.MustParse(req.SprintID)); err != nil {
			response.NewError(w, err)
			return
		}
	}

	var userId pgtype.UUID
	userId.Scan(req.UserID)

	var sprintId pgtype.UUID
	sprintId.Scan(req.SprintID)

	var featureId pgtype.UUID
	featureId.Scan(req.FeatureID)

	state, _ := strconv.Atoi(req.State)
	requirement, err := c.requirementRepository.Update(data.UpdateRequirementParams{
		ID:          id,
		UserID:      userId,
		SprintID:    sprintId,
		FeatureID:   featureId,
		State:       int32(state),
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, requirement)
}

func (c RequirementController) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)

	if err := c.requirementRepository.Delete(id); err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, "")
}
