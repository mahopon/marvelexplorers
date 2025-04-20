package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"tcy/marvelexplorers/handler"
	"tcy/marvelexplorers/middleware"
)

func Setup() http.Handler {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/", handler.Custom404Handler)
	r.HandleFunc("/favicon.ico", handler.GetFavicon).Methods("GET")
	apiRouter := r.PathPrefix("/api").Subrouter()
	RegisterCharacterRoutes(apiRouter)
	// RegisterEventRoutes(apiRouter)
	// RegisterSeriesRoutes(apiRouter)
	// RegisterStoryRoutes(apiRouter)

	muxWithMiddleware := middleware.ApplyMiddleware(r)
	return muxWithMiddleware
}
