package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Payment struct {
	PaymentId int     `json:"id"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

var PaymentDetails = make([]Payment, 0)

func createPayment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newPayment Payment
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

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		w.Header().Set("Content-Type", "Aplication/json")
		err := json.NewEncoder(w).Encode(PaymentDetails)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getSingleDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		paymentIdStr := r.URL.Query().Get("id")
		paymentId, err := strconv.Atoi(paymentIdStr)
		if err != nil {
			http.Error(w, "Invalid Payment ID", http.StatusBadRequest)
		}
		var idFound *Payment
		for _, payment := range PaymentDetails {
			if payment.PaymentId == paymentId {
				idFound = &payment
				w.Header().Set("Content-Type", "Aplication/json")

				err = json.NewEncoder(w).Encode(payment)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				fmt.Println(PaymentDetails[paymentId])
				break

			}
		}
		if idFound != nil {
			w.WriteHeader(http.StatusCreated)

		} else {
			http.Error(w, "Payment not found", http.StatusNotFound)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func processPayment(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		var updatedPayment Payment
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

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getPaymentStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		paymentIDStr := r.URL.Query().Get("id")
		paymentID, err := strconv.Atoi(paymentIDStr)
		if err != nil {
			http.Error(w, "Invalid payment ID", http.StatusBadRequest)
			return
		}
		var foundPayment *Payment
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
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/createPayment", createPayment)
	http.HandleFunc("/getPayment", index)
	http.HandleFunc("/processPayment", processPayment)
	http.HandleFunc("/paymentStatus", getPaymentStatus)
	http.HandleFunc("/singleDetail", getSingleDetail)
	fmt.Println("Server started..!")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
