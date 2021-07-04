package main

import (
	"fmt"
	"log"
	"net/http"
	"orders-service/internal/domain/orders"
	"orders-service/internal/domain/restaurants"
	"orders-service/internal/domain/users"
	"orders-service/internal/repository"
	"orders-service/internal/rest"
	"strings"

	"github.com/go-chi/chi"
)

func main() {

	orderRepository := repository.NewLocalOrderRepository()
	usersService := users.NewUserService()
	restaurantService := restaurants.NewService()
	orderService := orders.NewService(orderRepository, restaurantService, usersService)

	handler := rest.NewOrdersHandler(orderService)

	//Start App
	router := chi.NewRouter()
	rest.RegisterOrderRoutes(router, handler)

	log.Print("Availability routes")
	for _, a := range router.Routes() {
		for _, b := range a.SubRoutes.Routes() {
			log.Print(fmt.Sprint(strings.ReplaceAll(a.Pattern, "/*", ""), b.Pattern))
		}

	}
	port := ":8020"
	log.Print(fmt.Sprint("Starting server at port", port))
	log.Fatal(http.ListenAndServe(port, router))
}
