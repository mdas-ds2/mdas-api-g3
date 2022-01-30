package pokemon

import (
	pokemon "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon/domain"
)

type GetTypesByName struct {
	PokemonRepository pokemon.PokemonRepository
}

func (getTypesByName GetTypesByName) Execute(name pokemon.PokemonName) (string, error) {
	pokemon, error := getTypesByName.PokemonRepository.FindByName(name)

	if error != nil {
		return "", error
	}

	return pokemon.GetStringFormatedTypes(), nil
}
