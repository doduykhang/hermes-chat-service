package service

import (
	"context"
	"doduykhang/hermes-chat/pkg/dto"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue interface {
	SubAddUserEvent(chan dto.UserRoom)
	SubDeleteUserEvent(chan dto.UserRoom)
	PubAddMessage(dto.Message)
}

type queue struct {
	rabbitMq *amqp.Connection
}

func NewQueue(conn *amqp.Connection) Queue {
	return &queue{
		rabbitMq: conn,
	}
}

func (queue *queue) consumeQueue(ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
	q, err := ch.QueueDeclare(
  		queueName, // name
  		false,   // durable
  		false,   // delete when unused
  		false,   // exclusive
  		false,   // no-wait
  		nil,     // arguments
	)	

	if err != nil {
		log.Panicf("Error declare rabbit mq queue, %s\n", err)
		return nil, err
	}

	return ch.Consume(
  		q.Name, // queue
  		"",     // consumer
  		true,   // auto-ack
  		false,  // exclusive
  		false,  // no-local
  		false,  // no-wait
  		nil,    // args
	)
}

func (queue *queue) SubAddUserEvent(userCh chan dto.UserRoom) {
	ch, err := queue.rabbitMq.Channel()
	defer ch.Close()

	if err != nil {
		log.Panicf("Error open rabbit mq channel, %s\n", err)
		return	
	}

	msgs, err := queue.consumeQueue(ch, "add-user")

	for d := range msgs {
		var userRoom dto.UserRoom
		err := json.Unmarshal(d.Body, &userRoom)
		if err != nil {
			log.Printf("Error unmarshal message, %s\n", err)
			continue
		}
		userCh <- userRoom
  	}

}
func (queue *queue) SubDeleteUserEvent(userCh chan dto.UserRoom) {
	ch, err := queue.rabbitMq.Channel()
	defer ch.Close()

	if err != nil {
		log.Panicf("Error open rabbit mq channel, %s\n", err)
		return	
	}

	msgs, err := queue.consumeQueue(ch, "delete-user")

	for d := range msgs {
		var userRoom dto.UserRoom
		err := json.Unmarshal(d.Body, &userRoom)
		if err != nil {
			log.Printf("Error unmarshal message, %s\n", err)
			continue
		}
		userCh <- userRoom
  	}
}
func (queue queue) PubAddMessage(message dto.Message) {
	ch, err := queue.rabbitMq.Channel()
	defer ch.Close()

	if err != nil {
		log.Panicf("Error open rabbit mq channel, %s\n", err)
		return	
	}

	q, err := ch.QueueDeclare(
  		"hello", // name
  		false,   // durable
  		false,   // delete when unused
  		false,   // exclusive
  		false,   // no-wait
  		nil,     // arguments
	)

	if err != nil {
		log.Panicf("Error declare rabbit mq queue, %s\n", err)
		return	
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	d, err := json.Marshal(&message)

	if err != nil {
		log.Panicf("Error marshal message, %s\n", err)
		return	
	}

	err = ch.PublishWithContext(ctx,
  		"",     // exchange
  		q.Name, // routing key
  		false,  // mandatory
  		false,  // immediate
  		amqp.Publishing {
    			ContentType: "text/plain",
    			Body:        d,
  	})
}
