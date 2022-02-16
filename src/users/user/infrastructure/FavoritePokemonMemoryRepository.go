package user

import (
	domain "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

type FavoritePokemonMemoryRepository struct {
	database map[string][]string
}

func CreateFavoritePokemonMemoryRepository(database *map[string][]string) FavoritePokemonMemoryRepository {
	return FavoritePokemonMemoryRepository{*database}
}

func (repository FavoritePokemonMemoryRepository) Save(user domain.User) error {
	userId := user.GetId()
	for _, favoritePokemonId := range(user)
	if !repository.canBeAdded(userId, favoritePokemonId) {
		exception := domain.CreateFavoritePokemonDuplicatedException(favoritePokemonId)
		return exception.GetError()
	}

	repository.database[userId.GetValue()] = append(repository.database[userId.GetValue()], favoritePokemonId.GetValue())

	return nil
}


func (repository FavoritePokemonMemoryRepository) FindUser(userId domain.UserId) domain.User{
	id := userId.GetValue()
	favorites := repository.database[id]
	result := []domain.PokemonId{}

	for _, favoriteId := range favorites {
		pokemonId := domain.CreatePokemonId(favoriteId)
		result = append(result, pokemonId)
	}

	user := domain.CreateUser
	return result
}
