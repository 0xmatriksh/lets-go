package util

import (
	"encoding/json"
	"net/http"
)

func CreateResponse(writer http.ResponseWriter, data interface{}, status int) {
	var jsonResponse []byte
	var err error
	if data != "" {
		jsonResponse, err = json.Marshal(data)
	} else {
		jsonResponse, _ = json.Marshal(map[string]string{
			"message": "No Data Found",
		})
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	if err != nil {
		return
	}
	writer.Write(jsonResponse)

}
