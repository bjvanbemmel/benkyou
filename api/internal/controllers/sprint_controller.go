package controllers

import (
	"net/http"
	"time"

	"github.com/bjvanbemmel/benkyou/internal/data"
	"github.com/bjvanbemmel/benkyou/internal/repositories"
	"github.com/bjvanbemmel/benkyou/internal/requests"
	"github.com/bjvanbemmel/benkyou/internal/response"
	"github.com/google/uuid"
)

type SprintController struct {
	repository repositories.SprintRepository
}

func NewSprintController(repo repositories.SprintRepository) SprintController {
	return SprintController{
		repository: repo,
	}
}

func (s SprintController) Index(w http.ResponseWriter, r *http.Request) {
	sprints, err := s.repository.Index()
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, sprints)
}

func (s SprintController) Get(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)

	sprint, err := s.repository.Get(id)
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, sprint)
}

func (s SprintController) Create(w http.ResponseWriter, r *http.Request) {
	req, err := requests.Validate[requests.SprintCreateRequest](r)
	if err != nil {
		response.NewError(w, err)
		return
	}

	startDate, _ := time.Parse(time.RFC3339, req.StartDate)
	endDate, _ := time.Parse(time.RFC3339, req.EndDate)
	sprint, err := s.repository.Create(data.CreateSprintParams{
		Title:     req.Title,
		StartDate: startDate,
		EndDate:   endDate,
	})
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusCreated, sprint)
}

func (s SprintController) Update(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)

	req, err := requests.Validate[requests.SprintUpdateRequest](r)
	if err != nil {
		response.NewError(w, err)
		return
	}

	startDate, _ := time.Parse(time.RFC3339, req.StartDate)
	endDate, _ := time.Parse(time.RFC3339, req.EndDate)
	sprint, err := s.repository.Update(data.UpdateSprintParams{
		ID:        id,
		Title:     req.Title,
		StartDate: startDate,
		EndDate:   endDate,
	})
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, sprint)
}

func (s SprintController) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)

	if err := s.repository.Delete(id); err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, "")
}
