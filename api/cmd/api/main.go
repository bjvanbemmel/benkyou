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
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(time.Second * 60))

	userRepository, err := repositories.NewUserRepository(context.Background())
	if err != nil {
		panic(err)
	}

	userController := controllers.NewUserController(userRepository)

	r.Get("/users", userController.Index)
	r.Post("/users", userController.Create)

	r.Route("/users/{id}", func(r chi.Router) {
		r.Use(middlewares.Uuid)
		r.Get("/", userController.Get)
		r.Put("/", userController.Update)
		r.Delete("/", userController.Delete)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":80", r)
}
