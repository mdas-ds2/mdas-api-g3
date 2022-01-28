package pokeApi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	pokemon "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon/domain"
)

type PokeApiPokemonRepository struct{}

func (pokeApiRepository PokeApiPokemonRepository) FindByName(pokemonName pokemon.PokemonName) pokemon.Pokemon {
	pokemonApiResponse, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + pokemonName.GetValue())

	if err != nil {
		log.Fatalln(err)
	}

	defer pokemonApiResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(pokemonApiResponse.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var result Pokemon
	json.Unmarshal(responseBody, &result)

	var pokemonTypes = []pokemon.PokemonType{}

	for _, pokemonType := range result.Types {
		pokemonTypeType, err := pokemon.CreatePokemonType(pokemonType.Type.Name)

		if err != nil {
			log.Fatalln(err)
		}

		pokemonTypes = append(pokemonTypes, *pokemonTypeType)
	}

	pokemon, err := pokemon.CreatePokemon(pokemonName, pokemonTypes)

	if err != nil {
		log.Fatalln(err)
	}

	return *pokemon
}
