package pokemonType

import (
	httpClient "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/http-client"

	"testing"
)

func TestGetPokemonType(test *testing.T) {
	//Given
	url := "http://localhost:5001/pokemon-types?name=pikachu"

	//When
	httpGetType, _ := httpClient.Get(url)

	//Then
	result := string(httpGetType.Body)
	if result != "[\"electric\"]" {
		test.Errorf("Unexpected result getting the type of pikachu")
	}
}
