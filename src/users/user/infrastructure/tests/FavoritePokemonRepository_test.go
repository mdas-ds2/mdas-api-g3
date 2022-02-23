package user

import (
	"testing"

	domain "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
	infrastructure "github.com/mdas-ds2/mdas-api-g3/src/users/user/infrastructure"
)

func TestUserRepositoryFind(test *testing.T) {
	//Given
	userId := "pere"
	favoritePokemonIds := []string{"25", "678"}
	repository := infrastructure.CreateFavoritePokemonMemoryRepository(&map[string][]string{userId: favoritePokemonIds})

	//When
	result := repository.Find(domain.CreateUserId(userId))

	//Then
	if result.GetId().GetValue() != userId {
		test.Errorf("Did not get expected result. Wanted user ID %q, got: %q", userId, result.GetId().GetValue())
	}
}

func TestUserRepositorySave(test *testing.T) {
	//Given
	userId := domain.CreateUserId("pere")
	pokemonId := domain.CreatePokemonId("25")
	favorites := domain.CreatePokemonIdCollection([]domain.PokemonId{pokemonId})
	user := domain.CreateUser(userId, favorites)
	repository := infrastructure.CreateFavoritePokemonMemoryRepository(&map[string][]string{})

	//When
	repository.Save(*user)

	//Then
	result := repository.Find(userId)
	if result.GetId().GetValue() != userId.GetValue() {
		test.Errorf("Did not get expected result. Wanted user ID %q, got: %q", userId.GetValue(), result.GetId().GetValue())
	}

}
