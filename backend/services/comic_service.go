package services

import (
	model "tcy/marvelexplorers/model/db"
	repo "tcy/marvelexplorers/repository"
)

type ComicService struct {
	RedisRepo repo.CacheRepo[model.Comic_db]
	DBRepo    repo.DBRepo[model.Comic_db]
}
