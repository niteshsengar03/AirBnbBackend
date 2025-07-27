package middlewares

import (
	"Auth_Api_Gateway/config"
	"context"
	"fmt"
	"net/http"
	"strconv"
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

		// decoding the values from jwt 
		// id was intially int but when we decode in jwt internally it converts to 
		// float64  and writting .float64 is not i'm converting 
		// it's like saying that i know claims["id"] will return flaot64 so instead of 
		// keeping userId type any we define it.
		userId,okId := claims["id"].(float64)
		userEmail,okEmail := claims["email"].(string)

		if !okId || !okEmail {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}
		fmt.Println("UserId: ",int64(userId),"UserEmail: ",userEmail)
		ctx := context.WithValue(r.Context(),"userId", strconv.FormatFloat(userId,'f',0,64))
		ctx = context.WithValue(ctx,"email", userEmail)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
