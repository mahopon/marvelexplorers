package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		fmt.Printf("Handled request %s %s in %v\n", r.Method, r.URL.Path, duration)
	})
}
