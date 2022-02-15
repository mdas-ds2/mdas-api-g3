package domainTest

import (
	"testing"

	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/domain"
)

func TestType(test *testing.T) {
	// Given
	typeName := "electric"
	pokeTypeName, _ := domain.CreateTypeName(typeName)

	// When
	pokeType, _ := domain.CreateType(*pokeTypeName)
	result := pokeType.GetName().GetValue()

	// Then
	if result != typeName {
		test.Errorf("Test incorrect, expected a pokemon type with value %s, received %s", pokeType.GetName().GetValue(), pokeTypeName)
	}
}
