package middleware

import (
	"context"
	"net/http"
	"github.com/SyydMR/To-Do-List/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "authorization token not provided", http.StatusUnauthorized)
			return
		}

		userID, err := utils.VerifyJWT(tokenString)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}