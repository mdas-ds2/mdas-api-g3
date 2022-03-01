package pokemon_test

import (
	"errors"
	"fmt"
	"testing"

	application "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/application"
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
		domain.CreateTimesAsFavorite(0),
	), nil
}
func (repository PokemonRepositoryMock) Save(pokemon domain.Pokemon) {
	fmt.Print("saved correctly")
}

type PokemonRepositoryFindWithErrorMock struct{}

func (repository PokemonRepositoryFindWithErrorMock) Find(id domain.Id) (domain.Pokemon, error) {
	return domain.Pokemon{}, errors.New("Not found")
}
func (repository PokemonRepositoryFindWithErrorMock) Save(pokemon domain.Pokemon) {
	fmt.Print("saved correctly")
}

func TestGetPokemonDetails(test *testing.T) {
	// Given
	id := POKEMON_ID
	name := POKEMON_NAME
	height := POKEMON_HEIGHT
	weight := POKEMON_WEIGHT
	repository := PokemonRepositoryMock{}

	sut := application.GetPokemonDetails{Repository: repository}

	// When
	pokemon, error := sut.Execute(id)

	// Then
	if error != nil {
		test.Errorf("Error is not expected on this unit test: %s", error.Error())
	}

	if pokemon.Id != id {
		test.Errorf("Expected a returned pokemon detail with id %d, received %d", id, pokemon.Id)
	}

	if pokemon.Name != name {
		test.Errorf("Expected a returned pokemon detail with name %s, received %s", name, pokemon.Name)
	}

	if pokemon.Height != height {
		test.Errorf("Expected a returned pokemon detail with height %d, received %d", height, pokemon.Height)
	}

	if pokemon.Weight != weight {
		test.Errorf("Expected a returned pokemon detail with weight %d, received %d", weight, pokemon.Weight)
	}
}

func TestGetPokemonDetailsFail(test *testing.T) {
	// Given
	id := POKEMON_ID
	repository := PokemonRepositoryFindWithErrorMock{}

	sut := application.GetPokemonDetails{Repository: repository}

	// When
	_, error := sut.Execute(id)

	// Then
	if error == nil {
		test.Errorf("Error is expected on this unit test: %s", error.Error())
	}
}
