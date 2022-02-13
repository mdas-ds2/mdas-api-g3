package webserver

import "errors"

type BadRequestException struct {
	err error
}

func CreateBadRequestException(msg string) BadRequestException {
	exception := BadRequestException{errors.New(msg)}
	return exception
}

func (exception BadRequestException) GetError() error {
	return exception.err
}
