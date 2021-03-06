package restaurants

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"orders-service/internal/domain/model"
)

var (
	getUrl = "http://restaurant-service:8082/restaurants/%v/menu"
)

type Service interface {
	GetMenu(ctx context.Context, restaurantId uint64) (*model.Menu, error)
}

type service struct {
	client *http.Client
}

func (s service) GetMenu(ctx context.Context, restaurantId uint64) (*model.Menu, error) {
	url := fmt.Sprintf(getUrl, restaurantId)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("cannot create request for get menu for restaurant id %v: %v", restaurantId, err.Error())
	}
	response, err := s.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("cannot get menu for restaurant id %v: %v", restaurantId, err.Error())
	}
	if response.StatusCode != http.StatusOK {
		errorBody, err := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("cannot read body error menu for restaurant id %v: %v - %v", restaurantId, response.StatusCode, err.Error())
		}
		return nil, fmt.Errorf("cannot get menu for restaurant id %v: %v - %v", restaurantId, response.StatusCode, string(errorBody))
	}

	// b, _ := ioutil.ReadAll(response.Body)
	// log.Println(string(b))
	var menu *model.Menu
	decoder := json.NewDecoder(response.Body)

	err = decoder.Decode(&menu)
	if err != nil {
		return nil, fmt.Errorf("cannot parse body menu for restaurant id %v: %v", restaurantId, err)
	}
	defer response.Body.Close()
	return menu, nil
}

func NewService() Service {
	return &service{
		client: http.DefaultClient,
	}
}
