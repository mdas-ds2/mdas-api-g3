package domainTest

import (
	"testing"

	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/domain"
)

func TestTypeName(test *testing.T) {
	// Given
	typeName := "electric"

	// When
	pokeTypeName, _ := domain.CreateTypeName(typeName)
	result := pokeTypeName.GetValue()

	// Then
	if result != typeName {
		test.Errorf("Test incorrect, expected a pokemon type with value %s, received %s", result, pokeTypeName)
	}
}

func TestTypeNameWithEmptyValue(test *testing.T) {
	// Given
	typeName := ""

	// When
	_, error := domain.CreateTypeName(typeName)

	// Then
	if error == nil {
		test.Errorf("the test should return an error when creating a type name empty")
	}
}
