package rest

import (
	"net/http"

	"github.com/go-chi/chi"
)

func RegisterOrderRoutes(r chi.Router, handler *OrdersHandler) {
	r.Route("/orders", func(r chi.Router) {
		r.Get("/{id}", handler.getOrder)
		r.Post("/", handler.createOrder)
	})
}

func RegisterHealthRoutes(r chi.Router) {
	r.Route("/health", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
		})
	})
}
