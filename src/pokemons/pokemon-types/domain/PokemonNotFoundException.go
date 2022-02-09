package pokemonType

import "errors"

type PokemonNotFoundException struct {
	err error
}

func CreatePokemonNotFoundException(name PokemonName) PokemonNotFoundException {
	exception := PokemonNotFoundException{errors.New("Pokemon " + name.GetValue() + " not found")}
	return exception
}

func (exception PokemonNotFoundException) GetError() error {
	return exception.err
}
