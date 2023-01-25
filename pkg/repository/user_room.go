package repository

import (
	"log"

	"gorm.io/gorm"
)

type User interface {
	CheckUserInRoom(userID string, roomID string) (bool, error)
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
	var user struct {
		UserID string
		RoomID string
	}
	result := u.db.Raw("SELECT * FROM users_rooms WHERE user_id = ? AND room_id = ?", userID, roomID).Scan(&user)
	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			log.Printf("Error at repository.CheckUserInRoom: %s\n", result.Error)
			return false, result.Error
		}
		return true, nil
	}
	return true, nil
}
