package user_test

import (
	"testing"

	application "github.com/mdas-ds2/mdas-api-g3/src/users/user/application"
	domain "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

const REPEATED_POKEMON_ID = "pkchu9102"
const POKEMON_ID = "chrzd918239"
const USER_ID = "peze12038"

type FavoritePokemonRepositoryMock struct{}

func (repository FavoritePokemonRepositoryMock) Add(user domain.User, favoritePokemonId domain.PokemonId) error {
	if favoritePokemonId.GetValue() == REPEATED_POKEMON_ID {
		return domain.CreateFavoritePokemonDuplicatedException(favoritePokemonId).GetError()
	}
	return nil
}

func (repository FavoritePokemonRepositoryMock) FindAll(userId domain.UserId) []domain.PokemonId {
	pokemonId := domain.CreatePokemonId(REPEATED_POKEMON_ID)
	pokemonId2 := domain.CreatePokemonId(POKEMON_ID)
	return []domain.PokemonId{pokemonId, pokemonId2}
}

func TestAddFavoritePokemon(test *testing.T) {
	repo := FavoritePokemonRepositoryMock{}
	useCase := application.AddFavoritePokemon{Repository: repo}

	err := useCase.Execute(USER_ID, POKEMON_ID)
	if err != nil {
		test.Errorf("Error is not expected on this unit test: %s", err.Error())
	}

}

func TestAddFavoritePokemonDuplicated(test *testing.T) {
	repo := FavoritePokemonRepositoryMock{}
	useCase := application.AddFavoritePokemon{Repository: repo}

	err := useCase.Execute(USER_ID, REPEATED_POKEMON_ID)
	if err == nil {
		test.Errorf("Error is expected on this unit test")
	}
	expectedError := domain.CreateFavoritePokemonDuplicatedException(domain.CreatePokemonId(REPEATED_POKEMON_ID))
	if err.Error() != expectedError.GetError().Error() {
		test.Errorf("Error expected is %s but has been received %s", expectedError.GetError().Error(), err.Error())
	}
}
