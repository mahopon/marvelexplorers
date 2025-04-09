package model

type Character struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Modified    string             `json:"modified"`
	Thumbnail   CharacterThumbnail `json:"thumbnail"`
	ResourceURI string             `json:"resourceURI"`
	Comics      CharacterComics    `json:"comics"`
	Series      CharacterSeries    `json:"series"`
	Stories     CharacterStories   `json:"stories"`
	Events      CharacterEvents    `json:"events"`
	URLs        []CharacterURL     `json:"urls"`
}

type CharacterThumbnail struct {
	Path      string `json:"path"`
	Extension string `json:"extension"`
}

type CharacterComics struct {
	Available     int                  `json:"available"`
	CollectionURI string               `json:"collectionURI"`
	Items         []CharacterComicItem `json:"items"`
	Returned      int                  `json:"returned"`
}

type CharacterSeries struct {
	Available     int                   `json:"available"`
	CollectionURI string                `json:"collectionURI"`
	Items         []CharacterSeriesItem `json:"items"`
	Returned      int                   `json:"returned"`
}

type CharacterStories struct {
	Available     int                  `json:"available"`
	CollectionURI string               `json:"collectionURI"`
	Items         []CharacterStoryItem `json:"items"`
	Returned      int                  `json:"returned"`
}

type CharacterEvents struct {
	Available     int                  `json:"available"`
	CollectionURI string               `json:"collectionURI"`
	Items         []CharacterEventItem `json:"items"`
	Returned      int                  `json:"returned"`
}

type CharacterComicItem struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Type        string `json:"type,omitempty"`
}

type CharacterSeriesItem struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Type        string `json:"type,omitempty"`
}

type CharacterStoryItem struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Type        string `json:"type,omitempty"`
}

type CharacterEventItem struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Type        string `json:"type,omitempty"`
}

type CharacterURL struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}
