package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

// get titles from chatGPT
func (server *Server) GetTitles(w http.ResponseWriter, r *http.Request) {
	urlParams := r.URL.Query()

	title := urlParams.Get("title")
	howMany := urlParams.Get("count")

	//convert string to interger
	c, _ := strconv.Atoi(howMany)

	alternateTitles := connectToChatGPTAndGetTitles(title, c)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(alternateTitles)
}

func connectToChatGPTAndGetTitles(title string, howMany int) []string {
	var alternateTitles []string
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("OPENAI_API_KEY environment variable not set")
		return nil
	}

	c := openai.NewClient(apiKey)
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:     openai.GPT3Ada,
		MaxTokens: 5,
		Prompt:    title,
		N:         howMany,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return nil
	}

	if len(resp.Choices) > 0 {
		alternateTitles = append(alternateTitles, resp.Choices[0].Text)
	} else {
		fmt.Println("No completions returned")
	}

	return alternateTitles
}
