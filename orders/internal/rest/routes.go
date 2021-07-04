package rest

import (
	"github.com/go-chi/chi"
)

func RegisterOrderRoutes(r chi.Router, handler *OrdersHandler) {
	r.Route("/shipping", func(r chi.Router) {
		r.Get("/{id}", handler.getOrder)
		r.Post("/", handler.createOrder)
	})
}