package user

type UserId struct {
	value string
}

func (id UserId) GetValue() string {
	return id.value
}

func CreateUserId(id string) UserId {
	return UserId{id}
}
