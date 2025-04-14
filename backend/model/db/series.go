package model

import (
	"time"
)

type Series_db struct {
	Id          string
	Description string
	URL         []struct {
		Type string
		Path string
	}
	StartYear int
	EndYear   int
	Modified  time.Time
	Thumbnail struct {
		Path      string
		Extension string
	}
	Creators struct {
		Name string
		Role string
	}
	Characters []Character_db
	Comics     []Comic_db
}
