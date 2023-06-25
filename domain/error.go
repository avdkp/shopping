package domain

import "errors"

var (
	UserNotFoundError    = errors.New("user not found")
	CartNotFoundError    = errors.New("cart not found")
	ItemNotFoundError    = errors.New("item not found")
	ItemUnAvailableError = errors.New("item not available")
	InvalidTokenError    = errors.New("token invalid")
	UnAuthorizedError    = errors.New("request unauthorized")
)
