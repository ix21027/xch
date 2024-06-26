package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
    
	r.Route("/api", func(r chi.Router) {
		r.Get("/rate", getRate)
		r.Post("/subscribe", addSubscriber)
	})
	
	return r
}
