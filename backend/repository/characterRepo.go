package repository

import (
	"context"
	"tcy/marvelexplorers/model/db"
)

type CharacterRepo interface {
	GetCharacters(ctx context.Context, offset int) interface{}
	SearchCharacter(ctx context.Context, searchString string) []model.Character_db
}
