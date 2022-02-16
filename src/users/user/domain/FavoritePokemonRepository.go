package user

type FavoritePokemonRepository interface {
	Save(user User) error
	GetFavorites(userId UserId) FavoritePokemonIdCollection
	FindUser(userId UserId) *User
}
