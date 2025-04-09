package model

type Comic struct {
	ID                 int                `json:"id"`
	DigitalID          int                `json:"digitalId"`
	Title              string             `json:"title"`
	IssueNumber        int                `json:"issueNumber"`
	VariantDescription string             `json:"variantDescription"`
	Description        string             `json:"description"`
	Modified           string             `json:"modified"`
	ISBN               string             `json:"isbn"`
	UPC                string             `json:"upc"`
	DiamondCode        string             `json:"diamondCode"`
	EAN                string             `json:"ean"`
	ISSN               string             `json:"issn"`
	Format             string             `json:"format"`
	PageCount          int                `json:"pageCount"`
	TextObjects        []ComicsTextObject `json:"textObjects"`
	ResourceURI        string             `json:"resourceURI"`
	URLs               []ComicsURL        `json:"urls"`
	Series             ComicsSeries       `json:"series"`
	Variants           []ComicsVariant    `json:"variants"`
	Collections        []interface{}      `json:"collections"`
	CollectedIssues    []interface{}      `json:"collectedIssues"`
	Dates              []ComicsDate       `json:"dates"`
	Prices             []ComicsPrice      `json:"prices"`
	Thumbnail          ComicsThumbnail    `json:"thumbnail"`
	Images             []ComicsThumbnail  `json:"images"`
	Creators           ComicsCreators     `json:"creators"`
	Characters         ComicsCharacters   `json:"characters"`
	Stories            ComicsStories      `json:"stories"`
	Events             ComicsEvents       `json:"events"`
}

type ComicsTextObject struct {
	Type     string `json:"type"`
	Language string `json:"language"`
	Text     string `json:"text"`
}

type ComicsURL struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type ComicsSeries struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
}

type ComicsVariant struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
}

type ComicsDate struct {
	Type string `json:"type"`
	Date string `json:"date"`
}

type ComicsPrice struct {
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

type ComicsThumbnail struct {
	Path      string `json:"path"`
	Extension string `json:"extension"`
}

type ComicsCreators struct {
	Available     int             `json:"available"`
	CollectionURI string          `json:"collectionURI"`
	Items         []ComicsCreator `json:"items"`
	Returned      int             `json:"returned"`
}

type ComicsCreator struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Role        string `json:"role"`
}

type ComicsCharacters struct {
	Available     int               `json:"available"`
	CollectionURI string            `json:"collectionURI"`
	Items         []ComicsCharacter `json:"items"`
	Returned      int               `json:"returned"`
}

type ComicsCharacter struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
}

type ComicsStories struct {
	Available     int           `json:"available"`
	CollectionURI string        `json:"collectionURI"`
	Items         []ComicsStory `json:"items"`
	Returned      int           `json:"returned"`
}

type ComicsStory struct {
	ResourceURI string `json:"resourceURI"`
	Name        string `json:"name"`
	Type        string `json:"type"`
}

type ComicsEvents struct {
	Available     int           `json:"available"`
	CollectionURI string        `json:"collectionURI"`
	Items         []interface{} `json:"items"`
	Returned      int           `json:"returned"`
}
