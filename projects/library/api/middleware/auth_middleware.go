package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)


func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		JWT_SECRET := os.Getenv("JWT_SECRET")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return []byte(JWT_SECRET), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			ctx := context.WithValue(r.Context(), "userID", claims["sub"])
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}