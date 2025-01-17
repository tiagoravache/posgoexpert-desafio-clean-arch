package web

import (
	"encoding/json"
	"net/http"

	"desafio-cleanarch/internal/entity"
	"desafio-cleanarch/internal/usecase"
)

type WebListOrdersHandler struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewWebListOrdersHandler(
	OrderRepository entity.OrderRepositoryInterface,
) *WebListOrdersHandler {
	return &WebListOrdersHandler{
		OrderRepository: OrderRepository,
	}
}

func (h *WebListOrdersHandler) List(w http.ResponseWriter, r *http.Request) {
	listOrders := usecase.NewListOrdersUseCase(h.OrderRepository)
	output, err := listOrders.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
