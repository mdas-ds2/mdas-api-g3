package pokemonType

type PokemonTypes struct {
	values []PokemonType
}

func (pokemonTypes PokemonTypes) Create() PokemonTypes {
	types := []PokemonType{}
	pokeTypes := PokemonTypes{types}
	return pokeTypes
}

func (pokemonTypes PokemonTypes) GetValues() []PokemonType {
	return pokemonTypes.values
}

func (pokemonTypes *PokemonTypes) Add(pokemonType PokemonType) {
	pokemonTypes.values = append(pokemonTypes.values, pokemonType)
}
