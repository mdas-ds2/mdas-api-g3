package pokemonType

type PokemonTypeRepository interface {
	FindByPokemonName(pokemonName PokemonName) (TypeCollection, error)
}
