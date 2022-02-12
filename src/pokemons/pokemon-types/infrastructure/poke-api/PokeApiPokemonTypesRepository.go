package pokeApi

import (
	"net/http"

	httpClient "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/http-client"
	pokemonTypesDomain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/domain"
)

type PokeApiPokemonTypesRepository struct{}

const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon/"

func (pokeApiPokemonTypesRepository PokeApiPokemonTypesRepository) FindByPokemonName(pokemonName pokemonTypesDomain.PokemonName) (pokemonTypesDomain.PokemonTypes, error) {
	urlPath := pokeApiUrl + pokemonName.GetValue()
	response, statusCode, errorOnResponse := httpClient.Get(urlPath)

	if statusCode == http.StatusServiceUnavailable {
		serviceUnavailableException := pokemonTypesDomain.CreateRepositoryUnavailableException()
		return pokemonTypesDomain.PokemonTypes{}, serviceUnavailableException.GetError()
	}

	if errorOnResponse != nil {
		return pokemonTypesDomain.PokemonTypes{}, errorOnResponse
	}

	if statusCode == http.StatusNotFound {
		pokemonNotFoundException := pokemonTypesDomain.CreatePokemonNotFoundException(pokemonName)
		return pokemonTypesDomain.PokemonTypes{}, pokemonNotFoundException.GetError()
	}

	return mapResponseToPokemonTypes(pokemonName, response)
}
