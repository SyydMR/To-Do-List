package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/SyydMR/To-Do-List/models"
	"github.com/SyydMR/To-Do-List/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("id")

	if userID == nil {
		http.Error(w, "user not authorized", http.StatusUnauthorized)
		return
	}
	Id := userID.(int64)
	newItems := models.GetUserAllItem(Id)
	
	res, _ := json.Marshal(newItems)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("id")
	if userID == nil {
		http.Error(w, "user not authorized", http.StatusUnauthorized)
		return
	}
	Id := userID.(int64)

	user, err := models.GetUserById(Id)
	if err != nil || user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	var newItem models.Item
	err = utils.ParseBody(r, &newItem)
	if err != nil {
		http.Error(w, "error parsing item data", http.StatusBadRequest)
		return
	}

	newItem.UserID = Id
	err = user.AddItem(&newItem)
	if err != nil {
		http.Error(w, "error adding item", http.StatusInternalServerError)
		return
	}

	userItems := models.GetUserAllItem(Id)

	res, err := json.Marshal(userItems)
	if err != nil {
		http.Error(w, "error serializing response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func RemoveItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["ItemId"]
	ID, err := strconv.ParseInt(itemId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	err = models.RemoveItem(ID)
	if err := models.RemoveItem(ID); err != nil {
		http.Error(w, "Failed to remove item", http.StatusInternalServerError)
		return
	}
	userID := r.Context().Value("id")

	if userID == nil {
		http.Error(w, "user not authorized", http.StatusUnauthorized)
		return
	}
	Id := userID.(int64)
	newItems := models.GetUserAllItem(Id)
	res, _ := json.Marshal(newItems)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CheckItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["ItemId"]
	ID, err := strconv.ParseInt(itemId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	itemDetails, db := models.GetItemById(ID)

	if itemDetails.Status == "Done" {
		itemDetails.Status = "To Do"
	} else {
		itemDetails.Status = "Done"
	}

	db.Save(&itemDetails)





	userID := r.Context().Value("id")

	if userID == nil {
		http.Error(w, "user not authorized", http.StatusUnauthorized)
		return
	}
	Id := userID.(int64)
	newItems := models.GetUserAllItem(Id)
	res, _ := json.Marshal(newItems)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var updateItem = &models.Item{}
	utils.ParseBody(r, updateItem)

	vars := mux.Vars(r)
	itemId := vars["ItemId"]
	ID, err := strconv.ParseInt(itemId, 10, 64)

	if err != nil {
		fmt.Println("error while parsing")
	}

	itemDetails, db := models.GetItemById(ID)

	if updateItem.Title != "" {
		itemDetails.Title = updateItem.Title
	}
	if updateItem.Description != "" {
		itemDetails.Description = updateItem.Description
	}
	if updateItem.Status != "" {
		itemDetails.Status = updateItem.Status
	}

	db.Save(&itemDetails)

	userID := r.Context().Value("id")

	if userID == nil {
		http.Error(w, "user not authorized", http.StatusUnauthorized)
		return
	}
	Id := userID.(int64)
	newItems := models.GetUserAllItem(Id)
	res, _ := json.Marshal(newItems)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
