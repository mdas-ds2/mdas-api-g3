package user

type PokemonFavoriteRepository interface {
	InsertFavoritePokemon(user User, favoritePokemonId FavoritePokemonId)
}
