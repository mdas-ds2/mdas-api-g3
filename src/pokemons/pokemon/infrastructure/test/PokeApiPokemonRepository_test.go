package pokemon

import (
	"testing"

	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/domain"
	infrastructure "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/infrastructure/poke-api"
)

func TestFind(test *testing.T) {
	//Given
	repository := infrastructure.CreatePokeApiPokemonRepository()
	pokemonId := domain.CreateId(25)

	//When
	result, _ := repository.Find(pokemonId)

	//Then
	if result.GetId().GetValue() != pokemonId.GetValue() {
		test.Errorf("Error on getting pokemon info, expected %d, got %d", pokemonId.GetValue(), result.GetId().GetValue())
	}
}
