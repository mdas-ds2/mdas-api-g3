package user_test

import (
	"testing"

	domain "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

func TestGetId(t *testing.T) {
	// Given
	id := "1234"
	userId := domain.CreateUserId(id)
	user := domain.CreateUser(userId, domain.PokemonIdCollection{})

	// When
	result := user.GetId().GetValue()

	// Then
	if result != id {
		t.Errorf("Did not get expected result. Wanted user ID %q, got: %q", id, result)
	}
}

func TestAddFavorite(t *testing.T) {
	// Given
	id := "1234"
	userId := domain.CreateUserId(id)
	pokemonId := domain.CreatePokemonId("pika123")
	user := domain.CreateUser(userId, domain.PokemonIdCollection{})

	// When
	user.AddFavorite(pokemonId)
	result := user.GetFavorites().GetValues()[0]

	// Then
	if result != pokemonId {
		t.Errorf("Pokemon has not been inserted correctly")
	}
}

func TestAddFavoriteDuplicated(t *testing.T) {
	// Given
	id := "1234"
	userId := domain.CreateUserId(id)
	pokemonId := domain.CreatePokemonId("pika123")
	collection := domain.CreatePokemonIdCollection([]domain.PokemonId{pokemonId})
	exceptionError := domain.CreateFavoritePokemonDuplicatedException(pokemonId).GetError().Error()
	user := domain.CreateUser(userId, collection)

	// When
	error := user.AddFavorite(pokemonId)
	result := error.Error()

	// Then
	if result != exceptionError {
		t.Errorf("A duplicated pokemon has been inserted")
	}
}
