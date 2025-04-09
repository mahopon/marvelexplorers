package services

import (
	"context"
	model "tcy/marvelexplorers/model/db"
	repo "tcy/marvelexplorers/repository"
)

type CharacterService struct {
	Repo repo.CharacterRepo
}

func (s CharacterService) GetCharacters(ctx context.Context, offset int) interface{} {
	return s.Repo.GetCharacters(ctx, offset)
}

func (s CharacterService) SearchCharacter(ctx context.Context, searchString string) []model.Character_db {
	return s.Repo.SearchCharacter(ctx, searchString)
}
