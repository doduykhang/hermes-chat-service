package service

import (
	"doduykhang/hermes-chat/pkg/dto"
	"encoding/json"
	"log"

	"github.com/go-redis/redis"
)

type PubSub interface {
	Pub(message *dto.Message) error
	Sub(ch chan dto.Message) error
}

type pubSub struct {
	redis *redis.Client
}

func NewPubSubService(redis *redis.Client) PubSub {
	return &pubSub{
		redis: redis,
	}
}

func (p pubSub) Pub(message *dto.Message) (error) {
	byte, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshal message, %v\n", err)
	}
	err = p.redis.Publish("message", string(byte)).Err()
	if err != nil {
		log.Printf("Error publishing message, %v\n", err)
	}
	return nil
}
func (p pubSub) Sub(ch chan dto.Message) (error) {
	pubsub := p.redis.Subscribe("message")
	defer pubsub.Close()
	msgCh := pubsub.Channel()

	for msg := range msgCh {
		var message dto.Message
		err := json.Unmarshal([]byte(msg.Payload), &message)
		if err != nil {
			log.Printf("Error unmarshal message, %v\n", err)
		}

		ch <- message
	}
	return nil 
}


