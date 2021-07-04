package repository

import (
	"context"
	"fmt"
	"orders-service/internal/domain/model"
	"orders-service/internal/tools"
)

type Repository interface {
	Save(ctx context.Context, order *model.Order) (*model.Order, error)
	GetOrder(ctx context.Context, orderId uint64) (*model.Order, error)
}

type localRepository struct {
	orders map[uint64]*model.Order
}

func (l *localRepository) Save(ctx context.Context, order *model.Order) (*model.Order, error) {
	l.orders[order.OrderID] = order
	return order, nil
}

func (l *localRepository) GetOrder(ctx context.Context, orderId uint64) (*model.Order, error) {
	order, ok := l.orders[orderId]

	if !ok {
		return nil, tools.NewBusinessError(fmt.Sprintf("order not found: %v", orderId))
	}

	return order, nil
}

func NewLocalOrderRepository() Repository {
	return &localRepository{
		orders: map[uint64]*model.Order{},
	}
}
