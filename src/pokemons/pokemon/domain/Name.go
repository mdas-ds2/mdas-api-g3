package pokemon

type Name struct {
	value string
}

func (name Name) GetValue() string {
	return name.value
}

func CreateName(value string) Name {
	return Name{value}
}
