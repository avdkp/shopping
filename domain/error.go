package domain

import "errors"

var (
	UserNotFoundError = errors.New("user not found")
	UnAuthorizedError = errors.New("login failed")
)
