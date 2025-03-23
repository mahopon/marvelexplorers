package handler

import (
	"context"
	"encoding/json"
	"net/http"
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
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	output, _ := json.Marshal(db.GetPG().GetCharacters(ctx, offset))
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
