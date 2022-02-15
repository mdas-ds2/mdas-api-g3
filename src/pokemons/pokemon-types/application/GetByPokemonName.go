package domain

import (
	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/domain"
)

type GetByPokemonName struct {
	Repository domain.PokemonTypeRepository
}

func (useCase GetByPokemonName) Execute(name string) (domain.TypeCollection, error) {
	pokemonName, errorOnCreatePokemonName := domain.CreatePokemonName(name)

	if errorOnCreatePokemonName != nil {
		return domain.TypeCollection{}, errorOnCreatePokemonName
	}

	types, errorOnFindingPokemon := useCase.Repository.FindByPokemonName(*pokemonName)

	if errorOnFindingPokemon != nil {
		return domain.TypeCollection{}, errorOnFindingPokemon
	}

	return types, nil
}
