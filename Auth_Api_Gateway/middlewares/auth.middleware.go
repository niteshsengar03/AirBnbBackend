package middlewares

import (
	"Auth_Api_Gateway/config"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorisation header is required", http.StatusUnauthorized)
			return
		}
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Authorisation header must start with Bearer", http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			http.Error(w, "Token is required", http.StatusUnauthorized)
			return
		}
		claims := jwt.MapClaims{}
		_,err := jwt.ParseWithClaims(token , &claims, func(token *jwt.Token)(interface{},error){
			return []byte(config.GetString("JWT_SECRET_KEY", "YourToken")),nil
		})
		if err!=nil{
			http.Error(w,"Invalid token claims",http.StatusUnauthorized)
		}
		userId,okId := claims["id"].(float64)
		userEmail,okEmail := claims["email"].(string)
		if !okId || !okEmail {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}
		fmt.Println("UserId: ",userId,"UserEmail: ",userEmail)
		next.ServeHTTP(w, r)
	})

}
