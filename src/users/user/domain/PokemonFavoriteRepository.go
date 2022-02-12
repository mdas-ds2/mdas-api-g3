package user

type FavoritePokemonRepository interface {
	Add(userId Id, favoritePokemonId FavoritePokemonId)
	FindAll(userId Id) []FavoritePokemonId
}
