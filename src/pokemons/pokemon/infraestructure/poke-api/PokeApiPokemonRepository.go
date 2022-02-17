package pokeapi

import (
	"net/http"

	httpClient "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/http-client"
	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/domain"
)

type PokeApiPokemonRepository struct{}

const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon/"

func (repository PokeApiPokemonRepository) FindByPokemonName(id domain.Id) (domain.Pokemon, error) {
	urlPath := pokeApiUrl + string(id.GetValue())

	response, errorOnResponse := httpClient.Get(urlPath)

	if response.StatusCode == http.StatusServiceUnavailable {
		serviceUnavailableException := domain.CreateRepositoryUnavailableException()
		return domain.Pokemon{}, serviceUnavailableException.GetError()
	}

	if errorOnResponse != nil {
		return domain.Pokemon{}, errorOnResponse
	}

	if response.StatusCode == http.StatusNotFound {
		pokemonNotFoundException := domain.CreatePokemonNotFoundException(id)
		return domain.Pokemon{}, pokemonNotFoundException.GetError()
	}

	return mapResponseToPokemon(response.Body)
}
