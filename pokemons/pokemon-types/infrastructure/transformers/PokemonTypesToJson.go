package transformers

import (
	"encoding/json"

	pokemonType "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon-types/domain"
)

type PokemonTypesToJson struct{}

func (pokemonTypesToJson PokemonTypesToJson) Parse(types pokemonType.PokemonTypes) ([]byte, error) {
	var res = []string{}

	for _, value := range types.GetValues() {
		res = append(res, value.GetName())
	}

	jsonRes, errorOnCreatingResponse := json.Marshal(res)
	if errorOnCreatingResponse != nil {
		return nil, errorOnCreatingResponse
	}
	return jsonRes, nil
}
