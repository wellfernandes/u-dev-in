package controllers

import (
	"api/database"
	"api/repositories"
	"api/rest/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Create_user create a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a new user."))

	userRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(userRequest, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repositories.NewUserRepository(db)
	userRepo.Create(user)

}

// GetAll_users users search all users.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching all users."))

}

// GetById_user search user by ID.
func GetByIdUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching user by ID."))

}

// Update_user updates a user's data.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user."))

}

// Delete_user delete a user.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting a user."))

}
