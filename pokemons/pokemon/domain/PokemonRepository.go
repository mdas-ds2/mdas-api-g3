package pokemon

type PokemonRepository interface {
	FindByName(pokemonName PokemonName) (Pokemon, error)
}
