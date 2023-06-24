package services

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"sync"

	"example.com/shopping/domain"
)

type userService struct {
	lock   sync.Mutex
	lastId int
	users  map[string]domain.User
}

type UserService interface {
	Create(user domain.User) error
}

func NewUserService() UserService {
	return &userService{
		lastId: 1,
		users:  make(map[string]domain.User),
	}
}

func hashPassword(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func (us *userService) Create(user domain.User) error {
	_, exists := us.users[user.UserName]
	if exists {
		return errors.New("user already exists")
	}
	us.lock.Lock()
	us.lastId++
	user.Id = us.lastId
	user.Password = hashPassword(user.Password)
	us.users[user.UserName] = user
	us.lock.Unlock()
	return nil
}
