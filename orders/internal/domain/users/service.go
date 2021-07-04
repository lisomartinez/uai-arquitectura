package users

import (
	"context"
	"orders-service/internal/domain/model"
)

type Service interface {
	GetUser(ctx context.Context, userId uint64) (*model.User, error)
}

type service struct {
}

func (s service) GetUser(ctx context.Context, userId uint64) (*model.User, error) {
	panic("implement me")
}

func NewUserService() Service {
	return &service{}
}
