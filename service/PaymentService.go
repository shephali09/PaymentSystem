package service

import (
	"fmt"
	"paymentsystem/database"
	"paymentsystem/entity"
	"time"
)

/*
Payment Service struct
*/
type PaymentService struct {
	Database database.PaymentDataBase
}

/*
Create Payment service
*/
func (ps PaymentService) CreatePayment(newPayment entity.Payment) entity.Payment {

	// starting of the business logic if there is any.
	newPayment.CreatedAt = time.Now().Format("02-01-2006T15:04:05")
	newPayment.UpdatedAt = time.Now().Format("02-01-2006T15:04:05")

	// We will save to database after the business logic is done.
	return ps.Database.CreatePayment(newPayment)
}

/*
Get Payment service
*/
func (ps PaymentService) GetPayment() []entity.Payment {
	return ps.Database.GetPayment()
}

/*
Get Single Payment Detail service
*/
func (ps PaymentService) GetSingleDetail(PaymentId string) (entity.Payment, error) {
	paymentDetails := ps.Database.GetSingleDetail(PaymentId)
	for _, payment := range paymentDetails {
		if payment.PaymentId == PaymentId {
			return payment, nil
		}

	}
	return entity.Payment{}, fmt.Errorf("payment Not found")
}

/*
Process Payment service
*/
func (ps PaymentService) ProcessPayment(updatedPayment entity.Payment) (entity.Payment, error) {
	processedPayment, err := ps.Database.ProcessPayment(updatedPayment)
	if err != nil {
		return entity.Payment{}, nil
	}

	return processedPayment, nil

}

/*
Get Payment Status service
*/
func (ps PaymentService) GetPaymentStatus(PaymentId string) (string, error) {
	status, err := ps.Database.GetPaymentStatus(PaymentId)
	if err != nil {
		return "", err
	}
	return status, nil

}
