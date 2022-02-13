package domainTest

import (
	"testing"

	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/domain"
)

func TestTypeName(test *testing.T) {
	typeName := "electric"
	pokeTypeName, _ := domain.CreateTypeName(typeName)
	if typeName != pokeTypeName.GetValue() {
		test.Errorf("Test incorrect, expected a pokemon type with value %s, received %s", pokeTypeName.GetValue(), pokeTypeName)
	}
}

func TestTypeNameWithEmptyValue(test *testing.T) {
	typeName := ""
	_, error := domain.CreateTypeName(typeName)
	if error == nil {
		test.Errorf("the test should return an error when creating a type name empty")
	}
}
