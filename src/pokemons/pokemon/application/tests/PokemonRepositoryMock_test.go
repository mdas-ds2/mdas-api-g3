package pokemon_test

import (
	"fmt"

	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/domain"
)

const POKEMON_ID = 1234
const POKEMON_NAME = "Pikachu"
const POKEMON_HEIGHT = 1
const POKEMON_WEIGHT = 1

type PokemonRepositoryMock struct{}

func (repository PokemonRepositoryMock) Find(id domain.Id) (domain.Pokemon, error) {
	return domain.CreatePokemon(
		id,
		domain.CreateName(POKEMON_NAME),
		domain.CreateHeight(POKEMON_HEIGHT),
		domain.CreateWeight(POKEMON_WEIGHT),
		domain.CreateTimesAsFavorite(1),
	), nil
}
func (repository PokemonRepositoryMock) Save(pokemon domain.Pokemon) {
	fmt.Print("saved correctly")
}
