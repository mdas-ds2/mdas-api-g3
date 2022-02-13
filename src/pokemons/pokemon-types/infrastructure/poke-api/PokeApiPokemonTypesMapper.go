package pokeApi

import (
	"encoding/json"

	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/domain"
)

type PokeApiResponse = []byte

func mapResponseToPokemonTypes(pokeApiResponse PokeApiResponse) (domain.PokemonTypes, error) {
	var pokemonResponse PokemonModel
	json.Unmarshal(pokeApiResponse, &pokemonResponse)

	var pTypes = domain.PokemonTypes{}
	types := pTypes.Create()

	for _, pokemonTypeResponse := range pokemonResponse.Types {
		pokemonTypeName, errorOnCreatePokemonTypeName := domain.CreateTypeName(pokemonTypeResponse.Type.Name)

		if errorOnCreatePokemonTypeName != nil {
			return domain.PokemonTypes{}, errorOnCreatePokemonTypeName
		}

		pokemonType, errorOnCreatePokemonType := domain.CreatePokemonType(*pokemonTypeName)

		if errorOnCreatePokemonType != nil {
			return domain.PokemonTypes{}, errorOnCreatePokemonType
		}

		types.Add(*pokemonType)
	}

	return types, nil
}
