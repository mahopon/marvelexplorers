package model

import "time"

type Character_db struct {
	ID                 int       `db:"id"`
	Name               string    `db:"name"`
	Description        string    `db:"description"`
	Modified           time.Time `db:"modified"` // Change to time.Time if this is a date/timestamp
	ThumbnailPath      string    `db:"thumbnail_path"`
	ThumbnailExtension string    `db:"thumbnail_extension"`
	ResourceURI        string    `db:"resourceuri"`
}
