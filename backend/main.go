package main

import (
	"embed"
	handler "tcy/marvelexplorers/handler"
	router "tcy/marvelexplorers/routes"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	handler.StaticFiles = staticFiles
	router.Setup()
}
