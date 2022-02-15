package user

import (
	domain "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

type AddFavoritePokemon struct {
	Repository domain.FavoritePokemonRepository
}

func (useCase AddFavoritePokemon) Execute(userId string, pokemonId string) error {
	user := domain.CreateUser(domain.CreateUserId(userId))

	error := useCase.Repository.Add(
		*user,
		domain.CreatePokemonId(pokemonId),
	)

	if error != nil {
		return error
	}

	return nil
}
