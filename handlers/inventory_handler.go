package handlers

import (
	"encoding/json"
	"example.com/shopping/domain"
	"example.com/shopping/services"
	"fmt"
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

func (iH *inventoryHandler) AddItems(w http.ResponseWriter, r *http.Request) {
	var req []domain.Item

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = iH.inventoryService.AddItems(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (iH *inventoryHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	items := iH.inventoryService.GetAllItems()

	itemsJson, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(itemsJson)
	if err != nil {
		//TODO log error instead
		fmt.Print("could not write the response")
	}
}
