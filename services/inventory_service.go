package services

import (
	"example.com/shopping/domain"
	"sync"
)

type inventoryService struct {
	lock   sync.Mutex
	lastId int
	items  map[domain.ItemName][]domain.Item
}

type InventoryService interface {
	AddItems(items []domain.Item) error
}

func NewInventoryService() InventoryService {
	return &inventoryService{
		lastId: 0,
		items:  make(map[domain.ItemName][]domain.Item),
	}
}

func (is *inventoryService) addItemInInventory(item domain.Item) {
	is.lock.Lock()
	is.lastId++
	item.Id = is.lastId
	existingItemsWithSameName, exists := is.items[item.Name]
	if !exists {
		existingItemsWithSameName = make([]domain.Item, 1)
	}
	existingItemsWithSameName = append(existingItemsWithSameName, item)
	is.lock.Unlock()
}

func (is *inventoryService) AddItems(items []domain.Item) error {
	for _, item := range items {
		is.addItemInInventory(item)
	}
	return nil
}
