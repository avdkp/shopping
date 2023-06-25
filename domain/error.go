package domain

import "errors"

var (
	UserNotFoundError = errors.New("user not found")
	InvalidTokenError = errors.New("token invalid")
	UnAuthorizedError = errors.New("login failed")
)
