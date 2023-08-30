package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Notification struct {
	Id        int    `json:"id"`
	Type      string `json:"type"`
	Date      string `json:"date"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

var NotificationDetails = make([]Notification, 0)

func sendPaymentNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newNotification Notification
		err := json.NewDecoder(r.Body).Decode(&newNotification)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newNotification.CreatedAt = time.Now().Format("02-01-2006T15:04:05")
		newNotification.UpdatedAt = time.Now().Format("02-01-2006T15:04:05")
		NotificationDetails = append(NotificationDetails, newNotification)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, newNotification)

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "Application/json")
		err := json.NewEncoder(w).Encode(NotificationDetails)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/createNotification", sendPaymentNotification)
	http.HandleFunc("/getNotification", getNotification)
	fmt.Println("Server started..!")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
