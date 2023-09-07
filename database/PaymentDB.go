package database

import (
	"context"
	"fmt"
	"paymentsystem/entity"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// Local database == slice
var PaymentDetails = make([]entity.Payment, 0) //no need because we are using firestore

type PaymentDataBase struct {
	ConnectionObjectClient *firestore.Client
}

func (pdb PaymentDataBase) CreatePayment(newPayment entity.Payment) entity.Payment {

	pdb.ConnectionObjectClient.Collection("Payments").Doc(newPayment.PaymentId).Set(context.Background(), newPayment)
	return newPayment
}

func (pdb PaymentDataBase) GetPayment() []entity.Payment {
	var payment []entity.Payment
	iter := pdb.ConnectionObjectClient.Collection("Payments").Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		mapVar := doc.Data()
		var paymentObject = entity.Payment{
			PaymentId: mapVar["PaymentId"].(string),
			Amount:    mapVar["Amount"].(float64),
			Currency:  mapVar["Currency"].(string),
			Status:    mapVar["Status"].(string),
			CreatedAt: mapVar["CreatedAt"].(string),
			UpdatedAt: mapVar["UpdatedAt"].(string),
		}
		payment = append(payment, paymentObject)
	}

	return payment
}

func (pdb PaymentDataBase) GetSingleDetail(PaymentId string) []entity.Payment {

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

func (pdb PaymentDataBase) GetPaymentStatus(PaymentId string) (string, error) {
	for _, payment := range PaymentDetails {
		if payment.PaymentId == PaymentId {
			return payment.Status, nil

		}
	}
	return " ", fmt.Errorf("Payment Not Found")
}
