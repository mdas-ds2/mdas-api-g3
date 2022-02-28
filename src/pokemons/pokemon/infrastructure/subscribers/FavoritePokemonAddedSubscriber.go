package pokemon

import (
	"errors"
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

	q, err := createQueue(ch)

	if err != nil {
		return errors.New("error setting up the connection: " + err.Error())
	}

	msgs, err := registerConsumer(ch, q.Name)

	if err != nil {
		return errors.New("failed to register a consumer: " + err.Error())
	}

	listenEvents(msgs, useCase)

	return nil
}

func createQueue(channel *amqp.Channel) (amqp.Queue, error) {
	return channel.QueueDeclare(
		"notify_favorite_pokemon_added",
		false,
		false,
		false,
		false,
		nil,
	)
}

func registerConsumer(channel *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
	return channel.Consume(
		queueName,
		"pokemon-detail",
		true,
		false,
		false,
		false,
		nil,
	)
}

func listenEvents(msgs <-chan amqp.Delivery, useCase application.GetPokemonDetails) {
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			pokemon := findPokemon(d, useCase)
			increaseFavoriteTimes(pokemon, useCase)
		}
	}()

	<-forever
}

func findPokemon(delivery amqp.Delivery, useCase application.GetPokemonDetails) domain.Pokemon {
	pokemonId := getPokemonId(delivery)
	pokemon, _ := useCase.Repository.Find(pokemonId)
	return pokemon
}

func getPokemonId(delivery amqp.Delivery) domain.Id {
	pokemonId, _ := strconv.Atoi(string(delivery.Body))
	return domain.CreateId(pokemonId)
}

func increaseFavoriteTimes(pokemon domain.Pokemon, useCase application.GetPokemonDetails) {
	pokemon.IncreaseFavoriteTimes()
	useCase.Repository.Save(pokemon)
}
