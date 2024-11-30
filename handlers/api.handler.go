package handlers

import (
	"gopoc/api/task"

	chiMid "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func ApiHandler(chiMux *chi.Mux) {
	chiMux.Use(chiMid.StripSlashes)
	chiMux.Route("/tasks", func(r chi.Router) {

		r.Get("/", task.GetAlltask)
		r.Post("/", task.CreateTask)
		r.Get("/{taskId}", task.GetTask)
		r.Put("/{taskId}", task.UpdateTask)
		r.Delete("/{taskId}", task.DeleteTask)
	})

	chiMux.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"), //The url pointing to API definition
	))
}
