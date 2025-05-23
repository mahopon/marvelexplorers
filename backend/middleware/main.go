package middleware

import (
	"net/http"
)

func ApplyMiddleware(next http.Handler) http.Handler {
	return LogMiddleware(maintenanceMiddleware(CORSMiddleware(next)))
}
