package pokemonType

type PokemonType struct {
	name PokemonTypeName
}

func (pokemonType PokemonType) GetName() string {
	return pokemonType.name.GetValue()
}

func CreatePokemonType(pokemonTypeName PokemonTypeName) (*PokemonType, error) {
	pokemonType := &PokemonType{
		name: pokemonTypeName,
	}

	return pokemonType, nil
}
