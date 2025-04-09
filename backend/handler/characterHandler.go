package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	service "tcy/marvelexplorers/services"

	"github.com/gorilla/mux"
)

type CharacterHandler struct {
	Service *service.CharacterService
}

func (h *CharacterHandler) GetCharacters(w http.ResponseWriter, r *http.Request) {
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
	output, _ := json.Marshal(h.Service.GetCharacters(ctx, offset))
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func (h *CharacterHandler) SearchCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchStr := vars["partialName"]
	if searchStr == "" {
		http.Error(w, "No search parameter", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	output, _ := json.Marshal(h.Service.SearchCharacter(ctx, searchStr))
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
