package domain

type Pokemon struct {
	id Id
}

func CreatePokemon(id Id) Pokemon {
	return Pokemon{id}
}
