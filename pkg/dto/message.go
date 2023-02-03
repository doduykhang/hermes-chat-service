package dto

import "time"

type Sender struct {
	ID string `json:"id"`	
	FirstName string `json:"firstName"`	
	LastName string `json:"lastName"`	
	Avatar string `json:"avatar"`	
}

type Message struct {
	ID string `json:"id"`
	RoomID string `json:"roomID"`
	UserID string `json:"userID"`
	CreatedAt time.Time `json:"createdAt"`
	Content string `json:"content"`
	Sender Sender `json:"sender"`
}
