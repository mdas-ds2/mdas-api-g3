package pokeApi

import (
	"net/http"

	httpClient "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/http-client"
	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/domain"
	shared "github.com/mdas-ds2/mdas-api-g3/src/shared/infrastructure"
)

type PokeApiPokemonTypesRepository struct{}

const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon/"

func (pokeApiPokemonTypesRepository PokeApiPokemonTypesRepository) FindByPokemonName(pokemonName domain.PokemonName) (domain.TypeCollection, error) {
	urlPath := pokeApiUrl + pokemonName.GetValue()

	response, errorOnResponse := httpClient.Get(urlPath)

	if response.StatusCode == http.StatusServiceUnavailable {
		serviceUnavailableException := shared.CreatePokemonRepositoryUnavailableException()
		return domain.TypeCollection{}, serviceUnavailableException.GetError()
	}

	if errorOnResponse != nil {
		return domain.TypeCollection{}, errorOnResponse
	}

	if response.StatusCode == http.StatusNotFound {
		pokemonNotFoundException := domain.CreatePokemonNotFoundException(pokemonName)
		return domain.TypeCollection{}, pokemonNotFoundException.GetError()
	}

	return mapResponseToPokemonTypes(response.Body)
}
