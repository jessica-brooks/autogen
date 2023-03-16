package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// get titles from chatGPT
func (server *Server) GetTitles(w http.ResponseWriter, r *http.Request) {
	urlParams := r.URL.Query()
	title := urlParams.Get("title")
	n := urlParams.Get("count")
	brand := urlParams.Get("brand")

	//convert string to interger
	nAlternatives, _ := strconv.Atoi(n)

	var prompt string
	// Write an if statement to check if brand is empty or not
	if brand == "" {
		prompt = fmt.Sprintf("Give me %v alternate text for '%v' with effective SEO", nAlternatives, title)
	} else {
		prompt = fmt.Sprintf("Give me alternate text for '%v' in the style of %v", title, brand)
	}

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

	// alternatives := responseObject["choices"].([]interface{})[0].(map[string]interface{})["text"].(string)

	for _, v := range responseObject["choices"].([]interface{}) {
		alternateTitle := v.(map[string]interface{})["text"].(string)
		alternateTitles = append(alternateTitles, strings.Replace(alternateTitle, "\n\n", "", -1))
	}

	return alternateTitles
}
