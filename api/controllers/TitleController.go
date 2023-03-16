package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"context"
	"fmt"
	"os"

	openai "github.com/openai/openai-go/v2"
)

// GetOfferCode returns the offer code and updates the offers
// with site id
func (server *Server) GetTitles(w http.ResponseWriter, r *http.Request) {
	urlParams := r.URL.Query()

	title := urlParams.Get("title")
	howMany := urlParams.Get("count")

	var titles []string
	var i int

	//convert string to interger
	c, _ := strconv.Atoi(howMany)

	for i = 0; i < c; i++ {
		titles = append(titles, title)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(titles)
}

func connectToChatGPTAndGetTitles(title string, howMany int) []string {

	var alternateTitles []string
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("OPENAI_API_KEY environment variable not set")
		return nil
	}

	ctx := context.Background()
	client, err := openai.NewClient(apiKey)
	if err != nil {
		fmt.Printf("Failed to create client: %v\n", err)
		return nil
	}

	prompt := "The quick brown fox jumps over the lazy dog."
	model := "text-davinci-002"
	params := &openai.CompletionParams{
		Prompt:      prompt,
		MaxTokens:   50,
		Temperature: 0.7,
		N:           1,
		Stop:        []string{"\n"},
	}

	resp, err := client.Completions.Create(ctx, model, params)
	if err != nil {
		fmt.Printf("Failed to generate completions: %v\n", err)
		return nil
	}

	if len(resp.Choices) > 0 {
		alternateTitles = append(alternateTitles, resp.Choices[0].Text)
	} else {
		fmt.Println("No completions returned")
	}

	return alternateTitles
}
