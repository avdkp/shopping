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
	token domain.Token
}

type userService struct {
	lock           sync.Mutex
	lastId         int
	users          map[string]userDetails
	suspendedUsers map[string]userDetails
	tokens         map[domain.Token]string
}

type UserService interface {
	Create(user domain.User) error
	Login(user domain.User) (domain.Token, error)
	SuspendUser(userId string) error
}

var usrSvc *userService

type AuthService interface {
	MatchRoleType(token domain.Token, role string) (int, error)
}

func initializeUserService() {
	usrSvc = &userService{
		lastId: 1,
		users:  make(map[string]userDetails),
		tokens: make(map[domain.Token]string),
	}
}
func GetAuthService() AuthService {
	if usrSvc == nil {
		initializeUserService()
	}
	return usrSvc
}
func GetUserService() UserService {
	if usrSvc == nil {
		initializeUserService()
	}
	return usrSvc
}

func RandomString(length int) string {
	b := make([]byte, length)
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func (us *userService) generateRandomToken() domain.Token {
	for true {
		token := domain.Token(RandomString(tokenLength))
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

func (us *userService) Login(user domain.User) (domain.Token, error) {
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

func (us *userService) MatchRoleType(token domain.Token, role string) (int, error) {
	userName, exists := us.tokens[token]
	if !exists {
		return -1, domain.InvalidTokenError
	}
	user, exist := us.users[userName]
	if !exist {
		return -1, domain.InvalidTokenError
	}
	if user.Role == role {
		return user.Id, nil
	}
	return -1, domain.UnAuthorizedError
}

func (us *userService) SuspendUser(userName string) error {
	user, found := us.users[userName]
	if !found {
		return domain.UserNotFoundError
	}
	delete(us.users, userName)
	delete(us.tokens, user.token)
	us.suspendedUsers[userName] = user
	return nil
}
