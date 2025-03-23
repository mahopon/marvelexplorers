package routes

import (
	"github.com/gorilla/mux"
	"tcy/marvelexplorers/handler"
)

func RegisterStoryRoutes(r *mux.Router) {
	r.HandleFunc("/stories", handler.GetStories).Methods("GET")
}
