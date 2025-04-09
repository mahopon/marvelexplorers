package services

import (
	"context"
	model "tcy/marvelexplorers/model/db"
	db "tcy/marvelexplorers/repository"
)

func GetCharacters(ctx context.Context, offset int) interface{} {
	return db.GetPG().GetCharacters(ctx, offset)
}

func SearchCharacter(ctx context.Context, searchString string) []model.Character_db {
	return db.GetPG().SearchCharacter(ctx, searchString)
}
