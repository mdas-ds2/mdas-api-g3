package domainTest

import (
	"testing"

	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/domain"
)

func TestType(test *testing.T) {
	typeName := "electric"
	pokeTypeName, _ := domain.CreateTypeName(typeName)
	pokeType, _ := domain.CreatePokemonType(*pokeTypeName)
	if typeName != pokeType.GetName().GetValue() {
		test.Errorf("Test incorrect, expected a pokemon type with value %s, received %s", pokeType.GetName().GetValue(), pokeTypeName)
	}
}
