package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//JSON json encodes the response
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

//ERROR returns json error response
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			StatusCode int    `json:"statusCode"`
			Error      string `json:"error"`
		}{
			StatusCode: statusCode,
			Error:      err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}

//SUCCESS returns json success response
func SUCCESS(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusOK, struct {
		StatusCode int         `json:"statusCode"`
		Data       interface{} `json:"data"`
	}{
		StatusCode: http.StatusOK,
		Data:       data,
	})
}
