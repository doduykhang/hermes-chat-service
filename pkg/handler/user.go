package handler

import (
	"doduykhang/hermes-chat/pkg/dto"
	"doduykhang/hermes-chat/pkg/service"
)

type User struct{
	service service.User
	queue service.Queue
}

func NewUser(s service.User, q service.Queue) *User{
	return &User{
		service: s,
		queue: q,
	}
}

func (user *User) HandleAddUser() {
	ch := make(chan dto.UserRoom)
	go user.queue.SubAddUserEvent(ch)
	for request := range ch {
		user.service.AddUserToRoom(request.UserID, request.RoomID)
	}
}

func (user *User) HandleRemoveUser() {
	ch := make(chan dto.UserRoom)
	go user.queue.SubDeleteUserEvent(ch)
	for request := range ch {
		user.service.RemoveUserFromRoom(request.UserID, request.RoomID)
	}
}
