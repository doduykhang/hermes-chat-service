package dto

type Message struct {
	RoomID string `json:"roomID"`
	UserID string `json:"userID"`
	UserName string `json:"userName"`
	Avatar string `json:"avatar"`
	Content string `json:"content"`
}
