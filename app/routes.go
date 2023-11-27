package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/alex-appy-love-story/backend/handler"
)

func loadRoutes(a *App) *chi.Mux {
	router := chi.NewRouter()

	//taskHanlder := TaskHandler{app: a}
	router.Use(middleware.Logger)
	//router.Use(taskHanlder.routerContextMiddleware)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", loadOrderRoutes)

	return router
}

func loadOrderRoutes(router chi.Router) {
	router.Post("/", handler.Create)
}
