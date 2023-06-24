package services

import (
	"errors"

	"example.com/shopping/domain"
)

type userService struct {
	lastId int
	users  map[int]domain.User
}

type UserService interface {
	Create(user domain.User) error
}

func NewUserService() UserService {
	return &userService{
		lastId: 1,
		users:  make(map[int]domain.User),
	}
}

func (us *userService) Create(user domain.User) error {
	_, exists := us.users[user.Id]
	if exists {
		return errors.New("user already exists")
	}
	user.Id = us.lastId
	us.lastId++
	// user.Password = md5.
	us.users[user.Id] = user
	return nil
}
