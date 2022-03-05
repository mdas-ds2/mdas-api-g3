package user

import (
	domain "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

type EventPublisher interface {
	publishEvents(events []domain.FavoritePokemonAddedEvent) error
}
