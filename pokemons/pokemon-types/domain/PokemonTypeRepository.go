package pokemonType

type PokemonTypeRepository interface {
	FindByPokemonName(pokemonName PokemonName) (PokemonTypes, error)
}
