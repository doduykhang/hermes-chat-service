package config

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQConnection(config *Config) (*amqp.Connection, error) {
	conn := getRabbitMQConnString(config.RabbitMQ)
	fmt.Println(conn)
	return amqp.Dial(conn)
}

func getRabbitMQConnString(rabbitMQ RabbitMQ) string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s",
		rabbitMQ.Protocol,
		rabbitMQ.User,
		rabbitMQ.Password,
		rabbitMQ.Host,
		rabbitMQ.Port,
		rabbitMQ.VHost,
	)
}
