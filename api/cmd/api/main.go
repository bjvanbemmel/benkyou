package main

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/bjvanbemmel/benkyou/internal/controllers"
	"github.com/bjvanbemmel/benkyou/internal/middlewares"
	"github.com/bjvanbemmel/benkyou/internal/repositories"
	"github.com/bjvanbemmel/benkyou/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Second * 60))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
		MaxAge:         300,
	}))

	// Add repositories
	userRepository, err := repositories.NewUserRepository(context.Background())
	if err != nil {
		panic(err)
	}
	defer userRepository.Close()

	tokenRepository, err := repositories.NewTokenRepository(context.Background())
	if err != nil {
		panic(err)
	}
	defer tokenRepository.Close()

	sprintRepository, err := repositories.NewSprintRepository(context.Background())
	if err != nil {
		panic(err)
	}
	defer sprintRepository.Close()

	featureRepository, err := repositories.NewFeatureRepository(context.Background())
	if err != nil {
		panic(err)
	}
	defer featureRepository.Close()

	requirementRepository, err := repositories.NewRequirementRepository(context.Background())
	if err != nil {
		panic(err)
	}
	defer requirementRepository.Close()

	actionPointRepository, err := repositories.NewActionPointRepository(context.Background())
	if err != nil {
		panic(err)
	}
	defer actionPointRepository.Close()

	// Add controllers
	userController := controllers.NewUserController(userRepository)
	authController := controllers.NewAuthController(userRepository, tokenRepository)
	sprintController := controllers.NewSprintController(sprintRepository)
	featureController := controllers.NewFeatureController(featureRepository, userRepository, sprintRepository)
	requirementController := controllers.NewRequirementController(requirementRepository, featureRepository, sprintRepository, userRepository)
	actionPointController := controllers.NewActionPointController(actionPointRepository)

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(middlewares.Auth)
		r.Get("/users", userController.Index)
		r.Route("/users/{id}", func(r chi.Router) {
			r.Use(middlewares.Uuid)
			r.Get("/", userController.Get)
			r.Put("/", userController.Update)
			r.Delete("/", userController.Delete)
		})

		r.Get("/sprints", sprintController.Index)
		r.Post("/sprints", sprintController.Create)
		r.Route("/sprints/{id}", func(r chi.Router) {
			r.Use(middlewares.Uuid)
			r.Get("/", sprintController.Get)
			r.Put("/", sprintController.Update)
			r.Delete("/", sprintController.Delete)
		})

		r.Get("/features", featureController.Index)
		r.Post("/features", featureController.Create)
		r.Route("/features/{id}", func(r chi.Router) {
			r.Use(middlewares.Uuid)
			r.Get("/", featureController.Get)
			r.Put("/", featureController.Update)
			r.Delete("/", featureController.Delete)
		})

		r.Get("/requirements", requirementController.Index)
		r.Post("/requirements", requirementController.Create)
		r.Route("/requirements/{id}", func(r chi.Router) {
			r.Use(middlewares.Uuid)
			r.Get("/", requirementController.Get)
			r.Put("/", requirementController.Update)
			r.Delete("/", requirementController.Delete)
		})

		r.Get("/action-points", actionPointController.Index)
		r.Post("/action-points", actionPointController.Create)
		r.Route("/action-points/{id}", func(r chi.Router) {
			r.Use(middlewares.Uuid)
			r.Get("/", actionPointController.Get)
			r.Put("/", actionPointController.Update)
			r.Delete("/", actionPointController.Delete)
		})

		r.Post("/auth/logout", authController.Logout)
		r.Get("/auth/identity", authController.Identity)
	})

	// Unprotected routes
	r.Post("/auth/login", authController.Login)
	r.Post("/auth/register", authController.Register)

	// Generate random access token for session
	utils.NewAccessToken()
	slog.Info("Access token has been generated for the current session.", "token", utils.AccessToken)

	http.ListenAndServe(":80", r)
}
