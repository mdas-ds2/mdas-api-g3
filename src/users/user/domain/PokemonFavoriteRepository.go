package user

type FavoritePokemonRepository interface {
	Add(userId Id, favoritePokemonId FavoritePokemonId) error
	FindAll(userId Id) []FavoritePokemonId
}
