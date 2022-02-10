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
	favoritePokemonMemoryRepository.database[user.GetId().GetValue()] = append(favoritePokemonMemoryRepository.database[user.GetId().GetValue()], favoritePokemonId)
}

func (favoritePokemonMemoryRepository FavoritePokemonMemoryRepository) GetPokemonFavotites(user user.User) []user.FavoritePokemonId {
	return favoritePokemonMemoryRepository.database[user.GetId().GetValue()]
}
