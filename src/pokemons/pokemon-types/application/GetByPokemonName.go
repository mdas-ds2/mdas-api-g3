package pokemonType

import (
	pokemonType "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/domain"
)

type GetByPokemonName struct {
	Repository pokemonType.PokemonTypeRepository
}

func (useCase GetByPokemonName) Execute(name string) (pokemonType.PokemonTypes, error) {
	pokemonName, errorOnCreatePokemonName := pokemonType.CreatePokemonName(name)

	if errorOnCreatePokemonName != nil {
		return pokemonType.PokemonTypes{}, errorOnCreatePokemonName
	}

	types, errorOnFindingPokemon := useCase.Repository.FindByPokemonName(*pokemonName)

	if errorOnFindingPokemon != nil {
		return pokemonType.PokemonTypes{}, errorOnFindingPokemon
	}

	return types, nil
}
