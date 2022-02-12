package user

import (
	"errors"
	"fmt"
)

type FavoritePokemonDuplicatedException struct {
	err error
}

func CreateFavoritePokemonDuplicatedException(pokemonId FavoritePokemonId) FavoritePokemonDuplicatedException {
	message := fmt.Sprintf("Pokemon with id %s already in favorite list", pokemonId.GetValue())
	exception := FavoritePokemonDuplicatedException{errors.New(message)}
	return exception
}

func (exception FavoritePokemonDuplicatedException) GetError() error {
	return exception.err
}
