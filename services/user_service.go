package services

import (
	"example.com/shopping/domain"
	"math/rand"
	"sync"
)

const (
	tokenLength = 10
)

type userDetails struct {
	domain.User
	token string
}

type userService struct {
	lock   sync.Mutex
	lastId int
	users  map[string]userDetails
	tokens map[string]string
}

type UserService interface {
	Create(user domain.User) error
	Login(user domain.User) (string, error)
}

func NewUserService() UserService {
	return &userService{
		lastId: 1,
		users:  make(map[string]userDetails),
		tokens: make(map[string]string),
	}
}

func randomString(length int) string {
	b := make([]byte, length)
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func (us *userService) generateRandomToken() string {
	for true {
		token := randomString(tokenLength)
		_, exists := us.tokens[token]
		if !exists {
			return token
		}
	}
	return ""
}

func (us *userService) Create(user domain.User) error {
	_, exists := us.users[user.UserName]
	if exists {
		return domain.UserNotFoundError
	}
	us.lock.Lock()
	us.lastId++
	user.Id = us.lastId
	us.users[user.UserName] = userDetails{
		User: user,
	}
	us.lock.Unlock()
	return nil
}

func (us *userService) Login(user domain.User) (string, error) {
	user1, found := us.users[user.UserName]
	if !found {
		return "", domain.UserNotFoundError
	}
	if user.HashedPassword == user1.HashedPassword {
		if user1.token != "" {
			return user1.token, nil
		}
		token := us.generateRandomToken()
		user1.token = token
		us.lock.Lock()
		us.tokens[token] = user.UserName
		us.users[user.UserName] = user1
		us.lock.Unlock()
		return token, nil
	}
	return "", domain.UnAuthorizedError
}
