package handler

import (
	"context"
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
	output, err := h.Service.GetCharactersWithCache(ctx, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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
	result, err := h.Service.SearchCharacterWithCache(ctx, searchStr)
	if err != nil {
		if err.Error() == "No data" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
