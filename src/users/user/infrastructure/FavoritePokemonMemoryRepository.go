package infraestructure

import (
	user "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

type FavoritePokemonMemoryRepository struct {
	database map[string][]user.FavoritePokemonId
}

func CreateFavoritePokemonRepository() FavoritePokemonMemoryRepository {
	database := make(map[string][]user.FavoritePokemonId)
	return FavoritePokemonMemoryRepository{database}
}

func (favoritePokemonMemoryRepository FavoritePokemonMemoryRepository) InsertFavoritePokemon(user *user.User, favoritePokemonId user.FavoritePokemonId) {
	userId := user.GetId().GetValue()
	favoritePokemonMemoryRepository.database[userId] = append(favoritePokemonMemoryRepository.database[userId], favoritePokemonId)
}

func (favoritePokemonMemoryRepository FavoritePokemonMemoryRepository) GetPokemonFavorites(user user.User) []user.FavoritePokemonId {
	userId := user.GetId().GetValue()
	return favoritePokemonMemoryRepository.database[userId]
}
