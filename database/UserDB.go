package database

import (
	"fmt"
	"paymentsystem/entity"
	"time"
)

var UserDetails = make([]entity.User, 0)

type UserDataBase struct {
}

func (udb UserDataBase) GetUser() []entity.User {
	return UserDetails

}

func (udb UserDataBase) CreateUser(newUser entity.User) entity.User {
	UserDetails = append(UserDetails, newUser)
	return newUser

}

func (udb UserDataBase) UpdateUser(updatedUser entity.User) (entity.User, error) {
	for i, user := range UserDetails {
		if user.UserId == updatedUser.UserId {
			updatedUser.CreatedAt = UserDetails[i].CreatedAt
			updatedUser.UpdatedAt = time.Now().Format("02-01-2006T15:04:05")
			UserDetails[i] = updatedUser
			return updatedUser, nil
		}
	}
	return entity.User{}, fmt.Errorf("User not found")

}

func (udb UserDataBase) DeleteUser(userId int) error {

	for i, user := range UserDetails {
		if user.UserId == userId {
			UserDetails = append(UserDetails[:i], UserDetails[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User not found")

}
