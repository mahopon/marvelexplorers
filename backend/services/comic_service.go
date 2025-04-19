package services

import (
	repo "tcy/marvelexplorers/repository"
)

type ComicService struct {
	RedisRepo repo.CharacterCacheRepo
	DBRepo    repo.CharacterRepo
}
