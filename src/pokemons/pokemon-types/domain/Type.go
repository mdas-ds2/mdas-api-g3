package pokemonType

type PokemonType struct {
	name PokemonTypeName
}

func (pokemonType PokemonType) GetName() PokemonTypeName {
	return pokemonType.name
}

func CreatePokemonType(pokemonTypeName PokemonTypeName) (*PokemonType, error) {
	pokemonType := &PokemonType{
		name: pokemonTypeName,
	}

	return pokemonType, nil
}
