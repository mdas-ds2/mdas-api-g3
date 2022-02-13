package user

import (
	"fmt"

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

	// TODO: Helper log; remove this for production
	fmt.Println(useCase.Repository.FindAll(domain.CreateUserId(userId)))

	return nil
}
