package pokemonType

import "errors"

type TypeName struct {
	value string
}

func (name TypeName) GetValue() string {
	return name.value
}

func CreateTypeName(name string) (*TypeName, error) {
	if len(name) == 0 {
		return nil, errors.New("invalid argument \"type name\": it cannot be an empty string")
	}

	pokemonName := &TypeName{
		value: name,
	}

	return pokemonName, nil
}
