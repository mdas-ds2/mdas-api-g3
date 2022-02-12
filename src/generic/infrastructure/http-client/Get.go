package httpClient

import (
	"io/ioutil"
	"net/http"
)

func Get(url string) ([]byte, int, error) {
	response, errorOnGet := http.Get(url)

	if errorOnGet != nil {
		return nil, http.StatusServiceUnavailable, errorOnGet
	}

	defer response.Body.Close()

	parsedBodyResponse, errorOnParsedBodyResponse := ioutil.ReadAll(response.Body)

	if errorOnParsedBodyResponse != nil {
		return nil, http.StatusInternalServerError, errorOnParsedBodyResponse
	}

	return parsedBodyResponse, response.StatusCode, nil
}
