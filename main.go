package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"paymentsystem/config"
	"paymentsystem/controller"
	"paymentsystem/database"
	"paymentsystem/service"

	firestore "cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

func main() {

	config := config.InitConfig()
	if config == nil {
		// if any error occurs program will not run
		return
	}

	// create connection object here and send it to all the controller
	ctx := context.Background()
	projectId := "paymentsystem-4337a"
	options := option.WithCredentialsFile("paymentsystem-4337a-firebase-adminsdk-zwgan-b6fdcc7afd.json")
	connection, err := firestore.NewClient(ctx, projectId, options)
	if err != nil {
		fmt.Println("Error ", err)
		return
	}

	portNumber := config.String("portNumber")

	objectPaymentController := controller.PaymentController{
		Service: service.PaymentService{
			Database: database.PaymentDataBase{
				ConnectionObjectClient: connection,
			},
		},
	}

	objectUserController := controller.UserController{
		Service: service.UserService{
			Database: database.UserDataBase{
				ConnectionObjectClient: connection,
			},
		},
	}
	objectNotificationController := controller.NotificationController{
		Service: service.NotificationService{
			Database: database.NotificationDataBase{
				ConnectionObjectClient: connection,
			},
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
