package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GeminiRequest struct {
	Contents []Content `json:"contents"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected a word, provided 0")
		os.Exit(1)
	}

	word := os.Args[1]

	// hey just add your api key hehe 😊
	apiKey := ""
	if apiKey == "" {
		fmt.Println("GEMINI_API_KEY variable is not set")
		os.Exit(1)
	}

	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash-latest:generateContent?key=" + apiKey

	requestBody := GeminiRequest{
		Contents: []Content{
			{
				Parts: []Part{
					{Text: fmt.Sprintf("Provide a simple definition for the word '%s'. and also make sure to provide some example of it as well, make sure that the response don't contain any bold formatting", word)},
				},
			},
		},
	}

	fmt.Println("Searching the definition for you....... watieuuu~")
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Failed to create request body:", err)
		os.Exit(1)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Failed to make the API request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read the response:", err)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		fmt.Printf("API request failed with status code %d: %s\n", resp.StatusCode, string(body))
		os.Exit(1)
	}

	var geminiResponse GeminiResponse
	if err := json.Unmarshal(body, &geminiResponse); err != nil {
		fmt.Println("Failed to parse the JSON response:", err)
		os.Exit(1)
	}

	if len(geminiResponse.Candidates) > 0 && len(geminiResponse.Candidates[0].Content.Parts) > 0 {
		definition := geminiResponse.Candidates[0].Content.Parts[0].Text
		fmt.Printf("Definition for '%s':\n%s\n", word, definition)
	} else {
		fmt.Println("No definition found in the response.")
	}
}
