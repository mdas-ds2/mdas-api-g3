package pokeApi

import (
	generic "github.com/mdas-ds2/mdas-api-g3/generic"
	pokemon "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon/domain"
)

type PokeApiPokemonRepository struct{}

const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon/"

func (pokeApiRepository PokeApiPokemonRepository) FindByName(pokemonName pokemon.PokemonName) (pokemon.Pokemon, error) {
	responseBody, errorGetHttpPokeApi := generic.HttpGet(pokeApiUrl + pokemonName.GetValue())

	if errorGetHttpPokeApi != nil {
		return pokemon.Pokemon{}, errorGetHttpPokeApi
	}

	return mapResponseToPokemon(pokemonName, responseBody)
}
