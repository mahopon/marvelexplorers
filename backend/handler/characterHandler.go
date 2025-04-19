package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	model "tcy/marvelexplorers/model/db"
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
	output, _ := h.Service.GetCharactersWithCache(ctx, offset)
	characters, ok := output.([]model.Character_db)
	if !ok {
		http.Error(w, "Failed to parse characters", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(characters)
	if err != nil {
		http.Error(w, "Failed to encode characters", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
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
	result, _ := h.Service.SearchCharacterWithCache(ctx, searchStr)
	output, _ := json.Marshal(result)
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
