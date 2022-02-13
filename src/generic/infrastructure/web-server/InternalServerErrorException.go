package webserver

import "errors"

type InternalServerErrorException struct {
	err error
}

func CreateInternalServerErrorException(msg string) InternalServerErrorException {
	exception := InternalServerErrorException{errors.New(msg)}
	return exception
}

func (exception InternalServerErrorException) GetError() error {
	return exception.err
}
