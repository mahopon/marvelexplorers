package routes

import (
	"github.com/gorilla/mux"
	"tcy/marvelexplorers/handler"
)

func RegisterSeriesRoutes(r *mux.Router) {
	r.HandleFunc("/series", handler.GetSeries).Methods("GET")
}
