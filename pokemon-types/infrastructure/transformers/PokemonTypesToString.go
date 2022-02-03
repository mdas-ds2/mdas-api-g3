package transformers

import (
	pokemonType "github.com/mdas-ds2/mdas-api-g3/pokemon-types/domain"
)

type PokemonTypesToString struct{}

func (pokemonTypesToString PokemonTypesToString) Parse(types pokemonType.PokemonTypes) string {
	var result string
	for index, pType := range types.GetValues() {
		if index > 0 {
			result += ", "
		}
		result += pType.GetName()
	}
	return result
}
