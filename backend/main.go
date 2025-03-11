package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	PUBLIC_KEY := os.Getenv("PUBLIC_KEY")
	PRIVATE_KEY := os.Getenv("PRIVATE_KEY")
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	hashString := getHash(ts, PRIVATE_KEY, PUBLIC_KEY)

	var offset int64 = 0
	for {
		resp, err := http.Get("https://gateway.marvel.com/v1/public/characters?apikey=" + PUBLIC_KEY + "&hash=" + hashString + "&ts=" + ts + "&offset=" + strconv.FormatInt(offset, 10))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		var result map[string]interface{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			fmt.Println("Error parsing JSON: ", err)
			return
		}

		if data, ok := result["data"].(map[string]interface{}); ok {
			count := int(data["count"].(float64))
			fmt.Println("Hi", count)
			if count == 0 {
				break
			}
			var characters []Character
			if results, ok := data["results"].([]interface{}); ok {
				for _, charData := range results {
					if charMap, ok := charData.(map[string]interface{}); ok {
						jsonTemp, _ := json.Marshal(charMap)
						var character Character
						json.Unmarshal(jsonTemp, &character)
						characters = append(characters, character)
					}
				}
				for _, char := range characters {
					prettyCharacter, err := json.MarshalIndent(char, "", "  ")
					if err != nil {
						fmt.Println("Error marshalling character:", err)
					} else {
						fmt.Println(string(prettyCharacter))
					}
				}
			}
		}
		offset += 20
		time.Sleep(time.Second)
	}
}

func uploadDB(characters []Character) {
	for _, character := range characters {
		insertDB(character)
	}
}

func getHash(ts string, PRIVATE_KEY string, PUBLIC_KEY string) string {
	data := ts + PRIVATE_KEY + PUBLIC_KEY
	hash := md5.Sum([]byte(data))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

func insertDB(character Character) error {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to db %+v\n", err)
	}
	defer conn.Close(ctx)

	batch := &pgx.Batch{}

	// Insert Character
	batch.Queue(
		`INSERT INTO Characters (id, name, description, modified, thumbnail_path, thumbnail_extension, resourceURI) 
         VALUES ($1, $2, $3, $4, $5, $6, $7) 
         ON CONFLICT (id) DO NOTHING`,
		character.ID, character.Name, character.Description, character.Modified,
		character.Thumbnail.Path, character.Thumbnail.Extension, character.ResourceURI,
	)

	// Insert URLs
	for _, url := range character.URLs {
		batch.Queue(
			`INSERT INTO URLs (character_id, type, url) VALUES ($1, $2, $3) 
             ON CONFLICT DO NOTHING`,
			character.ID, url.Type, url.URL,
		)
	}

	// Insert Comics
	for _, comic := range character.Comics.Items {
		batch.Queue(
			`INSERT INTO Comics (title, collectionURI) 
             VALUES ($1, $2) 
             ON CONFLICT (collectionURI) DO NOTHING RETURNING comic_id`,
			comic.Name, character.Comics.CollectionURI,
		)
	}

	// Insert Series
	for _, series := range character.Series.Items {
		batch.Queue(
			`INSERT INTO Series (title, collectionURI) 
             VALUES ($1, $2) 
             ON CONFLICT (collectionURI) DO NOTHING RETURNING series_id`,
			series.Name, character.Series.CollectionURI,
		)
	}

	// Insert Stories
	for _, story := range character.Stories.Items {
		batch.Queue(
			`INSERT INTO Stories (title, type, collectionURI) 
             VALUES ($1, $2, $3) 
             ON CONFLICT (collectionURI) DO NOTHING RETURNING story_id`,
			story.Name, story.Type, character.Stories.CollectionURI,
		)
	}

	// Insert Events
	for _, event := range character.Events.Items {
		batch.Queue(
			`INSERT INTO Events (title, collectionURI) 
             VALUES ($1, $2) 
             ON CONFLICT (collectionURI) DO NOTHING RETURNING event_id`,
			event.Name, character.Events.CollectionURI,
		)
	}

	// Execute batch
	br := conn.SendBatch(ctx, batch)
	defer br.Close()

	// Retrieve IDs for Comics, Series, Stories, and Events
	var comicIDs []int
	for range character.Comics.Items {
		var comicID int
		err := br.QueryRow().Scan(&comicID)
		if err != nil && err != pgx.ErrNoRows {
			return fmt.Errorf("error retrieving comic ID: %w", err)
		}
		comicIDs = append(comicIDs, comicID)
	}

	var seriesIDs []int
	for range character.Series.Items {
		var seriesID int
		err := br.QueryRow().Scan(&seriesID)
		if err != nil && err != pgx.ErrNoRows {
			return fmt.Errorf("error retrieving series ID: %w", err)
		}
		seriesIDs = append(seriesIDs, seriesID)
	}

	var storyIDs []int
	for range character.Stories.Items {
		var storyID int
		err := br.QueryRow().Scan(&storyID)
		if err != nil && err != pgx.ErrNoRows {
			return fmt.Errorf("error retrieving story ID: %w", err)
		}
		storyIDs = append(storyIDs, storyID)
	}

	var eventIDs []int
	for range character.Events.Items {
		var eventID int
		err := br.QueryRow().Scan(&eventID)
		if err != nil && err != pgx.ErrNoRows {
			return fmt.Errorf("error retrieving event ID: %w", err)
		}
		eventIDs = append(eventIDs, eventID)
	}

	// Insert Character_Comics Join Table
	for _, comicID := range comicIDs {
		if comicID != 0 {
			_, err = conn.Exec(ctx,
				`INSERT INTO Character_Comics (character_id, comic_id) VALUES ($1, $2) 
                 ON CONFLICT DO NOTHING`,
				character.ID, comicID,
			)
			if err != nil {
				return fmt.Errorf("error inserting into Character_Comics: %w", err)
			}
		}
	}

	// Insert Character_Series Join Table
	for _, seriesID := range seriesIDs {
		if seriesID != 0 {
			_, err = conn.Exec(ctx,
				`INSERT INTO Character_Series (character_id, series_id) VALUES ($1, $2) 
                 ON CONFLICT DO NOTHING`,
				character.ID, seriesID,
			)
			if err != nil {
				return fmt.Errorf("error inserting into Character_Series: %w", err)
			}
		}
	}

	// Insert Character_Stories Join Table
	for _, storyID := range storyIDs {
		if storyID != 0 {
			_, err = conn.Exec(ctx,
				`INSERT INTO Character_Stories (character_id, story_id) VALUES ($1, $2) 
                 ON CONFLICT DO NOTHING`,
				character.ID, storyID,
			)
			if err != nil {
				return fmt.Errorf("error inserting into Character_Stories: %w", err)
			}
		}
	}

	// Insert Character_Events Join Table
	for _, eventID := range eventIDs {
		if eventID != 0 {
			_, err = conn.Exec(ctx,
				`INSERT INTO Character_Events (character_id, event_id) VALUES ($1, $2) 
                 ON CONFLICT DO NOTHING`,
				character.ID, eventID,
			)
			if err != nil {
				return fmt.Errorf("error inserting into Character_Events: %w", err)
			}
		}
	}

	return nil
}

func getDocs() {

	fmt.Println("Hello world!")
	resp, err := http.Get("https://gateway.marvel.com/docs/public")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print the response
	fmt.Println("Response:", string(body))
}
