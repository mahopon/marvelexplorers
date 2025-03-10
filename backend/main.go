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

	DATABASE_URL := os.Getenv("DATABASE_URL")
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

func uploadDB(DATABASE_URL string, characters []Character) {
	conn, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	// use dot operator to access struct members
}

func getHash(ts string, PRIVATE_KEY string, PUBLIC_KEY string) string {
	data := ts + PRIVATE_KEY + PUBLIC_KEY
	hash := md5.Sum([]byte(data))
	hashString := hex.EncodeToString(hash[:])
	return hashString
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
