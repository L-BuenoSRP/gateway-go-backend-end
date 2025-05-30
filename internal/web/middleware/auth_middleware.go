package middleware

import (
	"net/http"

	"github.com/L-BuenoSRP/imersao25/go-gateway/internal/domain"
	"github.com/L-BuenoSRP/imersao25/go-gateway/internal/service"
)

type AuthMiddleware struct {
	accountService service.AccountService
}

func NewAuthMiddleware(accountService *service.AccountService) *AuthMiddleware {
	return &AuthMiddleware{accountService: *accountService}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey == "" {
			http.Error(w, "X-API-KEY is required", http.StatusUnauthorized)
			return
		}

		_, err := m.accountService.FindByApiKey(apiKey)
		if err != nil {
			switch err {
			case domain.ErrAccountNotFound:
				http.Error(w, err.Error(), http.StatusUnauthorized)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		next.ServeHTTP(w, r)
	})
}
