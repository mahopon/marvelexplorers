package model

type Character struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Modified    string    `json:"modified"`
	Thumbnail   Thumbnail `json:"thumbnail"`
	ResourceURI string    `json:"resourceURI"`
	Comics      Comics    `json:"comics"`
	Series      Series    `json:"series"`
	Stories     Stories   `json:"stories"`
	Events      Events    `json:"events"`
	URLs        []URL     `json:"urls"`
}

type Thumbnail struct {
	Path      string `json:"path"`
	Extension string `json:"extension"`
}

type Comics struct {
	Available     int    `json:"available"`
	CollectionURI string `json:"collectionURI"`
	Items         []Item `json:"items"`
	Returned      int    `json:"returned"`
}

type Series struct {
	Available     int    `json:"available"`
	CollectionURI string `json:"collectionURI"`
	Items         []Item `json:"items"`
	Returned      int    `json:"returned"`
}

type Stories struct {
	Available     int    `json:"available"`
	CollectionURI string `json:"collectionURI"`
	Items         []Item `json:"items"`
	Returned      int    `json:"returned"`
}

type Events struct {
	Available     int    `json:"available"`
	CollectionURI string `json:"collectionURI"`
	Items         []Item `json:"items"`
	Returned      int    `json:"returned"`
}

type Item struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Type        string `json:"type,omitempty"`
}

type URL struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

// func MapToCharacter(charMap map[string]interface{}) (Character, string) {
// 	thumbnail := Thumbnail{}
// 	comics := Comics{}
// 	series := Series{}
// 	stories := Stories{}
// 	events := Events{}
// 	url := URL{}
// }
