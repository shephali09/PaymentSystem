package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"paymentsystem/entity"
	util "paymentsystem/utility"
	"strconv"
	"time"
)

var PaymentDetails = make([]entity.Payment, 0)

type PaymentController struct {
}

func (pc PaymentController) CreatePayment(w http.ResponseWriter, r *http.Request) {

	util.CheckMethod(r, w, http.MethodPost)

	//var PaymentDetails []entity.Payment
	var newPayment entity.Payment
	err := json.NewDecoder(r.Body).Decode(&newPayment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newPayment.CreatedAt = time.Now().Format("02-01-2006T15:04:05")
	newPayment.UpdatedAt = time.Now().Format("02-01-2006T15:04:05")
	PaymentDetails = append(PaymentDetails, newPayment)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, newPayment)

}

func (pc PaymentController) GetPayment(w http.ResponseWriter, r *http.Request) {

	util.CheckMethod(r, w, http.MethodGet)

	//var PaymentDetails []entity.Payment
	w.Header().Set("Content-Type", "aplication/json")
	err := json.NewEncoder(w).Encode(PaymentDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (pc PaymentController) GetSingleDetail(w http.ResponseWriter, r *http.Request) {
	util.CheckMethod(r, w, http.MethodGet)

	PaymentIdStr := r.URL.Query().Get("id")
	PaymentId, err := strconv.Atoi(PaymentIdStr)
	if err != nil {
		http.Error(w, "Invalid Payment ID", http.StatusBadRequest)
	}

	//var PaymentDetails []entity.Payment
	var idFound *entity.Payment
	for _, payment := range PaymentDetails {
		if payment.PaymentId == PaymentId {
			idFound = &payment
			w.Header().Set("Content-Type", "Aplication/json")

			err := json.NewEncoder(w).Encode(payment)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			fmt.Println(PaymentDetails[PaymentId])
			break
		}
	}
	if idFound != nil {
		w.WriteHeader(http.StatusCreated)

	} else {
		http.Error(w, "Payment not found", http.StatusNotFound)
	}

}

func (pc PaymentController) ProcessPayment(w http.ResponseWriter, r *http.Request) {

	util.CheckMethod(r, w, http.MethodPut)

	//var PaymentDetails []entity.Payment
	var updatedPayment entity.Payment
	err := json.NewDecoder(r.Body).Decode(&updatedPayment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, payment := range PaymentDetails {
		if payment.PaymentId == updatedPayment.PaymentId {
			updatedPayment.CreatedAt = PaymentDetails[i].CreatedAt
			updatedPayment.UpdatedAt = time.Now().Format("02-01-2006T15:04:05")
			PaymentDetails[i] = updatedPayment
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, updatedPayment)
		}
	}

}

func (pc PaymentController) GetPaymentStatus(w http.ResponseWriter, r *http.Request) {

	util.CheckMethod(r, w, http.MethodGet)

	paymentIDStr := r.URL.Query().Get("id")
	paymentID, err := strconv.Atoi(paymentIDStr)
	if err != nil {
		http.Error(w, "Invalid payment ID", http.StatusBadRequest)
		return
	}

	//var PaymentDetails []entity.Payment
	var foundPayment *entity.Payment
	for _, payment := range PaymentDetails {
		if payment.PaymentId == paymentID {
			foundPayment = &payment
			break
		}
	}
	if foundPayment != nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Payment Status : %s", foundPayment.Status)
	} else {
		http.Error(w, "Payment not found", http.StatusNotFound)
	}

}
