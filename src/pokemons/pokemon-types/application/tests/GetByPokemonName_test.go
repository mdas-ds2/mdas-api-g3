package applicationTest

import (
	"testing"

	application "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/application"
	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/domain"
	infraestructure "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/infrastructure/poke-api"
)

func TestGetTypesByPokemonName(test *testing.T) {
	pokemonName := "pikachu"
	pokeApiPokemonTypeRepository := infraestructure.PokeApiPokemonTypesRepository{}

	getByPokemonNameUseCase := application.GetByPokemonName{
		Repository: pokeApiPokemonTypeRepository,
	}

	pokemonTypes, error := getByPokemonNameUseCase.Execute(pokemonName)

	if error != nil {
		test.Errorf("Error is not expected on this unit test: %s", error.Error())
	}

	if pokemonTypes.GetValues()[0].GetName().GetValue() != "electric" {
		test.Errorf("Wrong type for pokemon named %s", pokemonName)

	}
}

func TestGetTypesByPokemonWIthEmptyName(test *testing.T) {
	pokemonName := ""
	pokeApiPokemonTypeRepository := infraestructure.PokeApiPokemonTypesRepository{}

	getByPokemonNameUseCase := application.GetByPokemonName{
		Repository: pokeApiPokemonTypeRepository,
	}

	_, error := getByPokemonNameUseCase.Execute(pokemonName)

	if error == nil {
		test.Errorf("An error should be returned when pokemon name is empty")
	}
}

func TestGetTypesByPokemonWIthNonExistingName(test *testing.T) {
	pokemonName := "pere"
	pokeApiPokemonTypeRepository := infraestructure.PokeApiPokemonTypesRepository{}

	getByPokemonNameUseCase := application.GetByPokemonName{
		Repository: pokeApiPokemonTypeRepository,
	}

	_, error := getByPokemonNameUseCase.Execute(pokemonName)

	if error == nil {
		test.Errorf("An error should be returned when pokemon name is empty")
	}
	pokeName, _ := domain.CreatePokemonName(pokemonName)
	expectedException := domain.CreatePokemonNotFoundException(*pokeName)

	if error.Error() != expectedException.GetError().Error() {
		test.Errorf("The error expected is %s but the function returned %s.", error.Error(), expectedException.GetError().Error())
	}
}
