package main

import (
	"encoding/json"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Services"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

// GetAllUsers handles the request to get all users.
func (app *application) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Call the service method to get all users.
	users, err := app.Service.Users.GetAll(r.Context())
	if err != nil {
		// If there's an error, respond with an internal server error status and message.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the appropriate response header.
	w.Header().Set("Content-Type", "application/json")

	// Encode the users slice into JSON and send it as the response.
	if err := json.NewEncoder(w).Encode(users); err != nil {
		// If encoding fails, respond with an internal server error status and message.
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) GetUserById(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the URL (e.g., /user/123)
	userId := chi.URLParam(r, "id")

	// Convert the userId from string to int
	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Call the GetById method from the UserService
	user, err := app.Service.Users.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.UserID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Return the user as JSON response
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *application) GetUserByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	// Call the service to get users by name
	users, err := app.Service.Users.GetByName(r.Context(), name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the list of users
	if len(users) == 0 {
		http.Error(w, "No users found", http.StatusNotFound)
		return
	}

	// Return the users in JSON format
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *application) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user Services.User

	// Decode the JSON request body into the User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the service to create the user
	err = app.Service.Users.Create(r.Context(), &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the created user as a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// DeleteUser handles the HTTP request to delete a user by ID.
func (app *application) DeleteUser(w http.ResponseWriter, r *http.Request) {

	// Retrieve the user_id from the request context
	tokenUserID, ok := r.Context().Value("token_user_id").(int)

	if !ok {
		http.Error(w, "User ID not found in token", http.StatusUnauthorized)
		return
	}

	// Extract the user ID from the URL parameters
	userIDStr := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if tokenUserID != userID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Call the service to delete the user and get a confirmation
	deleted, err := app.Service.Users.Delete(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	if !deleted {
		http.Error(w, "User not found or already deleted", http.StatusNotFound)
		return
	}

	// Respond with success if the user was deleted
	w.WriteHeader(http.StatusNoContent) // 204 No Content for a successful delete with no body
}

// UpdateUser handles the HTTP request to update a user's information.
func (app *application) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user Services.User

	// Retrieve the user_id from the request context
	tokenUserID, ok := r.Context().Value("token_user_id").(int)

	// Extract the user ID from the URL parameters
	userIDStr := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if tokenUserID != userID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if !ok {
		http.Error(w, "User ID not found in token", http.StatusUnauthorized)
		return
	}

	// Parse the request body into the User struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user.UserID = userID

	// Call the service to update the user
	err = app.Service.Users.Update(r.Context(), &user)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
}

func (app *application) authUser(w http.ResponseWriter, r *http.Request) {
	var authRequest struct {
		PhoneNumber string `json:"phone_number"`
		Password    string `json:"password"`
	}

	// Decode the JSON request body
	err := json.NewDecoder(r.Body).Decode(&authRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call the Auth function from UserService to generate the token
	token, user, err := app.Service.Users.Auth(r.Context(), authRequest.PhoneNumber, authRequest.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Return the JWT token
	w.Header().Set("Content-Type", "application/json")

	// Create a response struct to hold the token and user data
	response := struct {
		Token string        `json:"token"`
		User  Services.User `json:"user"`
	}{
		Token: token,
		User:  user,
	}

	// Encode the response struct to JSON and write it to the response writer
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
