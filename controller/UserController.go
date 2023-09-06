package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"paymentsystem/entity"
	"paymentsystem/service"
	util "paymentsystem/utility"
	"strconv"
)

var UserDetails = make([]entity.User, 0)

type UserController struct {
	Service service.UserService
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	util.CheckMethod(r, w, http.MethodGet)

	w.Header().Set("Content-Type", "Application/json")
	userDetails := uc.Service.GetUser()
	err := json.NewEncoder(w).Encode(userDetails)
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

	newUser = uc.Service.CreateUser(newUser)
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

	newUpdatedUser, err := uc.Service.UpdateUser(updatedUser)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, newUpdatedUser)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	util.CheckMethod(r, w, http.MethodDelete)

	userIdStr := r.URL.Query().Get("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = uc.Service.DeleteUser(userId)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "User deleted successfully!")
}
