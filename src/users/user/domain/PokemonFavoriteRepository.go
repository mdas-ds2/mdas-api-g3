package user

type FavoritePokemonRepository interface {
	Add(userId UserId, favoritePokemonId PokemonId) error
	FindAll(userId UserId) []PokemonId
}
