package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sgrkabadi96/bookAppClone/config"
	"github.com/sgrkabadi96/bookAppClone/models"
	"github.com/sgrkabadi96/bookAppClone/utils"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", " application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	IdString := vars["id"]
	id, err := strconv.ParseInt(IdString, 0, 0)

	if err != nil {
		panic(err)
	}
	user := models.GetUserById(int(id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(&user)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	NewUser := &models.User{}
	utils.ParseBody(r, NewUser)
	NewUser.CreateUser()
	res, _ := json.Marshal(NewUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	idStr := vars["id"]
	id, _ := strconv.ParseInt(idStr, 0, 0)
	user := models.DeleteById(int(id))
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAllUser(w http.ResponseWriter, r *http.Request) {
	users := models.DeleteAllUser()
	res, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	UserUpdate := &models.User{}

	utils.ParseBody(r, UserUpdate)

	vars := mux.Vars(r)
	id := vars["id"]
	idInt, _ := strconv.ParseInt(id, 0, 0)
	OldUser := models.GetUserById(int(idInt))

	if UserUpdate.Name != "" {
		OldUser.Name = UserUpdate.Name
	}

	if UserUpdate.Age != 0 {
		OldUser.Age = UserUpdate.Age
	}
	if UserUpdate.Website != "" {
		OldUser.Age = UserUpdate.Age
	}
	db := config.GetDB()
	db.Save(&OldUser)

	w.Header().Set("Content-Type", " application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(OldUser)
	w.Write(res)
}
