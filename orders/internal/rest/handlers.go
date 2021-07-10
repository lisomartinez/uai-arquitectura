package rest

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"orders-service/internal/domain/model"
	"orders-service/internal/tools"
	"strconv"

	"github.com/go-chi/chi"
)

type Service interface {
	CreateOrder(ctx context.Context, order *model.CreateOrderRequest) (*model.Order, error)
	GetOrder(ctx context.Context, orderId uint64) (*model.Order, error)
}

type OrdersHandler struct {
	service Service
}

func NewOrdersHandler(service Service) *OrdersHandler {
	return &OrdersHandler{
		service: service,
	}
}

func (oh *OrdersHandler) createOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	createOrderRequest := &model.CreateOrderRequest{}
	err := json.NewDecoder(r.Body).Decode(createOrderRequest)

	if err != nil {
		ErrorResponse(w, 400, "cannot parse request body")
	}

	order, err := oh.service.CreateOrder(ctx, createOrderRequest)

	var responseError tools.Custom
	if errors.As(err, &responseError) {
		ErrorResponse(w, responseError.Status(), err.Error())
	} else {
		ErrorResponse(w, 500, "internal server error")
	}

	WebResponse(w, 201, order)
}

func (oh *OrdersHandler) getOrder(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	if idParam == "" {
		ErrorResponse(w, 400, "id is mandatory")
		return
	}

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ErrorResponse(w, 400, "id is mandatory")
		return
	}

	ctx := r.Context()
	order, err := oh.service.GetOrder(ctx, id)

	var responseError tools.Custom

	if err == nil {
		WebResponse(w, 200, order)
	} else if errors.As(err, &responseError) {
		ErrorResponse(w, responseError.Status(), err.Error())
	} else {
		ErrorResponse(w, 500, "internal server error")
	}

}
