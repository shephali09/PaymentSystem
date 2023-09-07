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
Notification Controller struct
*/
type NotificationController struct {
	Service service.NotificationService
}

/*
Send Payment Notification endpoint
*/
func (nc NotificationController) SendPaymentNotification(w http.ResponseWriter, r *http.Request) {
	util.CheckMethod(r, w, http.MethodPost)

	var newNotification entity.Notification

	err := json.NewDecoder(r.Body).Decode(&newNotification)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newNotification = nc.Service.SendPaymentNotification(newNotification)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, newNotification)
}

/*
Get Notification endpoint
*/
func (nc NotificationController) GetNotification(w http.ResponseWriter, r *http.Request) {
	util.CheckMethod(r, w, http.MethodGet)

	w.Header().Set("Content-Type", "Application/json")
	notificationDetails := nc.Service.GetNotification()
	err := json.NewEncoder(w).Encode(notificationDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

}
