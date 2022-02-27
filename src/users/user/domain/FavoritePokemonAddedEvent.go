package user

type FavoritePokemonAddedEvent struct {
	pokemonId string
}

func (event FavoritePokemonAddedEvent) GetContent() string {
	return event.pokemonId
}

func CreateFavoritePokemonAddedEvent(id PokemonId) FavoritePokemonAddedEvent {
	return FavoritePokemonAddedEvent{id.GetValue()}
}
