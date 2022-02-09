# PokemonTypes (PokeApi)

PokemonTypes challenge for Software Design II (MDAS La Salle)

### Contributors:

- ezequiel.gomez@students.salle.url.edu
- perepadial@gmail.com

## Installation steps

_This project requires Go +1.13 and Go module support._

1. Clone the repository
   git clone https://github.com/mdas-ds2/mdas-api-g3.git .

2. Build docker image

```
docker build -t pokeapiwebserver .
```

3. Run the application in a docker container.

```
docker run -p 5001:5001 -it pokeapiwebserver /poke-api/main -getPokemonTypes [pokemonName]
```

After application started, a webserver will be held on port 5001, so that you can get types via http request:

example: http://localhost:5001/pokemon-types/pikachu
