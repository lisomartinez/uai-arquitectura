package main

import (
	"fmt"
	"log"
	"net/http"
	"orders-service/internal/domain/orders"
	"orders-service/internal/domain/restaurants"
	"orders-service/internal/domain/users"
	ordersRepository "orders-service/internal/repository/orders"
	"orders-service/internal/rest"
	"strings"

	"github.com/go-chi/chi"
)

func main() {

	orderRepository := ordersRepository.NewCosmosDbRepository()
	usersService := users.NewUserService()
	restaurantService := restaurants.NewService()
	orderService := orders.NewService(orderRepository, restaurantService, usersService)

	handler := rest.NewOrdersHandler(orderService)

	router := chi.NewRouter()
	rest.RegisterOrderRoutes(router, handler)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	log.Print("Availability routes")
	for _, a := range router.Routes() {
		for _, b := range a.SubRoutes.Routes() {
			log.Print(fmt.Sprint(strings.ReplaceAll(a.Pattern, "/*", ""), b.Pattern))
		}
	}

	port := ":8080"
	log.Print(fmt.Sprint("Starting server at port", port))
	log.Fatal(http.ListenAndServe(port, router))
}
