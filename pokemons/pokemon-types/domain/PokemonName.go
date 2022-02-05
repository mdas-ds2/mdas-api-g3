package pokemonType

import "errors"

type PokemonName struct {
	value string
}

func (name PokemonName) GetValue() string {
	return name.value
}

func CreatePokemonName(name string) (*PokemonName, error) {
	if len(name) == 0 {
		return nil, errors.New("invalid argument \"name\": it cannot be an empty string")
	}

	pokemonName := &PokemonName{
		value: name,
	}

	return pokemonName, nil
}
