package pokemonType

type PokemonTypeRepository interface {
	FindByPokemonName(pokemonName PokemonName) ([]PokemonType, error)
}
