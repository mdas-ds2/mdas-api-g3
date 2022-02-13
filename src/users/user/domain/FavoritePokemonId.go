// TODO: Este VO nos pasa como con el PokemonName en el PokemonTypes, no es parte de una Entidad
package user

type FavoritePokemonId struct {
	value string
}

func CreatePokemonId(pokeId string) FavoritePokemonId {
	return FavoritePokemonId{pokeId}
}

func (id FavoritePokemonId) GetValue() string {
	return id.value
}
