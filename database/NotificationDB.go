package database

import (
	"context"
	"paymentsystem/entity"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

/*
Notification Database Struct
*/

type NotificationDataBase struct {
	ConnectionObjectClient *firestore.Client
}

/*
	Send Payment Notification Database
*/

func (ndb NotificationDataBase) SendPaymentNotification(newNotification entity.Notification) entity.Notification {
	ndb.ConnectionObjectClient.Collection("PaymentNotifications").Doc(newNotification.Id).Set(context.Background(), newNotification)
	return newNotification

}

/*
Get Notification database
*/
func (ndb NotificationDataBase) GetNotification() []entity.Notification {
	var notification []entity.Notification
	iter := ndb.ConnectionObjectClient.Collection("PaymentNotifications").Documents(context.Background())
	for {

		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		mapVar := doc.Data()
		var notificationObject = entity.Notification{
			Id:        mapVar["Id"].(string),
			Type:      mapVar["Type"].(string),
			Date:      mapVar["Date"].(string),
			CreatedAt: mapVar["CreatedAt"].(string),
			UpdatedAt: mapVar["UpdatedAt"].(string),
		}

		notification = append(notification, notificationObject)
	}

	return notification
}
