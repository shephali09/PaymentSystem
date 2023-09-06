package database

import (
	"paymentsystem/entity"
)

type NotificationDataBase struct {
}

var NotificationDetails = make([]entity.Notification, 0)

func (ndb NotificationDataBase) SendPaymentNotification(newNotification entity.Notification) entity.Notification {
	NotificationDetails = append(NotificationDetails, newNotification)
	return newNotification

}

func (ndb NotificationDataBase) GetNotification() []entity.Notification {
	return NotificationDetails
}
