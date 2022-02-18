package pokemon

import "errors"

type PokemonRepositoryUnavailableException struct {
	err error
}

func CreatePokemonRepositoryUnavailableException() PokemonRepositoryUnavailableException {
	exception := PokemonRepositoryUnavailableException{errors.New("pokemon type service unavailable")}
	return exception
}

func (exception PokemonRepositoryUnavailableException) GetError() error {
	return exception.err
}
