package config

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQConnection(config *Config) (*amqp.Connection, error) {
	return amqp.Dial(getRabbitMQConnString(config.RabbitMQ))
}

func getRabbitMQConnString(rabbitMQ RabbitMQ) string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		rabbitMQ.User,
		rabbitMQ.Password,
		rabbitMQ.Host,
		rabbitMQ.Port,
	)
}
