package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"tcy/marvelexplorers/handler"
	"tcy/marvelexplorers/middleware"
	model "tcy/marvelexplorers/model/db"
	dbRepo "tcy/marvelexplorers/repository/postgres"
	redisRepo "tcy/marvelexplorers/repository/redis"
	"tcy/marvelexplorers/services"
)

func Setup() http.Handler {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/", handler.Custom404Handler)
	r.HandleFunc("/favicon.ico", handler.GetFavicon).Methods("GET")
	apiRouter := r.PathPrefix("/api").Subrouter()
	RegisterCharacterRoutes(apiRouter, &handler.CharacterHandler{Service: services.GetCharacterService(
		dbRepo.NewPGRepo[model.Character_db](),
		redisRepo.NewRedisRepo[model.Character_db](),
	)})
	// RegisterEventRoutes(apiRouter)
	// RegisterSeriesRoutes(apiRouter)
	// RegisterStoryRoutes(apiRouter)

	muxWithMiddleware := middleware.ApplyMiddleware(r)
	return muxWithMiddleware
}
