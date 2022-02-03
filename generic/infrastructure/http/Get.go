package generic

import (
	"io/ioutil"
	"net/http"
)

func Get(url string) ([]byte, error) {
	response, errorOnResponse := http.Get(url)

	if errorOnResponse != nil {
		return nil, errorOnResponse
	}

	defer response.Body.Close()

	parsedBodyResponse, errorOnParsedBodyResponse := ioutil.ReadAll(response.Body)

	if errorOnParsedBodyResponse != nil {
		return nil, errorOnParsedBodyResponse
	}

	return parsedBodyResponse, nil
}
