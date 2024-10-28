package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SyydMR/To-Do-List/models"
	"github.com/SyydMR/To-Do-List/utils"
	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	newUsers := models.GetAllUser()
	res, err := json.Marshal(newUsers)
	if err != nil {
		http.Error(w, "Failed to parse users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing")
	}

	newUser, err := models.GetUserById(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	
	res, err := json.Marshal(newUser)
	if err != nil {
		http.Error(w, "Failed to parse users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    createUser := &models.User{}

    if err := utils.ParseBody(r, createUser); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    registeredUser, err := createUser.Register()
    if err != nil {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    res, err := json.Marshal(registeredUser)
    if err != nil {
        http.Error(w, "Error creating response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write(res)
}


func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := utils.ParseBody(r, &loginData); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user := models.User{Username: loginData.Username}

	token, err := user.Login(loginData.Password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	response, err := json.Marshal(map[string]string{"token": token})
	if err != nil {
		http.Error(w, "Error creating response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}