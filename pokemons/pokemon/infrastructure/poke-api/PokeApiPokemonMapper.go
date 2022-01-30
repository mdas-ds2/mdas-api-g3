package pokeApi

import (
	"encoding/json"
	"log"

	pokemon "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon/domain"
)

type PokeApiResponse = []byte

func mapResponseToPokemon(pokemonName pokemon.PokemonName, pokeApiResponse PokeApiResponse) (pokemon.Pokemon, error) {
	var pokemonResponse PokemonModel
	json.Unmarshal(pokeApiResponse, &pokemonResponse)

	var pokemonTypes = []pokemon.PokemonType{}

	for _, pokemonTypeResponse := range pokemonResponse.Types {
		pokemonType, errorOnCreatePokemonType := pokemon.CreatePokemonType(pokemonTypeResponse.Type.Name)

		if errorOnCreatePokemonType != nil {
			log.Fatalln(errorOnCreatePokemonType)
		}

		pokemonTypes = append(pokemonTypes, *pokemonType)
	}

	pokemon, errOnCreatePokemon := pokemon.CreatePokemon(pokemonName, pokemonTypes)

	return *pokemon, errOnCreatePokemon
}
