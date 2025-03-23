package routes

import (
	"github.com/gorilla/mux"
	"tcy/marvelexplorers/handler"
)

func RegisterCharacterRoutes(r *mux.Router) {
	r.HandleFunc("/characters", handler.GetCharacters).Methods("GET")
}
