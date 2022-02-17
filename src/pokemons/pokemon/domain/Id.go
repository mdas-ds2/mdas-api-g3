package pokemon

type Id struct {
	value int
}

func (id Id) GetValue() int {
	return id.value
}

func CreateId(value int) Id {
	return Id{value}
}
