package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	handler "tcy/marvelexplorers/handler"
	db "tcy/marvelexplorers/repository"
	router "tcy/marvelexplorers/routes"
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
	ctx := context.Background()
	_, err = db.NewPG(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	router.Setup()
}
