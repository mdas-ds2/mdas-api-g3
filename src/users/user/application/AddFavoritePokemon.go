package user

import (
	"fmt"

	userPkg "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

type AddFavoritePokemon struct {
	Repository userPkg.FavoritePokemonRepository
}

func (useCase AddFavoritePokemon) Execute(userId string, pokemonId string) {
	useCase.Repository.Add(
		userPkg.CreateId(userId),
		userPkg.CreatePokemonId(pokemonId),
	)

	fmt.Println(useCase.Repository.FindAll(userPkg.CreateId(userId)))
}
