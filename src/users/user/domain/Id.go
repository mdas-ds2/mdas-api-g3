package user

type Id struct {
	value string
}

func (id Id) GetValue() string {
	return id.value
}

func CreateId(id string) Id {
	return Id{id}
}
