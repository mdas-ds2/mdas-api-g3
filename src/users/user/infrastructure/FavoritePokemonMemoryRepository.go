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

func (repository FavoritePokemonMemoryRepository) Add(userId domain.UserId, favoritePokemonId domain.PokemonId) error {
	id := userId.GetValue()
	if !repository.canBeAdded(userId, favoritePokemonId) {
		exception := domain.CreateFavoritePokemonDuplicatedException(favoritePokemonId)
		return exception.GetError()
	}
	repository.database[id] = append(repository.database[id], favoritePokemonId.GetValue())
	return nil
}

func (repository FavoritePokemonMemoryRepository) canBeAdded(userId domain.UserId, favoritePokemonId domain.PokemonId) bool {
	id := userId.GetValue()
	favorites := repository.database[id]

	if len(favorites) == 0 {
		return true
	}

	for _, pokemonId := range favorites {
		if pokemonId == favoritePokemonId.GetValue() {
			return false
		}
	}

	return true
}

func (repository FavoritePokemonMemoryRepository) FindAll(userId domain.UserId) []domain.PokemonId {
	id := userId.GetValue()
	favorites := repository.database[id]
	result := []domain.PokemonId{}

	for _, favoriteId := range favorites {
		pokemonId := domain.CreatePokemonId(favoriteId)
		result = append(result, pokemonId)
	}

	return result
}
