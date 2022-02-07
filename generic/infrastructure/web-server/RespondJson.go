package webserver

import "net/http"

func RespondJson(response http.ResponseWriter, responseBody []byte) {
	response.Header().Set("Content-Type", "application/json")

	response.Write(responseBody)
}
