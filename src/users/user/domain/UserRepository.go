package user

type UserRepository interface {
	Save(user User) error
	Find(id UserId) (User, error)
}
