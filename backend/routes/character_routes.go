package routes

import (
	"github.com/gorilla/mux"
	"tcy/marvelexplorers/handler"
	model "tcy/marvelexplorers/model/db"
	dbRepo "tcy/marvelexplorers/repository/postgres"
	redisRepo "tcy/marvelexplorers/repository/redis"
	"tcy/marvelexplorers/services"
)

func RegisterCharacterRoutes(r *mux.Router) {
	h := &handler.CharacterHandler{Service: services.GetCharacterService(
		dbRepo.NewPGRepo[model.Character_db](),
		redisRepo.NewRedisRepo[model.Character_db](),
	)}
	r.HandleFunc("/characters", h.GetCharacters).Methods("GET")
	r.HandleFunc("/characters/{partialName}", h.SearchCharacter).Methods("GET")
}
