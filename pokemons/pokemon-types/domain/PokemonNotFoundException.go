package pokemonType

import "errors"

type PokemonNotFoundException struct {
	err error
}

func CreatePokemonNotFoundException(name string) PokemonNotFoundException {
	exception := PokemonNotFoundException{errors.New("Pokemon " + name + " not found")}
	return exception
}

func (pokemonNotFoundException PokemonNotFoundException) GetError() error {
	return pokemonNotFoundException.err
}
