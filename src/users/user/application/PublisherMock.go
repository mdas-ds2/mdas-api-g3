package user

import (
	domain "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

type PublisherMock struct {
}

func (eventPublisher PublisherMock) publishEvents(events []domain.FavoritePokemonAddedEvent) error {
	return nil
}
