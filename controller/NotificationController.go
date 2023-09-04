package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"paymentsystem/entity"
	util "paymentsystem/utility"
	"time"
)

var NotificationDetails = make([]entity.Notification, 0)

type NotificationController struct {
}

func (nc NotificationController) SendPaymentNotification(w http.ResponseWriter, r *http.Request) {
	util.CheckMethod(r, w, http.MethodPost)

	var newNotification entity.Notification
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
}

func (nc NotificationController) GetNotification(w http.ResponseWriter, r *http.Request) {
	util.CheckMethod(r, w, http.MethodGet)

	w.Header().Set("Content-Type", "Application/json")
	err := json.NewEncoder(w).Encode(NotificationDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

}
