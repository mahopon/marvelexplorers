package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	db "tcy/marvelexplorers/repository"

	"github.com/gorilla/mux"
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

func SearchCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchStr := vars["partialName"]
	if searchStr == "" {
		http.Error(w, "No search parameter", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	output, _ := json.Marshal(db.GetPG().SearchCharacter(ctx, searchStr))
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
