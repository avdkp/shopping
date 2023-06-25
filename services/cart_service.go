package services

import "example.com/shopping/domain"

type cartService struct {
	carts            map[int]domain.Cart
	inventoryService InventoryService
}

type CartService interface {
	AddToCart(userId int, items []int) ([]int, error)
	RemoveFromCart(userId int, items []int) ([]int, error)
}

func NewCartService(invSvc InventoryService) CartService {
	return &cartService{
		carts:            make(map[int]domain.Cart),
		inventoryService: invSvc,
	}
}

func (cS *cartService) AddToCart(userId int, items []int) ([]int, error) {
	cart, found := cS.carts[userId]
	if !found {
		cart = createCart(userId)
		cS.carts[userId] = cart
	}
	itemsCouldNotBeAdded := make([]int, 0)
	for _, itemId := range items {
		err := cS.inventoryService.MarkUnavailable(itemId, userId)
		cart.Items[itemId] = struct{}{}
		if err != nil {
			itemsCouldNotBeAdded = append(itemsCouldNotBeAdded, itemId)
		}
	}
	return itemsCouldNotBeAdded, nil
}

func createCart(id int) domain.Cart {
	return domain.Cart{
		UserID: id,
		Items:  make(map[int]struct{}),
	}
}

func (cS *cartService) RemoveFromCart(userId int, items []int) ([]int, error) {
	cart, found := cS.carts[userId]
	if !found {
		return []int{}, domain.CartNotFoundError
	}

	itemsCouldNotBeRemoved := make([]int, 0)
	for _, itemId := range items {
		err := cS.inventoryService.MarkAvailable(itemId, userId)
		delete(cart.Items, itemId)
		if err != nil {
			itemsCouldNotBeRemoved = append(itemsCouldNotBeRemoved, itemId)
		}
	}
	return itemsCouldNotBeRemoved, nil
}
