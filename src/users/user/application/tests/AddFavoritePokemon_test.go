package user

import (
	"errors"
	"fmt"
	"testing"

	application "github.com/mdas-ds2/mdas-api-g3/src/users/user/application"
	domain "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

const REPEATED_POKEMON_ID = "pkchu9102"
const POKEMON_ID = "chrzd918239"
const USER_ID = "peze12038"

type UserRepositoryMock struct{}

func (repository UserRepositoryMock) Save(user domain.User) error {
	favoritePokemonId := user.GetFavorites().GetValues()[0]
	if favoritePokemonId.GetValue() == REPEATED_POKEMON_ID {
		return domain.CreateFavoritePokemonDuplicatedException(favoritePokemonId).GetError()
	}
	return nil
}

func (repository UserRepositoryMock) Find(userId domain.UserId) domain.User {
	pokemonId := domain.CreatePokemonId("pkchu9102")
	user := domain.CreateUser(domain.CreateUserId(userId.GetValue()), domain.CreatePokemonIdCollection([]domain.PokemonId{pokemonId}))
	return *user
}

type EventPublisherMock struct{}

func (eventPublisher EventPublisherMock) publishEvents(events []domain.FavoritePokemonAddedEvent) error {
	return nil
}

type MockSuper struct{}

func (eventPublisher MockSuper) publishEvents(events []domain.FavoritePokemonAddedEvent) error {
	fmt.Println(events)
	return errors.New("rejkl")
}

func TestAddFavoritePokemon(test *testing.T) {
	// Given
	userId := "peze12038"
	pokemonId := "8239"
	repository := UserRepositoryMock{}

	sut := application.AddFavoritePokemon{Repository: repository, Publisher: application.PublisherMock{}}
	// When
	error := sut.Execute(userId, pokemonId)

	// Then
	if error != nil {
		test.Errorf("Error is not expected on this unit test: %s", error.Error())
	}
}

func TestAddFavoritePokemonDuplicated(test *testing.T) {
	// Given
	userId := "peze12038"
	duplicatedPokemonId := "pkchu9102"
	expectedError := domain.CreateFavoritePokemonDuplicatedException(domain.CreatePokemonId(REPEATED_POKEMON_ID)).GetError().Error()
	repository := UserRepositoryMock{}

	sut := application.AddFavoritePokemon{Repository: repository}

	// When
	error := sut.Execute(userId, duplicatedPokemonId)
	result := error.Error()

	// Then
	if result != expectedError {
		test.Errorf("Error expected is %s but has been received %s", expectedError, result)
	}
}
