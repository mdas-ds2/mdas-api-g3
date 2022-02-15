package pokemonType

type TypeCollection struct {
	values []Type
}

func (pokemonTypes TypeCollection) Create() TypeCollection {
	types := []Type{}
	pokeTypes := TypeCollection{types}
	return pokeTypes
}

func (pokemonTypes TypeCollection) GetValues() []Type {
	return pokemonTypes.values
}

func (pokemonTypes *TypeCollection) Add(pokemonType Type) {
	pokemonTypes.values = append(pokemonTypes.values, pokemonType)
}
