package user_test

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
	// Given
	pokemonName := "pikachu"
	repository := pokemonApiRepositoryMock{}

	sut := application.GetByPokemonName{
		Repository: repository,
	}

	// When
	result, _ := sut.Execute(pokemonName)
	pokemonType := result.GetValues()[0].GetName().GetValue()

	// Then
	if pokemonType != "electric" {
		test.Errorf("Wrong type for pokemon named %s", pokemonName)
	}
}

func TestGetTypesByPokemonWithEmptyName(test *testing.T) {
	// Given
	pokemonName := ""
	repository := pokemonApiRepositoryMock{}

	sut := application.GetByPokemonName{
		Repository: repository,
	}

	// When
	_, result := sut.Execute(pokemonName)

	// Then
	if result == nil {
		test.Errorf("An error should be returned when pokemon name is empty")
	}
}

func TestGetTypesByPokemonWithNonExistingName(test *testing.T) {
	// Given
	inputPokemonName := "Pere"
	pokemonName, _ := domain.CreatePokemonName(inputPokemonName)
	exceptionError := domain.CreatePokemonNotFoundException(*pokemonName).GetError().Error()
	repository := pokemonApiRepositoryMock{}

	sut := application.GetByPokemonName{
		Repository: repository,
	}

	// When
	_, error := sut.Execute(inputPokemonName)
	result := error.Error()

	// Then
	if result != exceptionError {
		test.Errorf("The error expected is %s but the function returned %s.", exceptionError, result)
	}
}

func TestGetTypesByPokemonWithUnavailableRepo(test *testing.T) {
	// Given
	pokemonName := "pikachu"
	repository := pokemonApiUnavailableMock{}
	exceptionError := domain.CreateRepositoryUnavailableException().GetError().Error()

	sut := application.GetByPokemonName{
		Repository: repository,
	}

	// When
	_, error := sut.Execute(pokemonName)
	result := error.Error()

	// Then
	if result != exceptionError {
		test.Errorf("The error expected is %s but the function returned %s.", exceptionError, result)
	}
}
