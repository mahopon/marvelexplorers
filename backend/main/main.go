package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"strconv"
	"tcy/marvelexplorers/pg"
	"time"
)

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		fmt.Printf("Handled request %s %s in %v\n", r.Method, r.URL.Path, duration)
	})
}

func custom404Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<html><head><title>404 not found</title></head<body><h1>404 - Page Not Found</h1><p>Sorry, we couldn't find the page you're looking for.</p></body></html>")
}

func charHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/characters" {
		custom404Handler(w, r)
		return
	}
	offsetStr := r.URL.Query().Get("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}
	offset, _ := strconv.Atoi(offsetStr)
	ctx := context.Background()
	pg, err := pg.NewPG(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	output, _ := json.Marshal(pg.GetCharacters(ctx, offset))
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/characters", charHandler)
	mux.HandleFunc("/", custom404Handler)

	muxWithMiddleware := logMiddleware(mux)
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	http.ListenAndServe(":8000", muxWithMiddleware)
}
