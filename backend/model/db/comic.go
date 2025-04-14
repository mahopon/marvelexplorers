package model

import (
	"time"
)

type Comic_db struct {
	Id          string
	DigitalId   string
	IssueNumber int
	Title       string
	Description string
	Modified    time.Time
	PageCount   int
	ResourceURI string
	Series      struct {
		Name        string
		ResourceURI string
	}
	Thumbnail struct { // Consider making this a struct on its own for reusability
		Path      string
		Extension string
	}
	Writer struct {
		Name string
		Role string
	}
	Characters map[string]Character_db
}
