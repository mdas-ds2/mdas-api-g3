package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	pokemon "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon/domain"
)

func main() {
	pokemonType, error := pokemon.CreatePokemonType("electric")
	pokemonName, error := pokemon.CreatePokemonName("Pikachu")
	pokemonTypes := []pokemon.PokemonType{*pokemonType}

	pikachu, error := pokemon.CreatePokemon(*pokemonName, pokemonTypes)

	if error != nil {
		log.Fatalln(error)
	}

	pokemonNameToPrint := flag.String("pokemonName", pikachu.Name(), "The type of the pokemon we want to search for")

	flag.Parse()

	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + *pokemonNameToPrint)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var result Pokemon

	json.Unmarshal(body, &result)

	fmt.Println(pikachu.GetStringFormatedTypes())
}
