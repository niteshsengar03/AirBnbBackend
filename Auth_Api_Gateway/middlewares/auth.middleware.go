package middlewares

import (
	"Auth_Api_Gateway/config"
	dbConfig "Auth_Api_Gateway/config/db"
	db "Auth_Api_Gateway/db/repositories"
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
		_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetString("JWT_SECRET_KEY", "YourToken")), nil
		})
		if err != nil {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		}

		// decoding the values from jwt
		// id was intially int but when we decode in jwt internally it converts to
		// float64  and writting .float64 is not i'm converting
		// it's like saying that i know claims["id"] will return flaot64 so instead of
		// keeping userId type any we define it.
		userId, okId := claims["id"].(float64)
		userEmail, okEmail := claims["email"].(string)

		if !okId || !okEmail {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		fmt.Println("UserId: ", int64(userId), "UserEmail: ", userEmail)

		ctx := context.WithValue(r.Context(), "userId", strconv.FormatFloat(userId, 'f', 0, 64))
		ctx = context.WithValue(ctx, "email", userEmail)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

// roles is a string of []
// return type is a function which is a middleware
func RequireAllRoles(roles ...string) func(next http.Handler) http.Handler {
	// function that create a middleware for checking the above set of rules

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userIdStr := r.Context().Value("userID").(string)
			// string to integer
			userId, err := strconv.ParseInt(userIdStr, 10, 64)
			if err != nil {
				http.Error(w, "Invalid userId", http.StatusUnauthorized)
				return
			}
			DB, err := dbConfig.SetupDB()
			if err != nil {
				http.Error(w, "Cannot connect to database", http.StatusInternalServerError)
				return
			}

			urr := db.NewUserRoleRepository(DB)
			HasAllRoles, HasAllRolesErr := urr.HasAllRoles(userId, roles)
			if HasAllRolesErr != nil {
				http.Error(w, "Error checking user roles", http.StatusForbidden)
				return
			}
			if !HasAllRoles {
				http.Error(w, "Forbidden: You don't have forbidden roles", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})

	}

}
