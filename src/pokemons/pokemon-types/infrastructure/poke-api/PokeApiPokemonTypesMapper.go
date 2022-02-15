package pokeApi

import (
	"encoding/json"

	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/domain"
)

type PokeApiResponse = []byte

func mapResponseToPokemonTypes(pokeApiResponse PokeApiResponse) (domain.TypeCollection, error) {
	var pokemonResponse PokemonModel
	json.Unmarshal(pokeApiResponse, &pokemonResponse)

	var pTypes = domain.TypeCollection{}
	types := pTypes.Create()

	for _, pokemonTypeResponse := range pokemonResponse.Types {
		pokemonTypeName, errorOnCreatePokemonTypeName := domain.CreateTypeName(pokemonTypeResponse.Type.Name)

		if errorOnCreatePokemonTypeName != nil {
			return domain.TypeCollection{}, errorOnCreatePokemonTypeName
		}

		pokemonType, errorOnCreatePokemonType := domain.CreateType(*pokemonTypeName)

		if errorOnCreatePokemonType != nil {
			return domain.TypeCollection{}, errorOnCreatePokemonType
		}

		types.Add(*pokemonType)
	}

	return types, nil
}
