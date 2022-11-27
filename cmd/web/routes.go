package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)


func (app *Config) routes() http.Handler{
	mux := chi.NewRouter()

	// middleware
	mux.Use(middleware.Recover)

	mux.Get("/", app.HomePage)

	return mux
}