package console

import "flag"

const defaultInputValue = ""

func NewInputParameter(name, description string) string {
	inputValue := flag.String(name, defaultInputValue, description)
	flag.Parse()
	return *inputValue
}
