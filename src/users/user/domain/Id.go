package user

// TODO: User UUID instead of string for IDs
type Id struct {
	value string
}

func (id Id) GetValue() string {
	return id.value
}

func CreateId(id string) Id {
	return Id{id}
}
