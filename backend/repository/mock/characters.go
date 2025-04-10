package mock

import (
	"context"
	"tcy/marvelexplorers/model/db"
	"tcy/marvelexplorers/repository"
)

type MockCharacterRepo struct {
	Characters []model.Character_db
}

var _ repository.CharacterRepo = (*MockCharacterRepo)(nil)

func (m *MockCharacterRepo) GetCharacters(ctx context.Context, offset int) interface{} {
	return m.Characters
}

func (m *MockCharacterRepo) SearchCharacter(ctx context.Context, searchString string) []model.Character_db {
	var result []model.Character_db
	for _, c := range m.Characters {
		if searchString == "" || searchString == c.Name {
			result = append(result, c)
		}
	}
	return result
}
