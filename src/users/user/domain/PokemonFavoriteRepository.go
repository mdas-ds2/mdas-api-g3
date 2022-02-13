package user

type FavoritePokemonRepository interface {
	Add(userId UserId, favoritePokemonId FavoritePokemonId) error
	FindAll(userId UserId) []FavoritePokemonId
}
