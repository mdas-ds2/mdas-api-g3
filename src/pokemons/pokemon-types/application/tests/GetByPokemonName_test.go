package applicationTest

import (
	"errors"
	"testing"

	application "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/application"
	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/domain"
)

type pokemonApiUnavailableMock struct{}

func (pokeApiPokemonTypesRepository pokemonApiUnavailableMock) FindByPokemonName(pokemonName domain.PokemonName) (domain.PokemonTypes, error) {
	return domain.PokemonTypes{}, domain.CreateRepositoryUnavailableException().GetError()
}

type pokemonApiRepositoryMock struct{}

func (pokeApiPokemonTypesRepository pokemonApiRepositoryMock) FindByPokemonName(pokemonName domain.PokemonName) (domain.PokemonTypes, error) {
	if pokemonName.GetValue() == "" {
		return domain.PokemonTypes{}, errors.New("invalid argument \"type name\": it cannot be an empty string")
	}
	if pokemonName.GetValue() != "pikachu" {
		return domain.PokemonTypes{}, domain.CreatePokemonNotFoundException(pokemonName).GetError()
	}
	typeName, _ := domain.CreateTypeName("electric")
	pokeType, _ := domain.CreatePokemonType(*typeName)
	pokemonTypes := (domain.PokemonTypes{}).Create()
	pokemonTypes.Add(*pokeType)
	return pokemonTypes, nil
}

func TestGetTypesByPokemonName(test *testing.T) {
	pokemonName := "pikachu"
	pokeApiPokemonTypeRepository := pokemonApiRepositoryMock{}

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
	pokeApiPokemonTypeRepository := pokemonApiRepositoryMock{}

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
	pokeApiPokemonTypeRepository := pokemonApiRepositoryMock{}

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

func TestGetTypesByPokemonWithUnavailableRepo(test *testing.T) {
	pokemonName := "pikachu"
	pokeApiPokemonTypeRepository := pokemonApiUnavailableMock{}
	getByPokemonNameUseCase := application.GetByPokemonName{
		Repository: pokeApiPokemonTypeRepository,
	}

	_, error := getByPokemonNameUseCase.Execute(pokemonName)

	if error == nil {
		test.Errorf("An error should be returned when pokemon name is empty")
	}
	expectedException := domain.CreateRepositoryUnavailableException()

	if error.Error() != expectedException.GetError().Error() {
		test.Errorf("The error expected is %s but the function returned %s.", error.Error(), expectedException.GetError().Error())
	}
}
