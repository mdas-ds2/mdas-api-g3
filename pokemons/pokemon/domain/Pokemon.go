package pokemon

import (
	"errors"
)

type Pokemon struct {
	name  PokemonName
	types []PokemonType
}

func CreatePokemon(pokemonName PokemonName, pokemonTypes []PokemonType) (*Pokemon, error) {
	if len(pokemonName.GetValue()) == 0 {
		return nil, errors.New("invalid argument \"name\": it cannot be an empty string")
	}

	if len(pokemonTypes) == 0 {
		return nil, errors.New("invalid argument \"types\": a pokemon must have at least one type")
	}

	pokemon := &Pokemon{
		name:  pokemonName,
		types: pokemonTypes,
	}

	return pokemon, nil
}

func (pokemon Pokemon) Name() string {
	return pokemon.name.GetValue()
}

func (pokemon Pokemon) GetStringFormatedTypes() string {
	var pokemonTypes string

	for index, pokemonType := range pokemon.types {
		if index > 0 {
			pokemonTypes += ", "
		}
		pokemonTypes += pokemonType.GetValue()
	}

	return pokemonTypes
}
