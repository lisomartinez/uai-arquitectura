package orders

import (
	"context"
	"fmt"
	"math/rand"
	"orders-service/internal/domain/model"
	"orders-service/internal/domain/restaurants"
	"orders-service/internal/domain/users"
	"orders-service/internal/tools"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	CreateOrder(ctx context.Context, order *model.CreateOrderRequest) (*model.Order, error)
	GetOrder(ctx context.Context, orderId uint64) (*model.Order, error)
}

type Repository interface {
	Save(ctx context.Context, order *model.Order) (*model.Order, error)
	GetOrder(ctx context.Context, orderId uint64) (*model.Order, error)
}

type service struct {
	repo        Repository
	restaurants restaurants.Service
	users       users.Service
}

func (s service) GetOrder(ctx context.Context, orderId uint64) (*model.Order, error) {
	order, err := s.repo.GetOrder(ctx, orderId)

	if err == mongo.ErrNoDocuments {
		return nil, tools.NewCustom(404, fmt.Sprintf("document not found: %v", orderId))
	}

	return order, err
}

func (s service) CreateOrder(ctx context.Context, createOrderRequest *model.CreateOrderRequest) (*model.Order, error) {
	menu, err := s.restaurants.GetMenu(ctx, createOrderRequest.RestaurantID)

	if err != nil {
		return nil, err
	}

	err = validateUser(ctx, createOrderRequest, s)

	if err != nil {
		return nil, tools.NewCustom(400, fmt.Sprintf("invalid user: %v", err))
	}

	err = validateOrder(menu, createOrderRequest)

	if err != nil {
		return nil, tools.NewCustom(400, fmt.Sprintf("invalid order: %v", err))
	}

	order := &model.Order{
		OrderID:      createOrderRequest.OrderID,
		UserID:       createOrderRequest.UserID,
		RestaurantID: createOrderRequest.RestaurantID,
		Items:        menu.Items,
		DateCreated:  time.Now(),
		LastUpdated:  time.Now(),
		Total:        200 + rand.Float64()*300,
		Version:      0,
	}

	return s.repo.Save(ctx, order)
}

func validateUser(ctx context.Context, createOrderRequest *model.CreateOrderRequest, s service) error {
	user, err := s.users.GetUser(ctx, createOrderRequest.UserID)

	if err != nil {
		return err
	}

	if user == nil || user.UserID != createOrderRequest.UserID {
		return fmt.Errorf("invalid user for %v", createOrderRequest)
	}

	return nil
}

func validateOrder(menu *model.Menu, createOrderRequest *model.CreateOrderRequest) error {
	found := false
	notFound := make([]uint64, 0)
	for _, requestItem := range createOrderRequest.Items {
		for _, restoItem := range menu.Items {
			if restoItem.ID == requestItem {
				found = true
			}
		}

		if found == false {
			notFound = append(notFound, requestItem)
		}

		found = false
	}
	if len(notFound) == 0 {
		return nil
	} else {
		return fmt.Errorf("missing items: %v", notFound)
	}
}

func NewService(repository Repository, restaurants restaurants.Service, users users.Service) Service {
	return &service{
		repo:        repository,
		restaurants: restaurants,
		users:       users,
	}
}
