package user

type User struct {
	id UserId
}

func (user User) GetId() UserId {
	return user.id
}

func CreateUser(id UserId) *User {
	user := &User{
		id: id,
	}

	return user
}
