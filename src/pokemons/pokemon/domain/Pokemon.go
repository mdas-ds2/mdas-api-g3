package pokemon

type Pokemon struct {
	id              Id
	name            Name
	height          Height
	weight          Weight
	timesAsFavorite TimesAsFavorite
}

func CreatePokemon(id Id, name Name, height Height, weight Weight, timesAsFavorite TimesAsFavorite) Pokemon {
	return Pokemon{id, name, height, weight, timesAsFavorite}
}

func (pokemon *Pokemon) GetId() Id {
	return pokemon.id
}

func (pokemon *Pokemon) GetName() Name {
	return pokemon.name
}

func (pokemon *Pokemon) GetHeight() Height {
	return pokemon.height
}

func (pokemon *Pokemon) GetWeight() Weight {
	return pokemon.weight
}

func (pokemon *Pokemon) GetTimesAsFavorite() TimesAsFavorite {
	return pokemon.timesAsFavorite
}

func (pokemon *Pokemon) IncreaseFavoriteTimes() {
	pokemon.timesAsFavorite = CreateTimesAsFavorite(pokemon.timesAsFavorite.value + 1)
}
