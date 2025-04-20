package main

import (
	"embed"
	"fmt"
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
	router.Setup()
}
