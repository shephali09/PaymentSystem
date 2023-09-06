package main

import (
	"fmt"
	"log"
	"net/http"
	"paymentsystem/config"
	"paymentsystem/controller"
	"paymentsystem/database"
	"paymentsystem/service"
)

func main() {

	config := config.InitConfig()
	if config == nil {
		// if any error occurs program will not run
		return
	}

	portNumber := config.String("portNumber")

	objectPaymentController := controller.PaymentController{
		Service: service.PaymentService{
			Database: database.PaymentDataBase{},
		},
	}

	objectUserController := controller.UserController{
		Service: service.UserService{
			Database: database.UserDataBase{},
		},
	}
	objectNotificationController := controller.NotificationController{
		Service: service.NotificationService{
			Database: database.NotificationDataBase{},
		},
	}

	// Routes for the Payment
	http.HandleFunc("/createPayment", objectPaymentController.CreatePayment)
	http.HandleFunc("/getPayment", objectPaymentController.GetPayment)
	http.HandleFunc("/processPayment", objectPaymentController.ProcessPayment)
	http.HandleFunc("/paymentStatus", objectPaymentController.GetPaymentStatus)
	http.HandleFunc("/singleDetail", objectPaymentController.GetSingleDetail)

	//Routes for User
	http.HandleFunc("/getUser", objectUserController.GetUser)
	http.HandleFunc("/addUser", objectUserController.CreateUser)
	http.HandleFunc("/updateUser", objectUserController.UpdateUser)
	http.HandleFunc("/deleteUser", objectUserController.DeleteUser)

	//Routes for Notification
	http.HandleFunc("/createNotification", objectNotificationController.SendPaymentNotification)
	http.HandleFunc("/getNotification", objectNotificationController.GetNotification)

	fmt.Printf("Server started..!, On port no %s", portNumber)

	log.Fatal(http.ListenAndServe(":"+portNumber, nil))

	//fmt.Println(someConfigVar)
}
