package pokemonType

import (
	pokemonType "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon-types/domain"
)

type GetByPokemonName struct {
	PokemonTypeRepository pokemonType.PokemonTypeRepository
}

func (getByPokemonName GetByPokemonName) Execute(name string) ([]pokemonType.PokemonType, error) {
	pokemonName, errorOnCreatePokemonName := pokemonType.CreatePokemonName(name)

	if errorOnCreatePokemonName != nil {
		return nil, errorOnCreatePokemonName
	}

	types, errorOnFindingPokemon := getByPokemonName.PokemonTypeRepository.FindByPokemonName(*pokemonName)

	if errorOnFindingPokemon != nil {
		return nil, errorOnFindingPokemon
	}

	return types, nil
}
