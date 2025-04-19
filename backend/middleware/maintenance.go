package middleware

import (
	"net/http"
	"tcy/marvelexplorers/utils"
)

func maintenanceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if utils.CheckMaintenance() {
			http.Error(w, "Service is in maintenance.", http.StatusServiceUnavailable)
			return
		}
		next.ServeHTTP(w, r)
	})
}
