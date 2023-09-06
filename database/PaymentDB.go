package database

import (
	"fmt"
	"paymentsystem/entity"
	"time"
)

// Local database == slice
var PaymentDetails = make([]entity.Payment, 0)

type PaymentDataBase struct {
}

func (pdb PaymentDataBase) CreatePayment(newPayment entity.Payment) entity.Payment {

	PaymentDetails = append(PaymentDetails, newPayment)
	return newPayment
}

func (pdb PaymentDataBase) GetPayment() []entity.Payment {
	return PaymentDetails
}

func (pdb PaymentDataBase) GetSingleDetail(PaymentId int) []entity.Payment {
	return PaymentDetails
}

func (pdb PaymentDataBase) ProcessPayment(updatedPayment entity.Payment) (entity.Payment, error) {
	for i, payment := range PaymentDetails {
		if payment.PaymentId == updatedPayment.PaymentId {
			updatedPayment.CreatedAt = PaymentDetails[i].CreatedAt
			updatedPayment.UpdatedAt = time.Now().Format("02-01-2006T15:04:05")
			PaymentDetails[i] = updatedPayment
			return updatedPayment, nil

		}
	}
	return entity.Payment{}, fmt.Errorf("Payment not found")

}

func (pdb PaymentDataBase) GetPaymentStatus(PaymentId int) (string, error) {
	for _, payment := range PaymentDetails {
		if payment.PaymentId == PaymentId {
			return payment.Status, nil

		}
	}
	return " ", fmt.Errorf("Payment Not Found")
}
