package repository

import (
	"log"

	"gorm.io/gorm"
)

type User interface {
	CheckUserInRoom(userID string, roomID string) (bool, error)
	AddUserToRoom(userID string, roomID string) (bool, error)
	RemoveUserFromRoom(userID string, roomID string) (bool, error)
}

type user struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) User {
	return &user{
		db: db,
	}
}

func (u *user) CheckUserInRoom(userID string, roomID string) (bool, error) {	
	var users []struct {
		UserID string
		RoomID string
	}
	
	result := u.db.Raw("SELECT * FROM users_rooms WHERE user_id = ? AND room_id = ?", userID, roomID).Scan(&users)
	if result.Error != nil {
		log.Printf("Error at repository.CheckUserInRoom: %s\n", result.Error)
		return false, result.Error
	}

	return len(users) != 0, nil
}

func (u *user) AddUserToRoom(userID string, roomID string) (bool, error) {	
	result := u.db.Exec("INSERT INTO users_rooms (user_id, room_id) values (?, ?)", userID, roomID)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (u *user) RemoveUserFromRoom(userID string, roomID string) (bool, error) {	
	result := u.db.Exec("DELETE FROM users_rooms WHERE user_id = ? AND room_id = ?", userID, roomID)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil

}
