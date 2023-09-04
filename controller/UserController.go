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

var UserDetails = make([]entity.User, 0)

type UserController struct {
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	util.CheckMethod(r, w, http.MethodGet)

	w.Header().Set("Content-Type", "Application/json")
	err := json.NewEncoder(w).Encode(UserDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	util.CheckMethod(r, w, http.MethodPost)

	var newUser entity.User
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
}

func (uc UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	util.CheckMethod(r, w, http.MethodPut)

	var updatedUser entity.User
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

}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	util.CheckMethod(r, w, http.MethodDelete)

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

}

func (uc UserController) GetSingleUserDetail(w http.ResponseWriter, r *http.Request) {
	util.CheckMethod(r, w, http.MethodGet)
	userIdStr := r.URL.Query().Get("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
	}

	var userFound *entity.User
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
}
