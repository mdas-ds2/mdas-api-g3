package user

type FavoritePokemonIdCollection struct {
	values []PokemonId
}

func (collection *FavoritePokemonIdCollection) Add(pokemonId PokemonId) {
	collection.values = append(collection.values, pokemonId)
}

func (collection *FavoritePokemonIdCollection) Has(favoritePokemonToAdd PokemonId) bool {
	favorites := collection.values

	if len(favorites) == 0 {
		return true
	}

	for _, pokemonId := range favorites {
		if pokemonId.GetValue() == favoritePokemonToAdd.GetValue() {
			return false
		}
	}

	return true
}

func CreateFavoritePokemonIdCollection(favoritePokemonList []PokemonId) FavoritePokemonIdCollection {
	return FavoritePokemonIdCollection{favoritePokemonList}
}

func (collection FavoritePokemonIdCollection) GetValues() []PokemonId {
	return collection.values
}
