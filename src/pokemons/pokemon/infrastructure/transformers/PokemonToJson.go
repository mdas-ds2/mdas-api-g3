package transformers

import (
	"encoding/json"

	application "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/application"
)

type PokemonToJson struct{}

func (pokemonToJson PokemonToJson) Parse(pokemon application.PokemonDetails) ([]byte, error) {
	jsonRes, errorOnCreatingResponse := json.Marshal(pokemon)
	if errorOnCreatingResponse != nil {
		return nil, errorOnCreatingResponse
	}
	return jsonRes, nil
}
