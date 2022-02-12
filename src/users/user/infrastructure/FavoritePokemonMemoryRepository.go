package user

import (
	"errors"
	"fmt"

	userDomain "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

type FavoritePokemonMemoryRepository struct {
	database map[string][]string
}

func CreateFavoritePokemonMemoryRepository(database *map[string][]string) FavoritePokemonMemoryRepository {
	return FavoritePokemonMemoryRepository{*database}
}

func (repository FavoritePokemonMemoryRepository) Add(userId userDomain.Id, favoritePokemonId userDomain.FavoritePokemonId) error {
	id := userId.GetValue()
	fmt.Println(repository.canBeAdded(userId, favoritePokemonId))
	if !repository.canBeAdded(userId, favoritePokemonId) {
		return errors.New("Pokemon already in the favotite list")
	}
	repository.database[id] = append(repository.database[id], favoritePokemonId.GetValue())
	return nil
}

func (repository FavoritePokemonMemoryRepository) canBeAdded(userId userDomain.Id, favoritePokemonId userDomain.FavoritePokemonId) bool {
	id := userId.GetValue()
	favorites := repository.database[id]
	fmt.Println(favorites)
	if favorites != nil && len(favorites) == 0 {
		return true
	}
	for _, pokemonId := range favorites {
		if pokemonId == favoritePokemonId.GetValue() {
			return false
		}
	}
	return true
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
