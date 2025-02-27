package application

import (
	"net/http"

	"github.com/abdullahalsazib/orders_api/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRouters() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	router.Route("/orders", loadOrderRouters)

	return router
}

func loadOrderRouters(router chi.Router) {
    orderHandler := &handler.Order{}

    router.Post("/", orderHandler.Create)
    router.Get("/", orderHandler.List)
    router.Get("/{id}", orderHandler.GetByID)
    router.Put("/{id}", orderHandler.UpdateByID)
    router.Delete("/{id}", orderHandler.DeleteByID)
}
