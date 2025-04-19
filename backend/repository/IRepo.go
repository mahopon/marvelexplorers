package repository

import (
	"context"
	model "tcy/marvelexplorers/model/db"
)

type CharacterRepo interface {
	GetCharacters(ctx context.Context, offset int) (interface{}, error)
	SearchCharacter(ctx context.Context, searchString string) ([]model.Character_db, error)
}

type CharacterCacheRepo interface {
	GetCharacters(ctx context.Context, offset int) (interface{}, error)
	SearchCharacter(ctx context.Context, searchString string) ([]model.Character_db, error)
}
