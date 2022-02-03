package pokemonType

type PokemonTypes struct {
	values []PokemonType
}

func (pokemonTypes PokemonTypes) GetValues() []PokemonType {
	return pokemonTypes.values
}

func (pokemonTypes PokemonTypes) Add(pokemonType PokemonType) {
	pokemonTypes.values = append(pokemonTypes.values, pokemonType)
}
