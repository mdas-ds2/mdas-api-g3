package pokeApi

import (
	"errors"

	http "github.com/mdas-ds2/mdas-api-g3/generic/infrastructure/http"
	pokemonTypes "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon-types/domain"
)

type PokeApiPokemonTypesRepository struct{}

const pokeApiUrl = "http://127.0.0.1/"

func (pokeApiPokemonTypesRepository PokeApiPokemonTypesRepository) FindByPokemonName(pokemonName pokemonTypes.PokemonName) (pokemonTypes.PokemonTypes, error) {
	urlPath := pokeApiUrl + pokemonName.GetValue()
	response, statusCode, errorOnResponse := http.Get(urlPath)

	if statusCode == http.SERVICE_UNAVAILABLE {
		return pokemonTypes.PokemonTypes{}, errors.New("Service unavailable")
	}

	if errorOnResponse != nil {
		return pokemonTypes.PokemonTypes{}, errorOnResponse
	}

	if statusCode == http.NOT_FOUND {
		pokemonNotFoundException := pokemonTypes.CreatePokemonNotFoundException(pokemonName.GetValue())
		return pokemonTypes.PokemonTypes{}, pokemonNotFoundException.GetError()
	}

	return mapResponseToPokemonTypes(pokemonName, response)
}
