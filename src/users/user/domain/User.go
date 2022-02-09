package user

type User struct {
	id Id
}

func (user User) GetId() Id {
	return user.id
}

func CreateUser(id Id) (*User, error) {
	user := &User{
		id: id,
	}

	return user, nil
}
