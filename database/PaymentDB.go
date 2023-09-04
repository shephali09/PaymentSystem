package database

import (
	"paymentsystem/entity"
)

var PaymentDetails = make([]entity.Payment, 0)

type PaymentDataBase struct {
}

/*func (pdb PaymentDataBase) CreatePayment(newPayment entity.Payment) {
	newPayment.CreatedAt = time.Now().Format("02-01-2006T15:04:05")
	newPayment.UpdatedAt = time.Now().Format("02-01-2006T15:04:05")
	PaymentDetails = append(PaymentDetails, newPayment)
}*/

/*func (pdb PaymentDataBase) GetSingleDetail(controller.PaymentId) entity.Payment {
	for _, payment := range controller.PaymentDetails {
		if payment.PaymentId == entity.PaymentId {
			return payment
		}
	}

}*/
