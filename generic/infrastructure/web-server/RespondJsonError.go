package webserver

import (
	"encoding/json"
	"net/http"
)

func RespondJsonError(response http.ResponseWriter, err error) {
	errorMap := make(map[string]string)
	errorMap["error"] = err.Error()
	errorJson, _ := json.Marshal(errorMap)
	response.Write(errorJson)
}
