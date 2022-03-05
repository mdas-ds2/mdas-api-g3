package pokemon

type TimesAsFavorite struct {
	value int
}

func (times TimesAsFavorite) GetValue() int {
	return times.value
}

func CreateTimesAsFavorite(value int) TimesAsFavorite {
	return TimesAsFavorite{value}
}
