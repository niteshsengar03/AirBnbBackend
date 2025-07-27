package middlewares

import (
	"net/http"
	"time"
	"golang.org/x/time/rate"
)

var limitter = rate.NewLimiter(rate.Every(1*time.Minute), 5) // 5 req per min

func RateLimitter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limitter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
