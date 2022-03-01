package user

import (
	"errors"

	domain "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMqEventPublisher struct {
}

func (eventPublisher RabbitMqEventPublisher) publishEvents(events []domain.FavoritePokemonAddedEvent) error {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/user")

	if err != nil {
		return errors.New("Failed on create the connection: " + err.Error())
	}
	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		return errors.New("Failed on create the channel: " + err.Error())
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
		return errors.New("Failed on create the queue: " + err.Error())
	}

	for _, event := range events {
		err := eventPublisher.publishEventMq(ch, q, event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (eventPublisher RabbitMqEventPublisher) publishEventMq(ch *amqp.Channel, q amqp.Queue, event domain.FavoritePokemonAddedEvent) error {
	err := ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event.GetContent()),
		},
	)

	if err != nil {
		return errors.New("Failed on publish the message: " + err.Error())
	}

	return nil
}
