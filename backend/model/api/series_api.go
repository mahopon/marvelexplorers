package model

type Series struct {
	ID          int              `json:"id"`
	Title       string           `json:"title"`
	Description *string          `json:"description"`
	ResourceURI string           `json:"resourceURI"`
	URLs        []SeriesURL      `json:"urls"`
	StartYear   int              `json:"startYear"`
	EndYear     int              `json:"endYear"`
	Rating      string           `json:"rating"`
	Type        string           `json:"type"`
	Modified    string           `json:"modified"`
	Thumbnail   SeriesThumbnail  `json:"thumbnail"`
	Creators    SeriesCreators   `json:"creators"`
	Characters  SeriesCharacters `json:"characters"`
	Stories     SeriesStories    `json:"stories"`
	Comics      SeriesComics     `json:"comics"`
	Events      SeriesEvents     `json:"events"`
	Next        string           `json:"next"`
	Previous    string           `json:"previous"`
}

type SeriesURL struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type SeriesThumbnail struct {
	Path      string `json:"path"`
	Extension string `json:"extension"`
}

type SeriesCreators struct {
	Available     int             `json:"available"`
	CollectionURI string          `json:"collectionURI"`
	Items         []SeriesCreator `json:"items"`
	Returned      int             `json:"returned"`
}

type SeriesCreator struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Role        string `json:"role"`
}

type SeriesCharacters struct {
	Available     int               `json:"available"`
	CollectionURI string            `json:"collectionURI"`
	Items         []SeriesCharacter `json:"items"`
	Returned      int               `json:"returned"`
}

type SeriesCharacter struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
}

type SeriesStories struct {
	Available     int           `json:"available"`
	CollectionURI string        `json:"collectionURI"`
	Items         []SeriesStory `json:"items"`
	Returned      int           `json:"returned"`
}

type SeriesStory struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type SeriesComics struct {
	Available     int           `json:"available"`
	CollectionURI string        `json:"collectionURI"`
	Items         []SeriesComic `json:"items"`
	Returned      int           `json:"returned"`
}

type SeriesComic struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
}

type SeriesEvents struct {
	Available     int           `json:"available"`
	CollectionURI string        `json:"collectionURI"`
	Items         []interface{} `json:"items"`
	Returned      int           `json:"returned"`
}
