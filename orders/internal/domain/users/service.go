package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"orders-service/internal/domain/model"
)

var (
	getUrl = "http://users-service/users/%s"
)

type Service interface {
	GetUser(ctx context.Context, userId uint64) (*model.User, error)
}

type service struct {
	client *http.Client
}

func (s service) GetUser(ctx context.Context, userId uint64) (*model.User, error) {
	r, err := s.client.Get(fmt.Sprintf(getUrl, userId))

	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	var user *model.User
	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	return user, nil
}

func NewUserService() Service {
	return &service{
		client: http.DefaultClient,
	}
}
