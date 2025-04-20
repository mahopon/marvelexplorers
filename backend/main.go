package main

import (
	"embed"
	"fmt"
	"net/http"
	handler "tcy/marvelexplorers/handler"
	router "tcy/marvelexplorers/routes"

	"github.com/joho/godotenv"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	handler.StaticFiles = staticFiles

	server := router.Setup()
	fmt.Println("Started server on localhost:8000")
	http.ListenAndServe(":8000", server)
}
