package users

import (
	"context"
	"orders-service/internal/domain/model"
)

type Service interface {
	GetUser(ctx context.Context, userId uint64) *model.User
}
