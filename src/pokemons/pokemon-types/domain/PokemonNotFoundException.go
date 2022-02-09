package pokemonType

import (
	"errors"
	"fmt"
)

type PokemonNotFoundException struct {
	err error
}

func CreatePokemonNotFoundException(name PokemonName) PokemonNotFoundException {
	message := fmt.Sprintf("Pokemon %s not found", name.GetValue())
	exception := PokemonNotFoundException{errors.New(message)}
	return exception
}

func (exception PokemonNotFoundException) GetError() error {
	return exception.err
}
