package routes

import (
	"github.com/gorilla/mux"
	"tcy/marvelexplorers/handler"
)

func RegisterCharacterRoutes(r *mux.Router, h *handler.CharacterHandler) {
	r.HandleFunc("/characters", h.GetCharacters).Methods("GET")
	r.HandleFunc("/characters/{partialName}", h.SearchCharacter).Methods("GET")
}
