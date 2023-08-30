package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	UserId    int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdat"`
	UpdatedAt string `json:"updatedat"`
}

var UserDetails = make([]User, 0)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "Application/json")
		err := json.NewEncoder(w).Encode(UserDetails)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)

	} else {
		w.WriteHeader(http.StatusBadGateway)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newUser User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newUser.CreatedAt = time.Now().Format("02-01-2006T15:04:05")
		newUser.UpdatedAt = time.Now().Format("02-01-2006T15:04:05")

		UserDetails = append(UserDetails, newUser)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, newUser)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		var updatedUser User
		err := json.NewDecoder(r.Body).Decode(&updatedUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for i, user := range UserDetails {
			if user.UserId == updatedUser.UserId {
				updatedUser.CreatedAt = UserDetails[i].CreatedAt
				updatedUser.UpdatedAt = time.Now().Format("02-01-2006T15:04:05")
				UserDetails[i] = updatedUser
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, updatedUser)
			}
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		userIdStr := r.URL.Query().Get("id")
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for i, user := range UserDetails {
			if user.UserId == userId {
				UserDetails = append(UserDetails[:i], UserDetails[i+1:]...)
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, "User deleted successfully!")
			}
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func getSingleUserDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		userIdStr := r.URL.Query().Get("id")
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			http.Error(w, "Invalid User ID", http.StatusBadRequest)
		}

		var userFound *User
		for _, user := range UserDetails {
			if user.UserId == userId {
				userFound = &user
				w.Header().Set("Content-Type", "Aplication/json")

				err = json.NewEncoder(w).Encode(user)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				break

			}
		}
		if userFound != nil {
			w.WriteHeader(http.StatusCreated)

		} else {
			http.Error(w, "Payment not found", http.StatusNotFound)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {

	http.HandleFunc("/getUser", Index)
	http.HandleFunc("/addUser", createUser)
	http.HandleFunc("/updateUser", updateUser)
	http.HandleFunc("/deleteUser", deleteUser)
	http.HandleFunc("/singleDetail", getSingleUserDetail)
	fmt.Println("Server started..!")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
