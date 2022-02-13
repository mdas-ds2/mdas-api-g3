package pokemonType

type PokemonType struct {
	name TypeName
}

func (pokemonType PokemonType) GetName() TypeName {
	return pokemonType.name
}

func CreatePokemonType(pokemonTypeName TypeName) (*PokemonType, error) {
	pokemonType := &PokemonType{
		name: pokemonTypeName,
	}

	return pokemonType, nil
}
