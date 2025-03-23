package routes

import (
	"fmt"
	"net/http"
	"tcy/marvelexplorers/handler"
	"tcy/marvelexplorers/middleware"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func Setup() {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.Custom404Handler)
	r.HandleFunc("/favicon.ico", handler.GetFavicon).Methods("GET")

	apiRouter := r.PathPrefix("/api").Subrouter()
	RegisterCharacterRoutes(apiRouter)

	muxWithMiddleware := middleware.LogMiddleware(r)
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Started server on localhost:8000")
	http.ListenAndServe(":8000", muxWithMiddleware)
}
