package main

import (
	"context"
	"net/http"
	"time"

	"github.com/bjvanbemmel/benkyou/internal/controllers"
	"github.com/bjvanbemmel/benkyou/internal/middlewares"
	"github.com/bjvanbemmel/benkyou/internal/repositories"
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

	userRepository, err := repositories.NewUserRepository(context.Background())
	if err != nil {
		panic(err)
	}

	userController := controllers.NewUserController(userRepository)

	tokenRepository, err := repositories.NewTokenRepository(context.Background())
	if err != nil {
		panic(err)
	}

	authController := controllers.NewAuthController(userRepository, tokenRepository)

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

		r.Post("/auth/logout", authController.Logout)
		r.Get("/auth/identity", authController.Identity)
	})

	r.Post("/auth/login", authController.Login)
	r.Post("/auth/register", authController.Register)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":80", r)
}
