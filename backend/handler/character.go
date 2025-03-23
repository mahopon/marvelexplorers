package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	db "tcy/marvelexplorers/repository"
)

func GetCharacters(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/characters" {
		Custom404Handler(w, r)
		return
	}
	offsetStr := r.URL.Query().Get("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}
	offset, _ := strconv.Atoi(offsetStr)
	ctx := context.Background()
	pg, err := db.NewPG(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	output, _ := json.Marshal(pg.GetCharacters(ctx, offset))
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
