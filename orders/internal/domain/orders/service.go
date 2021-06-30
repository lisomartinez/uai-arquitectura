package orders

import (
	"context"
	"log"
	"orders-service/internal/domain/model"
	"orders-service/internal/domain/users"
	"orders-service/internal/restaurants"
	"time"
)

type Service interface {
	CreateOrder(ctx context.Context, order *model.CreateOrderRequest)
}

type Repository interface {
	Save(ctx context.Context, order *model.Order)
}

type service struct {
	repo        Repository
	restaurants restaurants.Service
	users       users.Service
}

func (s service) CreateOrder(ctx context.Context, createOrderRequest *model.CreateOrderRequest) {
	menu := s.restaurants.GetMenu(ctx, createOrderRequest.RestaurantId)
	validateUser(ctx, createOrderRequest, s)

	notFound := validateOrder(menu, createOrderRequest)
	if len(notFound) != 0 {
		log.Fatal("unfulfillable order")
	}

	order := &model.Order{
		Id:           createOrderRequest.Id,
		UserId:       createOrderRequest.UserId,
		RestaurantId: createOrderRequest.RestaurantId,
		Items:        menu.Items,
		DateCreated:  time.Now(),
		LastUpdated:  time.Now(),
		Version:      0,
	}

	s.repo.Save(ctx, order)
}

func validateUser(ctx context.Context, createOrderRequest *model.CreateOrderRequest, s service) {
	user := s.users.GetUser(ctx, createOrderRequest.UserId)

	if user == nil {
		log.Fatal("user not found")
	}
}

func validateOrder(menu *model.Menu, createOrderRequest *model.CreateOrderRequest) []uint64 {
	found := false
	notFound := make([]uint64, 0)
	for _, requestItem := range createOrderRequest.Items {
		for _, restoItem := range menu.Items {
			if restoItem.Id == requestItem {
				found = true
			}
		}

		if found == false {
			notFound = append(notFound, requestItem)
		}

		found = false
	}

	return notFound
}

func NewService(repository Repository) Service {
	return &service{
		repo: repository,
	}
}
