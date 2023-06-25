package handlers

import (
	"encoding/json"
	"example.com/shopping/domain"
	"example.com/shopping/services"
	"net/http"
)

type inventoryHandler struct {
	inventoryService services.InventoryService
}

type InventoryHandler interface {
	AddItems(w http.ResponseWriter, r *http.Request)
	GetItems(w http.ResponseWriter, r *http.Request)
}

func NewInventoryHandler(svc services.InventoryService) InventoryHandler {
	return &inventoryHandler{
		inventoryService: svc,
	}
}

func (sh *inventoryHandler) AddItems(w http.ResponseWriter, r *http.Request) {
	var req []domain.Item

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = sh.inventoryService.AddItems(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (sh *inventoryHandler) GetItems(w http.ResponseWriter, r *http.Request) {

}
