package shared

import "errors"

type PokemonRepositoryUnavailableException struct {
	err error
}

func CreatePokemonRepositoryUnavailableException() PokemonRepositoryUnavailableException {
	exception := PokemonRepositoryUnavailableException{errors.New("pokemon api service unavailable")}
	return exception
}

func (exception PokemonRepositoryUnavailableException) GetError() error {
	return exception.err
}
