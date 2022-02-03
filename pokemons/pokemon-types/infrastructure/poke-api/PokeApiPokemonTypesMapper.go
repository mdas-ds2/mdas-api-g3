package pokeApi

import (
	"encoding/json"

	pokemonTypes "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon-types/domain"
)

type PokeApiResponse = []byte

func mapResponseToPokemonTypes(pokemonName pokemonTypes.PokemonName, pokeApiResponse PokeApiResponse) ([]pokemonTypes.PokemonType, error) {
	var pokemonResponse PokemonModel
	json.Unmarshal(pokeApiResponse, &pokemonResponse)

	var types = []pokemonTypes.PokemonType{}

	for _, pokemonTypeResponse := range pokemonResponse.Types {
		pokemonTypeName, errorOnCreatePokemonTypeName := pokemonTypes.CreatePokemonTypeName(pokemonTypeResponse.Type.Name)

		if errorOnCreatePokemonTypeName != nil {
			return nil, errorOnCreatePokemonTypeName
		}

		pokemonType, errorOnCreatePokemonType := pokemonTypes.CreatePokemonType(*pokemonTypeName)

		if errorOnCreatePokemonType != nil {
			return nil, errorOnCreatePokemonType
		}

		types = append(types, *pokemonType)
	}

	return types, nil
}
