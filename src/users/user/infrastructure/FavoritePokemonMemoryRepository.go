package user

import (
	userDomain "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

type FavoritePokemonMemoryRepository struct {
	database map[string][]string
}

func CreateFavoritePokemonMemoryRepository(database *map[string][]string) FavoritePokemonMemoryRepository {
	return FavoritePokemonMemoryRepository{*database}
}

func (repository FavoritePokemonMemoryRepository) Add(userId userDomain.Id, favoritePokemonId userDomain.FavoritePokemonId) {
	id := userId.GetValue()
	repository.database[id] = append(repository.database[id], favoritePokemonId.GetValue())
}

func (repository FavoritePokemonMemoryRepository) FindAll(userId userDomain.Id) []userDomain.FavoritePokemonId {
	id := userId.GetValue()
	favorites := repository.database[id]
	result := []userDomain.FavoritePokemonId{}

	for _, favoriteId := range favorites {
		pokemonId := userDomain.CreatePokemonId(favoriteId)
		result = append(result, pokemonId)
	}

	return result
}
