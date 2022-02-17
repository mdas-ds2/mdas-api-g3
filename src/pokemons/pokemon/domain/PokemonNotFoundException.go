package pokemon

import (
	"errors"
	"fmt"
)

type PokemonNotFoundException struct {
	err error
}

func CreatePokemonNotFoundException(id Id) PokemonNotFoundException {
	message := fmt.Sprintf("Pokemon %d not found", id.GetValue())
	exception := PokemonNotFoundException{errors.New(message)}
	return exception
}

func (exception PokemonNotFoundException) GetError() error {
	return exception.err
}
