package pokemonType

import "errors"

type RepositoryUnavailableException struct {
	err error
}

func CreateRepositoryUnavailableException() RepositoryUnavailableException {
	exception := RepositoryUnavailableException{errors.New("pokemon type service unavailable")}
	return exception
}

func (exception RepositoryUnavailableException) GetError() error {
	return exception.err
}
