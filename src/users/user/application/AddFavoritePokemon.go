package user

import (
	domain "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

type AddFavoritePokemon struct {
	Repository domain.UserRepository
	Publisher  EventPublisher
}

func (useCase AddFavoritePokemon) Execute(userId, pokemonId string) error {
	user := useCase.Repository.Find(
		domain.CreateUserId(userId),
	)

	err := user.AddFavorite(
		domain.CreatePokemonId(pokemonId),
	)

	if err != nil {
		return err
	}

	useCase.Repository.Save(user)

	events := user.GetEvents()

	err = useCase.Publisher.publishEvents(events)

	if err != nil {
		return err
	}

	return nil
}
