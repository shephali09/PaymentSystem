package service

import (
	"paymentsystem/database"
	"paymentsystem/entity"
	"time"
)

/*
Notification Service struct
*/
type NotificationService struct {
	Database database.NotificationDataBase
}

/*
	Send Payment Notification service
*/

func (ns NotificationService) SendPaymentNotification(newNotification entity.Notification) entity.Notification {
	newNotification.CreatedAt = time.Now().Format("02-01-2006T15:04:05")
	newNotification.UpdatedAt = time.Now().Format("02-01-2006T15:04:05")
	return ns.Database.SendPaymentNotification(newNotification)

}

/*
	Get Notification service
*/

func (ns NotificationService) GetNotification() []entity.Notification {
	return ns.Database.GetNotification()

}
