package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Router - Endpoints in Http Server
func (s *Server) Router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/tickerwatch", func(r chi.Router) {
		r.Get("/", s.Get)       // GET /todos/{id} - read a single todo by :id
		r.Put("/", s.Update)    // PUT /todos/{id} - update a single todo by :id
		r.Delete("/", s.Delete) // DELETE /todos/{id} - delete a single todo by :id
	})
	return r
}
