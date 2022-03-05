package pokemon

type Repository interface {
	Find(id Id) (Pokemon, error)
	Save(pokemon Pokemon)
}
