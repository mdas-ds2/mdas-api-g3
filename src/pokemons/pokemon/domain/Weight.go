package pokemon

type Weight struct {
	value int
}

func (weight Weight) GetValue() int {
	return weight.value
}

func CreateWeight(value int) Weight {
	return Weight{value}
}
