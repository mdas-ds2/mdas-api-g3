package generic

import (
	"io/ioutil"
	"net/http"
)

const SERVICE_UNAVAILABLE = 503
const NOT_FOUND = 404
const BAD_REQUEST = 400
const INTERNAL_SERVER_ERROR = 500

func Get(url string) ([]byte, int, error) {
	response, errorOnGet := http.Get(url)

	if errorOnGet != nil {
		return nil, SERVICE_UNAVAILABLE, errorOnGet
	}

	defer response.Body.Close()

	parsedBodyResponse, errorOnParsedBodyResponse := ioutil.ReadAll(response.Body)

	if errorOnParsedBodyResponse != nil {
		return nil, INTERNAL_SERVER_ERROR, errorOnParsedBodyResponse
	}

	return parsedBodyResponse, response.StatusCode, nil
}
