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
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func extract() {
	err := godotenv.Load("../.env")
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
				go uploadDB(characters)
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
	// Set up context
	ctx := context.Background()

	// Connect to the connection pool (create once at the start of your application)
	connPool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to the database: %+v\n", err)
	}
	defer connPool.Close() // Close the pool when done

	// Use the pool to create a batch request
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
			 ON CONFLICT (title) DO NOTHING RETURNING comic_id`,
			comic.Name, comic.ResourceURI,
		)
	}

	// Insert Series
	for _, series := range character.Series.Items {
		batch.Queue(
			`INSERT INTO Series (title, collectionURI) 
			 VALUES ($1, $2) 
			 ON CONFLICT (title) DO NOTHING RETURNING series_id`,
			series.Name, series.ResourceURI,
		)
	}

	// Insert Stories
	for _, story := range character.Stories.Items {
		batch.Queue(
			`INSERT INTO Stories (title, type, collectionURI) 
			 VALUES ($1, $2, $3) 
			 ON CONFLICT (title) DO NOTHING RETURNING story_id`,
			story.Name, story.Type, story.ResourceURI,
		)
	}

	// Insert Events
	for _, event := range character.Events.Items {
		batch.Queue(
			`INSERT INTO Events (title, collectionURI) 
			 VALUES ($1, $2) 
			 ON CONFLICT (title) DO NOTHING RETURNING event_id`,
			event.Name, event.ResourceURI,
		)
	}

	// Execute batch
	br := connPool.SendBatch(ctx, batch)
	defer br.Close()

	// Retrieve IDs for Comics, Series, Stories, and Events
	var comicIDs []int
	for _, comic := range character.Comics.Items {
		var comicID int
		err := connPool.QueryRow(ctx, "SELECT comic_id FROM Comics WHERE title = $1", comic.Name).Scan(&comicID)
		if err != nil && err != pgx.ErrNoRows {
			return fmt.Errorf("error retrieving comic ID: %w", err)
		} else if err == pgx.ErrNoRows {
			return fmt.Errorf("no rows")
		}
		comicIDs = append(comicIDs, comicID)
	}

	var seriesIDs []int
	for _, series := range character.Series.Items {
		var seriesID int
		err := connPool.QueryRow(ctx, "SELECT series_id FROM Series WHERE title = $1", series.Name).Scan(&seriesID)
		if err != nil && err != pgx.ErrNoRows {
			return fmt.Errorf("error retrieving series ID: %w", err)
		} else if err == pgx.ErrNoRows {
			return fmt.Errorf("no rows")
		}
		seriesIDs = append(seriesIDs, seriesID)
	}

	var storyIDs []int
	for _, story := range character.Stories.Items {
		var storyID int
		err := connPool.QueryRow(ctx, "SELECT story_id FROM Stories WHERE title = $1", story.Name).Scan(&storyID)
		if err != nil && err != pgx.ErrNoRows {
			return fmt.Errorf("error retrieving story ID: %w", err)
		} else if err == pgx.ErrNoRows {
			return fmt.Errorf("no rows")
		}
		storyIDs = append(storyIDs, storyID)
	}

	var eventIDs []int
	for _, event := range character.Events.Items {
		var eventID int
		err := connPool.QueryRow(ctx, "SELECT event_id FROM Events WHERE title = $1", event.Name).Scan(&eventID)
		if err != nil && err != pgx.ErrNoRows {
			return fmt.Errorf("error retrieving event ID: %w", err)
		} else if err == pgx.ErrNoRows {
			return fmt.Errorf("no rows")
		}
		eventIDs = append(eventIDs, eventID)
	}

	// Insert Character_Comics Join Table
	for _, comicID := range comicIDs {
		if comicID != 0 {
			_, err = connPool.Exec(ctx,
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
			_, err = connPool.Exec(ctx,
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
			_, err = connPool.Exec(ctx,
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
			_, err = connPool.Exec(ctx,
				`INSERT INTO Character_Events (character_id, event_id) VALUES ($1, $2) 
                 ON CONFLICT DO NOTHING`,
				character.ID, eventID,
			)
			if err != nil {
				return fmt.Errorf("error inserting into Character_Events: %w", err)
			}
		}
	}
	fmt.Println("Done: ", character.Name)
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
