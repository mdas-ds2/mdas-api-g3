package user

type FavoritePokemonId struct {
	value string
}

func CreatePokemonId(pokeId string) FavoritePokemonId {
	return FavoritePokemonId{pokeId}
}

func (id FavoritePokemonId) GetValue() string {
	return id.value
}
