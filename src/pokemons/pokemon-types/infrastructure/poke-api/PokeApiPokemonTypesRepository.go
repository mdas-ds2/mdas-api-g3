package pokeApi

import (
	"net/http"

	httpClient "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/http-client"
	pokemonTypes "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/domain"
)

type PokeApiPokemonTypesRepository struct{}

const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon/"

func (pokeApiPokemonTypesRepository PokeApiPokemonTypesRepository) FindByPokemonName(pokemonName pokemonTypes.PokemonName) (pokemonTypes.PokemonTypes, error) {
	urlPath := pokeApiUrl + pokemonName.GetValue()
	response, statusCode, errorOnResponse := httpClient.Get(urlPath)

	if statusCode == http.StatusServiceUnavailable {
		serviceUnavailableException := pokemonTypes.CreateRepositoryUnavailableException()
		return pokemonTypes.PokemonTypes{}, serviceUnavailableException.GetError()
	}

	if errorOnResponse != nil {
		return pokemonTypes.PokemonTypes{}, errorOnResponse
	}

	if statusCode == http.StatusNotFound {
		pokemonNotFoundException := pokemonTypes.CreatePokemonNotFoundException(pokemonName)
		return pokemonTypes.PokemonTypes{}, pokemonNotFoundException.GetError()
	}

	return mapResponseToPokemonTypes(pokemonName, response)
}
