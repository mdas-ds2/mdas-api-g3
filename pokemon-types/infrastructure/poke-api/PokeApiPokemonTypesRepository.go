package pokeApi

import (
	http "github.com/mdas-ds2/mdas-api-g3/generic/infrastructure/http"
	pokemonTypes "github.com/mdas-ds2/mdas-api-g3/pokemon-types/domain"
)

type PokeApiPokemonTypesRepository struct{}

const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon/"

func (pokeApiPokemonTypesRepository PokeApiPokemonTypesRepository) FindByPokemonName(pokemonName pokemonTypes.PokemonName) ([]pokemonTypes.PokemonType, error) {
	urlPath := pokeApiUrl + pokemonName.GetValue()
	response, errorOnResponse := http.Get(urlPath)

	if errorOnResponse != nil {
		return nil, errorOnResponse
	}

	return mapResponseToPokemonTypes(pokemonName, response)
}
