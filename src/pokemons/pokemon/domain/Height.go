package pokemon

type Height struct {
	value int
}

func (height Height) GetValue() int {
	return height.value
}

func CreateHeight(value int) Height {
	return Height{value}
}
