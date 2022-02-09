package pokemonType

import (
	pokemonType "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/domain"
)

type GetByPokemonName struct {
	PokemonTypeRepository pokemonType.PokemonTypeRepository
}

func (getByPokemonName GetByPokemonName) Execute(name string) (pokemonType.PokemonTypes, error) {
	pokemonName, errorOnCreatePokemonName := pokemonType.CreatePokemonName(name)

	if errorOnCreatePokemonName != nil {
		return pokemonType.PokemonTypes{}, errorOnCreatePokemonName
	}

	types, errorOnFindingPokemon := getByPokemonName.PokemonTypeRepository.FindByPokemonName(*pokemonName)

	if errorOnFindingPokemon != nil {
		return pokemonType.PokemonTypes{}, errorOnFindingPokemon
	}

	return types, nil
}
