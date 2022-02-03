package pokeApi

import (
	"encoding/json"

	pokemonTypes "github.com/mdas-ds2/mdas-api-g3/pokemon-types/domain"
)

type PokeApiResponse = []byte

func mapResponseToPokemonTypes(pokemonName pokemonTypes.PokemonName, pokeApiResponse PokeApiResponse) (pokemonTypes.PokemonTypes, error) {
	var pokemonResponse PokemonModel
	json.Unmarshal(pokeApiResponse, &pokemonResponse)

	var types = pokemonTypes.PokemonTypes{}

	for _, pokemonTypeResponse := range pokemonResponse.Types {
		pokemonTypeName, errorOnCreatePokemonTypeName := pokemonTypes.CreatePokemonTypeName(pokemonTypeResponse.Type.Name)

		if errorOnCreatePokemonTypeName != nil {
			return pokemonTypes.PokemonTypes{}, errorOnCreatePokemonTypeName
		}

		pokemonType, errorOnCreatePokemonType := pokemonTypes.CreatePokemonType(*pokemonTypeName)

		if errorOnCreatePokemonType != nil {
			return pokemonTypes.PokemonTypes{}, errorOnCreatePokemonType
		}

		types.Add(*pokemonType)
	}

	return types, nil
}
