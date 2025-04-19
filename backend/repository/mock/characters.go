package mock

import (
	"context"
	"tcy/marvelexplorers/model/db"
)

type MockCharacterRepo struct {
	Characters []model.Character_db
}

func (m *MockCharacterRepo) GetCharacters(ctx context.Context, offset int) (any, error) {
	return m.Characters, nil
}

func (m *MockCharacterRepo) SearchCharacter(ctx context.Context, searchString string) ([]model.Character_db, error) {
	var result []model.Character_db
	for _, c := range m.Characters {
		if searchString == "" || searchString == c.Name {
			result = append(result, c)
		}
	}
	return result, nil
}
