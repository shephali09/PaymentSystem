package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"paymentsystem/entity"
	"paymentsystem/service"
	util "paymentsystem/utility"
)

/*
Payment Controller Struct
*/
type PaymentController struct {
	Service service.PaymentService
}

/*
Create Payment endpoint
*/
func (pc PaymentController) CreatePayment(w http.ResponseWriter, r *http.Request) {

	// checking for the method and responding with the error
	util.CheckMethod(r, w, http.MethodPost)

	// getting the data from the request body
	var newPayment entity.Payment
	err := json.NewDecoder(r.Body).Decode(&newPayment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newPayment = pc.Service.CreatePayment(newPayment)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, newPayment)

}

/*
Get Payment endpoint
*/
func (pc PaymentController) GetPayment(w http.ResponseWriter, r *http.Request) {

	util.CheckMethod(r, w, http.MethodGet)

	w.Header().Set("Content-Type", "aplication/json")
	paymentDetails := pc.Service.GetPayment()
	err := json.NewEncoder(w).Encode(paymentDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}

/*
Get Single Payment Detail endpoint
*/
func (pc PaymentController) GetSingleDetail(w http.ResponseWriter, r *http.Request) {
	util.CheckMethod(r, w, http.MethodGet)

	PaymentId := r.URL.Query().Get("id")

	payment, err := pc.Service.GetSingleDetail(PaymentId)
	if err != nil {
		http.Error(w, "Invalid Payment ID", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "Aplication/json")
	err = json.NewEncoder(w).Encode(payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

/*
	Process Payment endpoint
*/

func (pc PaymentController) ProcessPayment(w http.ResponseWriter, r *http.Request) {

	util.CheckMethod(r, w, http.MethodPut)

	var updatedPayment entity.Payment
	err := json.NewDecoder(r.Body).Decode(&updatedPayment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	processedPayment, err := pc.Service.ProcessPayment(updatedPayment)
	if err != nil {
		http.Error(w, "Payment Not Found", http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, processedPayment)
}

/*
Get Payment Status endpoint
*/
func (pc PaymentController) GetPaymentStatus(w http.ResponseWriter, r *http.Request) {

	util.CheckMethod(r, w, http.MethodGet)

	paymentID := r.URL.Query().Get("id")

	status, err := pc.Service.GetPaymentStatus(paymentID)
	if err != nil {
		http.Error(w, "PAyment Not Found", http.StatusNotFound)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Payment Status :%s", status)

}
