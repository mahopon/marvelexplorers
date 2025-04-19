package redis

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"tcy/marvelexplorers/model/db"
)

func (repo *CharacterRepoRedis) GetCharacters(ctx context.Context, offset int) (any, error) {
	var builder strings.Builder
	builder.WriteString("characters|offset:")
	builder.WriteString(strconv.Itoa(offset))
	result, err := repo.client.Get(builder.String())
	if err != nil {
		log.Printf("Error getting characters: %v", err)
		return nil, err
	}
	return result, nil
}

func (repo *CharacterRepoRedis) SearchCharacter(ctx context.Context, searchString string) ([]model.Character_db, error) {
	var builder strings.Builder
	builder.WriteString("characters:")
	builder.WriteString(searchString)
	result, err := repo.client.Get(builder.String())
	if err != nil {
		log.Printf("Error getting character, %s: %v", searchString, err)
		return nil, err
	}
	var characters []model.Character_db
	err = json.Unmarshal([]byte(result), &characters)
	if err != nil {
		log.Printf("Error unmarshalling character, %s: %v", searchString, err)
		return nil, err
	}
	return characters, nil
}
