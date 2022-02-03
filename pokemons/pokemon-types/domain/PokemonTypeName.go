package pokemonType

import "errors"

type PokemonTypeName struct {
	value string
}

func (name PokemonTypeName) GetValue() string {
	return name.value
}

func CreatePokemonTypeName(name string) (*PokemonTypeName, error) {
	if len(name) == 0 {
		return nil, errors.New("invalid argument \"type name\": it cannot be an empty string")
	}

	pokemonName := &PokemonTypeName{
		value: name,
	}

	return pokemonName, nil
}
