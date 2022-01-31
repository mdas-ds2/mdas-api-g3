package pokeApi

import (
	generic "github.com/mdas-ds2/mdas-api-g3/generic"
	pokemon "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon/domain"
)

type PokeApiPokemonRepository struct{}

const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon/"

func (pokeApiRepository PokeApiPokemonRepository) FindByName(pokemonName pokemon.PokemonName) (pokemon.Pokemon, error) {
	urlPath := pokeApiUrl + pokemonName.GetValue()
	response, errorOnResponse := generic.HttpGet(urlPath)

	if errorOnResponse != nil {
		return pokemon.Pokemon{}, errorOnResponse
	}

	return mapResponseToPokemon(pokemonName, response)
}
