package pokeapi

import (
	"encoding/json"

	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/domain"
)

type PokeApiResponse = []byte

func mapResponseToPokemon(response PokeApiResponse) (domain.Pokemon, error) {
	var pokemonResponse PokemonModel
	json.Unmarshal(response, &pokemonResponse)

	pokeId := domain.CreateId(pokemonResponse.ID)
	name := domain.CreateName(pokemonResponse.Name)
	height := domain.CreateHeight(pokemonResponse.Height)
	weight := domain.CreateWeight(pokemonResponse.Weight)

	pokemon := domain.CreatePokemon(pokeId, name, height, weight)

	return pokemon, nil
}
