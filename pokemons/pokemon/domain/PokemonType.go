package pokemon

import "errors"

type PokemonType struct {
	value string
}

func (pokemonType PokemonType) GetValue() string {
	return pokemonType.value
}

func CreatePokemonType(pType string) (*PokemonType, error) {
	if len(pType) == 0 {
		return nil, errors.New("invalid argument \"type\": it cannot be an empty string")
	}

	pokemonType := &PokemonType{
		value: pType,
	}

	return pokemonType, nil
}
