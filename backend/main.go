package main

import (
	"context"
	"embed"
	"fmt"
	"os"
	handler "tcy/marvelexplorers/handler"
	db "tcy/marvelexplorers/repository/postgres"
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
	ctx := context.Background()
	_, err = db.NewPG(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	router.Setup()
}
