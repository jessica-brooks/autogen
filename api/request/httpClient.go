package request

import (
	//"log"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

//GET sends the GET http request
func GET(apiURL string, requestConfig Config) ([]byte, error) {
	httpClinet := http.Client{Timeout: time.Duration(requestConfig.Timeout.ConnectionTimeout) * time.Second}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	AddRequestHeaders(req, requestConfig.Headers.RequestHeaders)

	resp, err := httpClinet.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}

//POST sends the POST http request
func POST(apiURL string, requestConfig Config, payload map[string]interface{}) ([]byte, error) {
	httpClinet := http.Client{Timeout: time.Duration(requestConfig.Timeout.ConnectionTimeout) * time.Second}

	encodedPayload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(encodedPayload))
	if err != nil {
		return nil, err
	}

	AddRequestHeaders(req, requestConfig.Headers.RequestHeaders)

	resp, err := httpClinet.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}

// AddRequestHeaders add the headers to the request
func AddRequestHeaders(request *http.Request, requestHeaders map[string]string) {
	for headerKey, headerValue := range requestHeaders {
		request.Header.Add(headerKey, headerValue)
	}
}
