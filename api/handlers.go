package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHandler(db *Application) http.Handler {
	r := chi.NewMux()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		SendJSON(w, Response{Data: "System is healthy and running"}, http.StatusOK)
	})

	return r
}
