package transformers

import (
	"encoding/json"

	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/domain"
)

type PokemonTypesToJson struct{}

func (pokemonTypesToJson PokemonTypesToJson) Parse(types domain.TypeCollection) ([]byte, error) {
	var res = []string{}

	for _, value := range types.GetValues() {
		res = append(res, value.GetName().GetValue())
	}

	jsonRes, errorOnCreatingResponse := json.Marshal(res)
	if errorOnCreatingResponse != nil {
		return nil, errorOnCreatingResponse
	}
	return jsonRes, nil
}
