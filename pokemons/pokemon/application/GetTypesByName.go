package pokemon

import (
	pokemon "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon/domain"
)

type GetTypesByName struct {
	PokemonRepository pokemon.PokemonRepository
}

func (getTypesByName GetTypesByName) Execute(name string) (string, error) {
	pokemonName, errorOnCreatePokemonName := pokemon.CreatePokemonName(name)

	if errorOnCreatePokemonName != nil {
		return "", errorOnCreatePokemonName
	}

	pokemon, errorOnFindingPokemon := getTypesByName.PokemonRepository.FindByName(*pokemonName)

	if errorOnFindingPokemon != nil {
		return "", errorOnFindingPokemon
	}

	return pokemon.GetStringFormatedTypes(), nil
}
