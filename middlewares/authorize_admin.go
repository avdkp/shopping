package middlewares

import (
	"example.com/shopping/domain"
	"example.com/shopping/services"
	"net/http"
)

type authorizationMiddleware struct {
	authService services.AuthService
}

type AuthorizationMiddleware interface {
	AuthorizeAdmin(next http.Handler) http.Handler
	IsLoggedIn(next http.Handler) http.Handler
}

func NewAuthMiddleware(authSvc services.AuthService) AuthorizationMiddleware {
	return &authorizationMiddleware{
		authService: authSvc,
	}
}

func (aM *authorizationMiddleware) AuthorizeAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Auth-Token")
		err := aM.authService.IsAdmin(domain.Token(token))
		if err == nil {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
	})
}

func (aM *authorizationMiddleware) IsLoggedIn(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Auth-Token")
		err := aM.authService.IsLoggedIn(domain.Token(token))
		if err == nil {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
	})
}
