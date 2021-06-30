package restaurants

import (
	"context"
	"orders-service/internal/domain/model"
)

type Service interface {
	GetMenu(ctx context.Context, restaurantId uint64) *model.Menu
}

type service struct {
}

func (s service) GetMenu(ctx context.Context, restaurantId uint64) *model.Menu {
	panic("implement me")
}

func NewService() Service {
	return &service{}
}
