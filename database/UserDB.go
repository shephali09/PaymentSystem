package database

import (
	"context"
	"fmt"
	"paymentsystem/entity"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

var UserDetails = make([]entity.User, 0)

type UserDataBase struct {
	ConnectionObjectClient *firestore.Client
}

func (udb UserDataBase) GetUser() []entity.User {
	var user []entity.User
	iter := udb.ConnectionObjectClient.Collection("User").Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		mapVar := doc.Data()
		var userObject = entity.User{
			UserId:    mapVar["UserId"].(string),
			Name:      mapVar["Name"].(string),
			Email:     mapVar["Email"].(string),
			Password:  mapVar["Password"].(string),
			CreatedAt: mapVar["CreatedAt"].(string),
			UpdatedAt: mapVar["UpdatedAt"].(string),
		}
		user = append(user, userObject)
	}
	return user
}

func (udb UserDataBase) CreateUser(newUser entity.User) entity.User {
	udb.ConnectionObjectClient.Collection("User").Doc(newUser.UserId).Set(context.Background(), newUser)
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

func (udb UserDataBase) DeleteUser(userId string) error {

	/*for i, user := range UserDetails {
		if user.UserId == userId {
			UserDetails = append(UserDetails[:i], UserDetails[i+1:]...)
			return nil
		}
	}*/
	//return fmt.Errorf("User not found")

	fmt.Println(userId)
	_, error := udb.ConnectionObjectClient.Collection("User").Doc(userId).Delete(context.Background())
	return error

}
