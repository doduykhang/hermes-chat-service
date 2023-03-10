package service

import (
	"doduykhang/hermes-chat/pkg/dto"
	"doduykhang/hermes-chat/pkg/repository"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/olahol/melody"
)

type Chat interface {
	ConnectToRoom(w http.ResponseWriter, r *http.Request, roomID string, userID string) error
	BroadcastToRoom(roomID string, message *dto.Message) error
	HandleMessage() error
	HandleWaitForMessage() error
}
 
type chat struct {
	melody *melody.Melody
	pubSub PubSub
	userRepo repository.User
	queue Queue
}

func NewChat(m *melody.Melody, p PubSub, userRepo repository.User, q Queue) Chat {
	return &chat{
		melody: m,	
		pubSub: p,
		userRepo: userRepo,
		queue: q,
	}
}	

func (chat *chat) ConnectToRoom(w http.ResponseWriter, r *http.Request, roomID string, userID string) (error) {
	
	check, err := chat.userRepo.CheckUserInRoom(userID, roomID)
	if err != nil {
		log.Printf("Error at service.chat.ConnectToRoom: %v\n", err)
		return err
	}

	if !check {
		log.Printf("You don't have access to this room")
		return errors.New("You don't have access to this room")
	}

	keys := make(map[string]interface{})
	keys["roomID"] = roomID
	keys["userID"] = userID

	return chat.melody.HandleRequestWithKeys(w, r, keys)
}

func (chat *chat) HandleMessage() error {
	chat.melody.HandleMessage(func(s *melody.Session, data []byte) {
		var message dto.Message
		err := json.Unmarshal(data, &message)
		message.CreatedAt = time.Now()
		message.ID = uuid.New().String()

		if err != nil {
			log.Printf("Error unmarshal message, %s\n", err)
			return
		}

		roomIDKey, ok := s.Keys["roomID"]
		userIDkey, ok := s.Keys["userID"]
		if !ok {
			log.Printf("Error getting room id\n")
			return
		}

		message.RoomID = roomIDKey.(string)
		message.UserID = userIDkey.(string)	
		chat.pubSub.Pub(&message)	
		chat.queue.PubAddMessage(message)
	})
	return nil
}


func (chat *chat) HandleWaitForMessage() error {
	ch := make(chan dto.Message)
	go chat.pubSub.Sub(ch)	

	for msg := range ch {
		fmt.Println("message sub", msg)
		chat.BroadcastToRoom(msg.RoomID, &msg)
	}
	return nil
}

func (chat *chat) BroadcastToRoom(roomID string, message *dto.Message) (error) {
	b, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshaling message %s\n", err)
		return err
	}
	return chat.melody.BroadcastFilter(b, func (s *melody.Session) bool {
		roomIDKey, ok := s.Keys["roomID"]
		if !ok {
			log.Printf("Error getting room id")
			return false	
		} 
		return roomIDKey == roomID
	})	
}




