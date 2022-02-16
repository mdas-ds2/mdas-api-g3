package user

type FavoritePokemonRepository interface {
	Save(user User) error
	Find(userId UserId) *User
}
