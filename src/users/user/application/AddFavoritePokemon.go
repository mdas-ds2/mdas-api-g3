package user

import (
	domain "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

type AddFavoritePokemon struct {
	Repository domain.FavoritePokemonRepository
}

func (useCase AddFavoritePokemon) Execute(userId string, pokemonId string) error {
	error := useCase.Repository.Add(
		domain.CreateUserId(userId),
		domain.CreatePokemonId(pokemonId),
	)

	if error != nil {
		return error
	}

	return nil
}
