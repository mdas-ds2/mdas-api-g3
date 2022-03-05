package user

import (
	"bytes"
	"encoding/json"
	"testing"

	"net/http"
)

func TestAddPokemonToFavorites(test *testing.T) {
	//Given
	client := http.Client{}
	postBody, _ := json.Marshal(map[string]string{
		"pokemonId": "25",
	})
	postBuffered := bytes.NewBuffer(postBody)
	req, _ := http.NewRequest("POST", "http://localhost:5001/favorite-pokemon/", postBuffered)
	req.Header.Set("UserId", "pere")

	//When
	result, _ := client.Do(req)

	//Then
	if result.StatusCode != http.StatusOK {
		test.Errorf("the webserver returned an unexpected result: %d", result.StatusCode)
	}
}
