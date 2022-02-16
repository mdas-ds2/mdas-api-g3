package user

import (
	domain "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

type AddFavoritePokemon struct {
	Repository domain.UserRepository
}

func (useCase AddFavoritePokemon) Execute(userId, pokemonId string) error {
	user, err := useCase.Repository.Find(
		domain.CreateUserId(userId),
	)

	if err != nil {
		return err
	}

	user.AddFavorite(
		domain.CreatePokemonId(pokemonId),
	)

	useCase.Repository.Save(user)

	if err != nil {
		return err
	}

	return nil
}
