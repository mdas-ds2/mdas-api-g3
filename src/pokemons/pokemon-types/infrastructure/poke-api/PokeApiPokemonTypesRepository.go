package pokeApi

import (
	http "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/http"
	pokemonTypes "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/domain"
)

type PokeApiPokemonTypesRepository struct{}

const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon/"

func (pokeApiPokemonTypesRepository PokeApiPokemonTypesRepository) FindByPokemonName(pokemonName pokemonTypes.PokemonName) (pokemonTypes.PokemonTypes, error) {
	urlPath := pokeApiUrl + pokemonName.GetValue()
	response, statusCode, errorOnResponse := http.Get(urlPath)

	if statusCode == http.SERVICE_UNAVAILABLE {
		serviceUnavailableException := pokemonTypes.CreateRepositoryUnavailableException()
		return pokemonTypes.PokemonTypes{}, serviceUnavailableException.GetError()
	}

	if errorOnResponse != nil {
		return pokemonTypes.PokemonTypes{}, errorOnResponse
	}

	if statusCode == http.NOT_FOUND {
		pokemonNotFoundException := pokemonTypes.CreatePokemonNotFoundException(pokemonName)
		return pokemonTypes.PokemonTypes{}, pokemonNotFoundException.GetError()
	}

	return mapResponseToPokemonTypes(pokemonName, response)
}
