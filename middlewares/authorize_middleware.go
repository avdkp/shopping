package middlewares

import (
	"example.com/shopping/domain"
	"example.com/shopping/services"
	"fmt"
	"net/http"
)

type authorizationMiddleware struct {
	authService services.AuthService
}

type AuthorizationMiddleware interface {
	AuthorizationCreator(role string) func(next http.Handler) http.Handler
}

func NewAuthMiddleware(authSvc services.AuthService) AuthorizationMiddleware {
	return &authorizationMiddleware{
		authService: authSvc,
	}
}

func (aM *authorizationMiddleware) AuthorizationCreator(roleType string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Auth-Token")
			userId, err := aM.authService.MatchRoleType(domain.Token(token), roleType)
			r.Header.Set("X-User-Id", fmt.Sprintf("%d", userId))
			if err == nil {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, err.Error(), http.StatusUnauthorized)
			}
		})
	}
}
