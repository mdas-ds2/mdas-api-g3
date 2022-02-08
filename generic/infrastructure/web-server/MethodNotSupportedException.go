package webserver

import "errors"

type MethodNotSupportedException struct {
	err error
}

func CreateMethodNotSupportedException() MethodNotSupportedException {
	exception := MethodNotSupportedException{errors.New("method not supported")}
	return exception
}

func (exception MethodNotSupportedException) GetError() error {
	return exception.err
}
