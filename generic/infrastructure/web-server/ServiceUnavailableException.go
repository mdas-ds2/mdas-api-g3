package webserver

import "errors"

type ServiceUnavailableException struct {
	err error
}

func CreateServiceUnavailableException() ServiceUnavailableException {
	exception := ServiceUnavailableException{errors.New("service unavailable")}
	return exception
}

func (exception ServiceUnavailableException) GetError() error {
	return exception.err
}
