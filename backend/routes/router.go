package routes

import (
	"fmt"
	"net/http"
	"tcy/marvelexplorers/handler"
	"tcy/marvelexplorers/middleware"

	"github.com/gorilla/mux"
)

func Setup() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/", handler.Custom404Handler)
	r.HandleFunc("/favicon.ico", handler.GetFavicon).Methods("GET")

	apiRouter := r.PathPrefix("/api").Subrouter()
	RegisterCharacterRoutes(apiRouter)
	RegisterEventRoutes(apiRouter)
	RegisterSeriesRoutes(apiRouter)
	RegisterStoryRoutes(apiRouter)

	muxWithMiddleware := middleware.ApplyMiddleware(r)
	fmt.Println("Started server on localhost:8000")
	http.ListenAndServe(":8000", muxWithMiddleware)
}
