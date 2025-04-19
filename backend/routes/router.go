package routes

import (
	"fmt"
	"net/http"
	"tcy/marvelexplorers/handler"
	"tcy/marvelexplorers/middleware"
	db "tcy/marvelexplorers/repository/postgres"
	"tcy/marvelexplorers/services"

	"github.com/gorilla/mux"
)

func Setup() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/", handler.Custom404Handler)
	r.HandleFunc("/favicon.ico", handler.GetFavicon).Methods("GET")
	pgRepo := db.GetPG()
	characterHandler := &handler.CharacterHandler{
		Service: &services.CharacterService{
			DBRepo: pgRepo,
		},
	}
	apiRouter := r.PathPrefix("/api").Subrouter()
	RegisterCharacterRoutes(apiRouter, characterHandler)
	// RegisterEventRoutes(apiRouter)
	// RegisterSeriesRoutes(apiRouter)
	// RegisterStoryRoutes(apiRouter)

	muxWithMiddleware := middleware.ApplyMiddleware(r)
	fmt.Println("Started server on localhost:8000")
	http.ListenAndServe(":8000", muxWithMiddleware)
}
