# PokemonTypes (PokeApi)

PokemonTypes challenge for Software Design II (MDAS La Salle)

### Contributors:

ezequiel.gomez@students.salle.url.edu
perepadial@gmail.com

## Installation steps

_This project requires Go +1.13 and Go module support._

1. Install go following these instructions:
   https://go.dev/doc/install

2. Run application in command mode

```
go run . -getPokemonTypes [pokemonName]
```

After application started, a webserver will be held on port 5001, so that you can get types via http request:

example: http://localhost:5001/pokemon-types/pikachu
