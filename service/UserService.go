package service

import (
	"paymentsystem/database"
	"paymentsystem/entity"
	"time"
)

type UserService struct {
	Database database.UserDataBase
}

func (us UserService) GetUser() []entity.User {
	return us.Database.GetUser()

}

func (us UserService) CreateUser(newUser entity.User) entity.User {
	newUser.CreatedAt = time.Now().Format("02-01-2006T15:04:05")
	newUser.UpdatedAt = time.Now().Format("02-01-2006T15:04:05")
	return us.Database.CreateUser(newUser)

}

func (us UserService) UpdateUser(updatedUser entity.User) (entity.User, error) {
	newUpdatedUser, err := us.Database.UpdateUser(updatedUser)
	if err != nil {
		return entity.User{}, nil
	}

	return newUpdatedUser, nil

}

func (us UserService) DeleteUser(userId int) error {
	err := us.Database.DeleteUser(userId)

	if err != nil {
		return err
	}
	return nil

}
