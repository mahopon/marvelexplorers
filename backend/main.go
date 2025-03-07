package main

import (
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

	data := ts + PRIVATE_KEY + PUBLIC_KEY
	hash := md5.Sum([]byte(data))             // Generate MD5 hash
	hashString := hex.EncodeToString(hash[:]) // Convert to hex string

	resp, err := http.Get("https://gateway.marvel.com/v1/public/characters?apikey=" + PUBLIC_KEY + "&hash=" + hashString + "&ts=" + ts)
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

	// Datatype assertion
	if data, ok := result["data"].(map[string]interface{}); ok {
		fmt.Println("hi")
		if results, ok := data["results"].([]interface{}); ok {
			var characters []Character
			for _, charData := range results {
				if charMap, ok := charData.(map[string]interface{}); ok {
					jsonTemp, _ := json.Marshal(charMap)
					var character Character
					json.Unmarshal(jsonTemp, &character)
					characters = append(characters, character)
					// characters = append(characters, charMap) // Beautify the output using json.MarshalIndent
					// prettyCharacter, err := json.MarshalIndent(charMap, "", "  ")
					// if err != nil {
					// 	fmt.Println("Error marshalling character:", err)
					// } else {
					// 	fmt.Println(string(prettyCharacter))
					// }
				}
			}
			fmt.Printf("%+v\n", characters)
		}
	}

	// fmt.Println("Response:", string(body))
}

func getAllCharacters() {

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
