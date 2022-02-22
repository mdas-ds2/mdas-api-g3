package user

type PokemonId struct {
	value string
}

func CreatePokemonId(pokeId string) PokemonId {
	return PokemonId{pokeId}
}

func (id PokemonId) GetValue() string {
	return id.value
}
