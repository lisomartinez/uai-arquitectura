package restaurants

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"orders-service/internal/domain/model"
)

var (
	getUrl = ""
)

type Service interface {
	GetMenu(ctx context.Context, restaurantId uint64) (*model.Menu, error)
}

type service struct {
	client *http.Client
}

func (s service) GetMenu(ctx context.Context, restaurantId uint64) (*model.Menu, error) {
	request, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("cannot create request for get menu for restaurant id %v: %v", restaurantId, err.Error()))
	}
	response, err := s.client.Do(request)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("cannot get menu for restaurant id %v: %v", restaurantId, err.Error()))
	}
	if response.StatusCode != http.StatusOK {
		errorBody, err := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			return nil, errors.New(fmt.Sprintf("cannot read body error menu for restaurant id %v: %v - %v", restaurantId, response.StatusCode, err.Error()))
		}
		return nil, errors.New(fmt.Sprintf("cannot get menu for restaurant id %v: %v - %v", restaurantId, response.StatusCode, string(errorBody)))
	}

	var menu *model.Menu
	decoder := json.NewDecoder(response.Body)

	err = decoder.Decode(menu)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("cannot parse body menu for restaurant id %v", restaurantId))
	}
	defer response.Body.Close()
	return menu, nil
}

func NewService() Service {
	return &service{
		client: http.DefaultClient,
	}
}
