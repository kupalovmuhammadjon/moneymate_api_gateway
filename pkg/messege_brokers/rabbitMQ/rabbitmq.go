package rabbitmq

import (
	"context"

	"github.com/cenkalti/backoff/v4"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQProducer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

type IRabbitMQProducer interface {
	ProduceMessage(context.Context, string, []byte) error
	Close()
}

func NewRabbitMQProducer() (IRabbitMQProducer, error) {
	var conn *amqp.Connection
	var err error

	operation := func() error {
		conn, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		return err
	}

	err = backoff.Retry(operation, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = setUpExchanges(channel)
	if err != nil {
		return nil, err
	}

	return &RabbitMQProducer{
		connection: conn,
		channel:    channel,
	}, nil
}

func setUpExchanges(channel *amqp.Channel) error {
	var (
		err error
	)
	_, err = channel.QueueDeclare(
		"lessons", // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	err = channel.ExchangeDeclare(
		"create",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = channel.ExchangeDeclare(
		"update",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = channel.QueueBind(
		"lessons",
		"",
		"create",
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = channel.QueueBind(
		"lessons",
		"",
		"update",
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQProducer) ProduceMessage(ctx context.Context, exchangeName string, message []byte) error {

	err := r.channel.PublishWithContext(
		ctx,
		exchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQProducer) Close() {
	r.channel.Close()
	r.connection.Close()
}
