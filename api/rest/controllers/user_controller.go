package controllers

import "net/http"

// Create_user create a new user.
func Create_user(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a new user."))

}

// GetAll_users users search all users.
func GetAll_users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching all users."))

}

// GetById_user search user by ID.
func GetById_user(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching user by ID."))

}

// Update_user updates a user's data.
func Update_user(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user."))

}

// Delete_user delete a user.
func Delete_user(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting a user."))

}
