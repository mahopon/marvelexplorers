package services

import (
	"context"
	"errors"
	"sync"
	model "tcy/marvelexplorers/model/db"
	repo "tcy/marvelexplorers/repository"

	"encoding/json"
	"strconv"
	"time"

	"github.com/mahopon/gobackend/redis"
)

type CharacterService struct {
	RedisRepo repo.CacheRepo[model.Character_db]
	DBRepo    repo.DBRepo[model.Character_db]
}

var (
	character_table          string = "Characters"
	characterServiceInstance *CharacterService
	characterServiceOnce     sync.Once
)

func newCharacterService(repo repo.DBRepo[model.Character_db], redisRepo repo.CacheRepo[model.Character_db]) *CharacterService {
	return &CharacterService{DBRepo: repo, RedisRepo: redisRepo}
}

func GetCharacterService(dbRepo repo.DBRepo[model.Character_db], redisRepo repo.CacheRepo[model.Character_db]) *CharacterService {
	characterServiceOnce.Do(func() {
		characterServiceInstance = newCharacterService(dbRepo, redisRepo)
	})
	return characterServiceInstance
}

func (s CharacterService) GetCharactersFromDB(ctx context.Context, offset int) ([]model.Character_db, error) {
	return s.DBRepo.Get(ctx, character_table, offset)
}

func (s CharacterService) SearchCharacterFromDB(ctx context.Context, searchString string) ([]model.Character_db, error) {
	return s.DBRepo.Search(ctx, character_table, searchString)
}

func (s CharacterService) GetCharactersFromCache(ctx context.Context, offset int) (string, error) {
	cacheKey := "offset:" + strconv.Itoa(offset)
	return s.RedisRepo.Get(ctx, character_table, cacheKey)
}

func (s CharacterService) SearchCharacterFromCache(ctx context.Context, searchString string) (string, error) {
	cacheKey := "search:" + searchString
	return s.RedisRepo.Get(ctx, character_table, cacheKey)
}

func (s CharacterService) InsertCharactersIntoCache(ctx context.Context, key string, value any, ttl time.Duration) error {
	return s.RedisRepo.Insert(ctx, character_table, key, value, ttl)
}

func (s CharacterService) GetCharactersWithCache(ctx context.Context, offset int) ([]byte, error) {
	raw, err := s.GetCharactersFromCache(ctx, offset)
	if err == nil {
		// Cache hit
		if raw != "" {
			return []byte(raw), nil
		}
	}

	// Cache miss - get from DB
	data, err := s.GetCharactersFromDB(ctx, offset)
	if err != nil {
		return nil, err
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Store in cache
	cacheKey := "offset:" + strconv.Itoa(offset)
	_ = s.InsertCharactersIntoCache(ctx, cacheKey, jsonData, redis.CACHE_TTL_LONG)

	return jsonData, nil
}

func (s CharacterService) SearchCharacterWithCache(ctx context.Context, searchString string) ([]byte, error) {
	raw, err := s.SearchCharacterFromCache(ctx, searchString)
	if err == nil {
		// Cache hit
		if raw != "" {
			return []byte(raw), nil
		}
	}

	// Cache miss
	result, err := s.SearchCharacterFromDB(ctx, searchString)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.New("no data")
	}

	// Marshall to JSON
	jsonEncoding, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	// Store in cache
	cacheKey := "search:" + searchString
	_ = s.InsertCharactersIntoCache(ctx, cacheKey, jsonEncoding, redis.CACHE_TTL_LONG)

	return jsonEncoding, nil
}
