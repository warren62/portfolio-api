package router

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"

	"portfolio-api/api/resource/health"
	"portfolio-api/api/resource/post"
)

func New(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", health.Read)

	r.Route("/v1", func(r chi.Router) {
		postAPI := post.New(db)
		r.Get("/posts", postAPI.List)
		r.Post("/posts", postAPI.Create)
		r.Get("/posts/{id}", postAPI.Read)
		r.Put("/posts/{id}", postAPI.Update)
		r.Delete("/posts/{id}", postAPI.Delete)
	})

	return r
}
