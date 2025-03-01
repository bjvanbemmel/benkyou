package controllers

import (
	"net/http"

	"github.com/bjvanbemmel/benkyou/internal/data"
	"github.com/bjvanbemmel/benkyou/internal/repositories"
	"github.com/bjvanbemmel/benkyou/internal/requests"
	"github.com/bjvanbemmel/benkyou/internal/response"
	"github.com/google/uuid"
)

type ActionPointController struct {
	repo repositories.ActionPointRepository
}

func NewActionPointController(repo repositories.ActionPointRepository) ActionPointController {
	return ActionPointController{
		repo: repo,
	}
}

func (a ActionPointController) Index(w http.ResponseWriter, r *http.Request) {
	actionPoints, err := a.repo.Index()
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, actionPoints)
}

func (a ActionPointController) Get(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)

	actionPoint, err := a.repo.Get(id)
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusOK, actionPoint)
}

func (a ActionPointController) Create(w http.ResponseWriter, r *http.Request) {
	req, err := requests.Validate[requests.ActionPointCreateRequest](r)
	if err != nil {
		response.NewError(w, err)
		return
	}

	actionPoint, err := a.repo.Create(req.Title)
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusCreated, actionPoint)
}

func (a ActionPointController) Update(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)

	req, err := requests.Validate[requests.ActionPointCreateRequest](r)
	if err != nil {
		response.NewError(w, err)
		return
	}

	actionPoint, err := a.repo.Update(data.UpdateActionPointParams{
		ID:    id,
		Title: req.Title,
	})
	if err != nil {
		response.NewError(w, err)
		return
	}

	response.New(w, http.StatusCreated, actionPoint)
}

func (a ActionPointController) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("uuid").(uuid.UUID)
	if err := a.repo.Delete(id); err != nil {
		response.NewError(w, err)
		return
	}
	response.New(w, http.StatusOK, "")
}
