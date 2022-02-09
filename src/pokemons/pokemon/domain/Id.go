package pokemon

type Id struct {
	value string
}

func CreatePokemonId(pokeId string) Id {
	return Id{pokeId}
}
