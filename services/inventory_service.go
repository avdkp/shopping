package services

import (
	"example.com/shopping/domain"
	"sync"
)

type inventoryService struct {
	lock      sync.Mutex
	lastId    int
	items     map[int]domain.Item
	itemLocks sync.Map
}

type InventoryService interface {
	AddItems(items []domain.Item) error
	GetAllItems() []domain.Item
	MarkUnavailable(itemId int) error
}

func NewInventoryService() InventoryService {
	return &inventoryService{
		lastId: 0,
		items:  make(map[int]domain.Item),
	}
}

func (iS *inventoryService) addItemInInventory(item domain.Item) {
	iS.lock.Lock()
	iS.lastId++
	item.Id = iS.lastId
	iS.items[item.Id] = item
	iS.lock.Unlock()
}

func (iS *inventoryService) AddItems(items []domain.Item) error {
	for _, item := range items {
		item.Available = true
		iS.addItemInInventory(item)
	}
	return nil
}

func (iS *inventoryService) GetAllItems() []domain.Item {
	result := make([]domain.Item, 0)
	for _, item := range iS.items {
		if item.Available {
			result = append(result, item)
		}
	}
	return result
}

func (iS *inventoryService) MarkUnavailable(itemId int) error {
	item, found := iS.items[itemId]
	if !found {
		return domain.ItemNotFoundError
	}
	if item.Available == false {
		return domain.ItemUnAvailableError
	}

	randomKey := RandomString(10)
	_, alreadyPresent := iS.itemLocks.LoadOrStore(itemId, randomKey)
	if alreadyPresent {
		//Lock was acquired by someone else
		return domain.ItemUnAvailableError
	}
	//lock was acquired for itemID
	defer iS.itemLocks.Delete(itemId)

	item.Available = false
	iS.items[item.Id] = item
	return nil
}
