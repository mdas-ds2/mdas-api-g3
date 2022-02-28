package pokemon

import (
	"errors"
	"log"
	"strconv"

	application "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/application"
	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

type FavoritePokemonAddedSubscriber struct {
}

func (subscriber FavoritePokemonAddedSubscriber) RegisterSubscriber(useCase application.GetPokemonDetails) error {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/user")

	if err != nil {
		return errors.New("error setting up the connection: " + err.Error())
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return errors.New("error setting up the channel: " + err.Error())
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"notify_favorite_pokemon_added",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return errors.New("error setting up the connection: " + err.Error())
	}

	msgs, err := ch.Consume(
		q.Name,
		"pokemon-detail",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return errors.New("failed to register a consumer: " + err.Error())
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			pokemonId, _ := strconv.Atoi(string(d.Body))
			pokemon, _ := useCase.Repository.Find(domain.CreateId(pokemonId))
			pokemon.IncreaseFavoriteTimes()
			useCase.Repository.Save(pokemon)
			log.Printf("pokemon favorite addetd corrrectly")
		}
	}()

	<-forever

	return nil
}
