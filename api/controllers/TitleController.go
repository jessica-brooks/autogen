package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// get titles from chatGPT
func (server *Server) GetTitles(w http.ResponseWriter, r *http.Request) {
	urlParams := r.URL.Query()
	title := urlParams.Get("title")
	brand := urlParams.Get("brand")
	n := urlParams.Get("count")

	//convert string to interger
	nAlternatives, _ := strconv.Atoi(n)
	//"Give me 3 alternate text for 'Enjoy 25% off orders in the sale when using this ASOS voucher code'"
	prompt := fmt.Sprintf("Give me %v alternate text for '%v' in style of %v", nAlternatives, title, brand)
	alternateTitles := connectToChatGPTAndGetTitles(prompt, nAlternatives)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alternateTitles)
}

func connectToChatGPTAndGetTitles(prompt string, nAlternatives int) []string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	var alternateTitles []string
	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		fmt.Println("OPENAI_API_KEY environment variable not set")
		return nil
	}

	var modelEngine = "text-davinci-002"
	var temperature = 0.7
	var maxTokens = 60

	// Generate the request body
	requestBody, err := json.Marshal(map[string]interface{}{
		"prompt":      prompt,
		"temperature": temperature,
		"max_tokens":  maxTokens,
		"n":           nAlternatives,
	})
	if err != nil {
		panic(err)
	}

	// Generate the API request
	request, err := http.NewRequest("POST", "https://api.openai.com/v1/engines/"+modelEngine+"/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+apiKey)

	// Send the API request and parse the response
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var responseObject map[string]interface{}
	json.Unmarshal(responseBody, &responseObject)

	for _, choice := range responseObject["choices"].([]interface{}) {
		alternateTitles = append(alternateTitles, choice.(map[string]interface{})["text"].(string))
	}

	return alternateTitles
}
