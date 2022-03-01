package user

type User struct {
	id               UserId
	favoritePokemons PokemonIdCollection
	events           []FavoritePokemonAddedEvent
}

func (user User) GetId() UserId {
	return user.id
}

func CreateUser(id UserId, favoritePokemons PokemonIdCollection) *User {
	user := &User{
		id:               id,
		favoritePokemons: favoritePokemons,
		events:           make([]FavoritePokemonAddedEvent, 0),
	}

	return user
}

func (user *User) AddFavorite(pokemonId PokemonId) error {
	if user.favoritePokemons.Has(pokemonId) {
		exception := CreateFavoritePokemonDuplicatedException(pokemonId)
		return exception.GetError()
	}

	user.favoritePokemons.Add(pokemonId)
	event := CreateFavoritePokemonAddedEvent(pokemonId)
	user.raise(event)
	return nil
}

func (user *User) GetFavorites() PokemonIdCollection {
	return user.favoritePokemons
}

func (user *User) raise(event FavoritePokemonAddedEvent) {
	user.events = append(user.events, event)
}

func (user *User) pullDomainEvents() []FavoritePokemonAddedEvent {
	events := user.events
	user.events = nil
	return events
}

func (user *User) GetEvents() []FavoritePokemonAddedEvent {
	return user.events
}
