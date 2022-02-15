package pokemonType

type Type struct {
	name TypeName
}

func (pokemonType Type) GetName() TypeName {
	return pokemonType.name
}

func CreateType(pokemonTypeName TypeName) (*Type, error) {
	pokemonType := &Type{
		name: pokemonTypeName,
	}

	return pokemonType, nil
}
