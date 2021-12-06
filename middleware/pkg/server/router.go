package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func setupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	return r
}
