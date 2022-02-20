package pokemonType

import (
	"testing"

	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/domain"
	infrastructure "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/infrastructure/poke-api"
)

func TestFindByPokemonName(test *testing.T) {
	//Given
	repository := infrastructure.PokeApiPokemonTypesRepository{}
	pokemonType, _ := domain.CreatePokemonName("pikachu")

	//When
	result, _ := repository.FindByPokemonName(*pokemonType)

	//Then
	expectedType := "electric"
	resultType := result.GetValues()[0].GetName().GetValue()
	if resultType != expectedType {
		test.Errorf("error on getting pokemon type: expected %s, got %s", expectedType, resultType)
	}
}
