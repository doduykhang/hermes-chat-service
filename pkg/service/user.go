package service

import "doduykhang/hermes-chat/pkg/repository"

type User interface {
	AddUserToRoom(userID string, roomID string) (bool, error)
	RemoveUserFromRoom(userID string, roomID string) (bool, error)
}

type user struct {
	userRepo repository.User
}

func NewUserService(userRepo repository.User) User {
	return &user{
		userRepo: userRepo,
	}
}

func (user *user) AddUserToRoom(userID string, roomID string) (bool, error) {
	return user.userRepo.AddUserToRoom(userID, roomID)
}

func (user *user) RemoveUserFromRoom(userID string, roomID string) (bool, error) {
	return user.userRepo.RemoveUserFromRoom(userID, roomID) 
}
