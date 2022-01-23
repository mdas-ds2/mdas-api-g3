package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	pokemonName := flag.String("pokemonName", "", "The type of the pokemon we want to search for")

	flag.Parse()

	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + *pokemonName)

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

	var pokemonTypes string

	for index, value := range result.Types {
		if index > 0 {
			pokemonTypes += ", "
		}
		pokemonTypes += value.Type.Name
	}

	fmt.Println(pokemonTypes)
}
