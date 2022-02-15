package user_test

import (
	"testing"

	domain "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

func TestPokemonId(t *testing.T) {
	originalId := "1234"
	pokemonId := domain.CreatePokemonId(originalId)
	if pokemonId.GetValue() != originalId {
		t.Errorf("Did not get expected result. Wanted user ID %q, got: %q", originalId, pokemonId.GetValue())
	}
}
