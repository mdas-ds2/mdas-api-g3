package pokemon

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	httpClient "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/http-client"
	application "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/application"
)

func GetPokemonFavoriteDetailsTest(test *testing.T) {
	//Given
	userId := "Pere26"
	pokemonId := 26
	addFavoriteUrl := "http://localhost:5001/favorite-pokemon/"
	client := http.Client{}
	postBody, _ := json.Marshal(map[string]string{
		"pokemonId": strconv.Itoa(pokemonId),
	})
	postBuffered := bytes.NewBuffer(postBody)
	req, _ := http.NewRequest("POST", addFavoriteUrl, postBuffered)
	req.Header.Set("UserId", userId)
	client.Do(req)
	getPokemonDetailsUrl := "http://localhost:5001/pokemon/?id=" + strconv.Itoa(pokemonId)

	//When
	httpGetType, _ := httpClient.Get(getPokemonDetailsUrl)

	//Then
	pokemon := application.PokemonDetailsDTO{}
	json.Unmarshal(httpGetType.Body, &pokemon)
	expected := 1
	if pokemon.FavoriteTimesAdded != expected {
		test.Errorf("pokemon added favorite times are not correct, expected: %d, got: %d", expected, pokemon.FavoriteTimesAdded)
	}
}
