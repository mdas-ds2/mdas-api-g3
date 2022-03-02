package pokemon

import (
	"encoding/json"

	"testing"

	httpClient "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/http-client"
	application "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/application"
)

func TestGetPokemonType(test *testing.T) {
	//Given
	url := "http://localhost:5001/pokemon/?id=25"

	//When
	httpGetType, _ := httpClient.Get(url)

	//Then
	pokemon := application.PokemonDetailsDTO{}
	json.Unmarshal(httpGetType.Body, &pokemon)
	if pokemon.Name != "pikachu" {
		test.Errorf("Unexpected result getting type, expected %s received %s", "pikachu", pokemon.Name)
	}
}
