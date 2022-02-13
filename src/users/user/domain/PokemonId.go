// TODO: Este VO nos pasa como con el PokemonName en el PokemonTypes, no es parte de una Entidad
package user

type PokemonId struct {
	value string
}

func CreatePokemonId(pokeId string) PokemonId {
	return PokemonId{pokeId}
}

func (id PokemonId) GetValue() string {
	return id.value
}
