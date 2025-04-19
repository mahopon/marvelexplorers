package services

import (
	"context"
	model "tcy/marvelexplorers/model/db"
	repo "tcy/marvelexplorers/repository"
	redis "tcy/marvelexplorers/repository/redis"
)

type CharacterService struct {
	RedisRepo redis.CharacterRepoRedis
	DBRepo    repo.CharacterRepo
}

func NewCharacterService(repo repo.CharacterRepo, redisRepo redis.CharacterRepoRedis) *CharacterService {
	return &CharacterService{DBRepo: repo, RedisRepo: redisRepo}
}

func (s CharacterService) GetCharactersFromDB(ctx context.Context, offset int) (any, error) {
	return s.DBRepo.GetCharacters(ctx, offset)
}

func (s CharacterService) SearchCharacterFromDB(ctx context.Context, searchString string) ([]model.Character_db, error) {
	return s.DBRepo.SearchCharacter(ctx, searchString)
}

func (s CharacterService) GetCharactersFromCache(ctx context.Context, offset int) (any, error) {
	return s.RedisRepo.GetCharacters(ctx, offset)
}

func (s CharacterService) SearchCharacterFromCache(ctx context.Context, searchString string) ([]model.Character_db, any) {
	return s.RedisRepo.SearchCharacter(ctx, searchString)
}

func (s CharacterService) GetCharactersWithCache(ctx context.Context, offset int) (any, error) {
	result, _ := s.GetCharactersFromCache(ctx, offset)
	var err error
	// Cache miss
	if result == "" {
		result, err = s.GetCharactersFromDB(ctx, offset)
		// Add in write to Redis
	}
	return result, err
}

func (s CharacterService) SearchCharacterWithCache(ctx context.Context, searchString string) ([]model.Character_db, error) {
	result, _ := s.SearchCharacterFromCache(ctx, searchString)
	var err error
	// Cache miss
	if result == nil {
		result, err = s.SearchCharacterFromDB(ctx, searchString)
		// Add in write to Redis
	}
	return result, err
}
