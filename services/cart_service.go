package services

import "example.com/shopping/domain"

type cartService struct {
	userCart         map[int]int //map of userID to cartId, currently we are supporting 1-1 mapping
	carts            map[int]domain.Cart
	inventoryService InventoryService
}

type CartService interface {
	GetCart(userId int) domain.Cart
	AddToCart(userId int, items []int) ([]int, error)
}

func (cS *cartService) AddToCart(cartId int, items []int) ([]int, error) {
	cart, found := cS.carts[cartId]
	if !found {
		return []int{}, domain.CartNotFoundError
	}
	itemsCouldNotBeAdded := make([]int, 0)
	for _, itemId := range items {
		err := cS.inventoryService.MarkUnavailable(itemId)
		cart.Items[itemId] = struct{}{}
		if err != nil {
			itemsCouldNotBeAdded = append(itemsCouldNotBeAdded, itemId)
		}
	}
	return itemsCouldNotBeAdded, nil
}

func (cS *cartService) GetCart(userId int) (domain.Cart, error) {
	cartId, found := cS.userCart[userId]
	if !found {
		return domain.Cart{}, domain.CartNotFoundError
	}
	cart, found := cS.carts[cartId]
	if !found {
		return domain.Cart{}, domain.CartNotFoundError
	}
	return cart, nil
}
