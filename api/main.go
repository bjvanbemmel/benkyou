package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
  fmt.Println("Hello, World!")

  r := chi.NewRouter()
  r.Use(middleware.Logger)
  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
  })

  http.ListenAndServe(":80", r)
}
