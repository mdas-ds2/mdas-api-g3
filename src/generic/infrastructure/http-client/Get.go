package httpClient

import (
	"io/ioutil"
	"net/http"
)

type Response struct {
	Body       []byte
	StatusCode int
}

func Get(url string) (Response, error) {
	res, error := http.Get(url)

	if error != nil {
		return Response{nil, http.StatusServiceUnavailable}, error
	}

	defer res.Body.Close()

	parsedBodyResponse, error := ioutil.ReadAll(res.Body)

	if error != nil {
		return Response{nil, http.StatusInternalServerError}, error
	}

	return Response{parsedBodyResponse, res.StatusCode}, nil
}
