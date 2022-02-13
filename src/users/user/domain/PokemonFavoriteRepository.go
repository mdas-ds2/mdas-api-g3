package user

type FavoritePokemonRepository interface {
	Add(user User, favoritePokemonId PokemonId) error
	FindAll(userId UserId) []PokemonId
}
