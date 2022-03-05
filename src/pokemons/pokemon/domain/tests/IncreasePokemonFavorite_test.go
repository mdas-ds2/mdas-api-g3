package pokemon_test

import (
	"testing"

	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/domain"
)

func TestIncreasePokemonFavoriteTimes(t *testing.T) {
	//Given
	pokemon := domain.CreatePokemon(
		domain.CreateId(46),
		domain.CreateName("charizard"),
		domain.CreateHeight(30),
		domain.CreateWeight(35),
		domain.CreateTimesAsFavorite(0),
	)

	//When
	pokemon.IncreaseFavoriteTimes()

	//Then
	expected := 1
	result := pokemon.GetTimesAsFavorite().GetValue()

	if expected != result {
		t.Errorf("Favorite times hasn´t been increased properly, expected: %d, got: %d", expected, result)
	}
}
