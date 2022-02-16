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

func (repository *FavoritePokemonMemoryRepository) Save(user domain.User) error {
	userId := user.GetId().GetValue()
	favoriteCollection := user.GetFavorites().GetValues()
	var favoriteIdList []string
	for _, collectionItem := range favoriteCollection {
		favoriteIdList = append(favoriteIdList, collectionItem.GetValue())
	}
	repository.database[userId] = favoriteIdList

	return nil
}

func (repository FavoritePokemonMemoryRepository) FindUser(userId domain.UserId) *domain.User {
	id := userId.GetValue()
	favorites := repository.database[id]
	result := []domain.PokemonId{}

	for _, favoriteId := range favorites {
		pokemonId := domain.CreatePokemonId(favoriteId)
		result = append(result, pokemonId)
	}
	favoriteCollection := domain.CreateFavoritePokemonIdCollection(result)
	user := domain.CreateUser(userId, favoriteCollection)
	return user
}
