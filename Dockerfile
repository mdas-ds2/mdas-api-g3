FROM golang:rc-alpine3.15

EXPOSE 5001

RUN mkdir /poke-api

COPY . /poke-api

WORKDIR /poke-api

RUN apk add git

RUN go mod download

RUN go build -o app .

CMD ./app -getPokemonTypes pikachu