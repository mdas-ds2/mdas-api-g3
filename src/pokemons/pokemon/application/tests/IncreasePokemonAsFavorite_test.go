package pokemon_test

import (
	"testing"

	application "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/application"
	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/domain"
)

func TestIncreasePokemonAsFavorite(t *testing.T) {
	//Given
	pokemonId := 25
	repository := PokemonRepositoryMock{}
	useCase := application.IncreasePokemonAsFavorite{repository}

	//When
	useCase.Execute(pokemonId)

	//Then
	pokemon, _ := repository.Find(domain.CreateId(pokemonId))
	result := pokemon.GetTimesAsFavorite().GetValue()
	expected := 1
	if result != expected {
		t.Errorf("result not correct, expected %d, got %d", expected, result)
	}
}
