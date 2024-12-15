package main

import (
	"encoding/json"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Services"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

// GetAllUsers godoc
//	@Summary		Get all users
//	@Description	Retrieve a list of all users in the system
//	@Tags			users
//	@Produce		json
//	@Success		200	{array}		Services.User
//	@Failure		500	{object}	string	"Internal Server Error"
//	@Router			/user/users [get]
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

// GetUserById godoc
//	@Summary		Get user by ID
//	@Description	Retrieve user information by their unique ID
//	@Tags			users
//	@Param			id	path	int	true	"User ID"
//	@Produce		json
//	@Success		200	{object}	Services.User
//	@Failure		400	{object}	string	"Bad Request"
//	@Failure		404	{object}	string	"User not found"
//	@Failure		500	{object}	string	"Internal Server Error"
//	@Router			/user/userId/{id} [get]
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

// GetUserByName godoc
//	@Summary		Get users by name
//	@Description	Retrieve users with matching name
//	@Tags			users
//	@Param			name	path	string	true	"User Name"
//	@Produce		json
//	@Success		200	{array}		Services.User
//	@Failure		404	{object}	string	"No users found"
//	@Failure		500	{object}	string	"Internal Server Error"
//	@Router			/user/userName/{name} [get]
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

// CreateUser godoc
//	@Summary		Create a new user
//	@Description	Add a new user to the system
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		Services.User	true	"User data"
//	@Success		201		{object}	Services.User
//	@Failure		400		{object}	string	"Bad Request"
//	@Failure		500		{object}	string	"Internal Server Error"
//	@Router			/user/create [post]
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

// DeleteUser godoc
//	@Summary		Delete a user
//	@Description	Remove a user from the system by their ID
//	@Tags			users
//	@Param			id	path	int	true	"User ID"
//	@Security		BearerAuth
//	@Success		204	"No Content"
//	@Failure		400	{object}	string	"Bad Request"
//	@Failure		401	{object}	string	"Unauthorized"
//	@Failure		404	{object}	string	"User not found"
//	@Failure		500	{object}	string	"Internal Server Error"
//	@Router			/user/delete/{id} [delete]
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

// UpdateUser godoc
//	@Summary		Update a user
//	@Description	Update user information by their ID
//	@Tags			users
//	@Param			id		path	int				true	"User ID"
//	@Param			user	body	Services.User	true	"Updated user data"
//	@Security		BearerAuth
//	@Success		200	"User updated successfully"
//	@Failure		400	{object}	string	"Bad Request"
//	@Failure		401	{object}	string	"Unauthorized"
//	@Failure		500	{object}	string	"Internal Server Error"
//	@Router			/user/update/{id} [put]
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

// authUser godoc
//	@Summary		Authenticate user
//	@Description	Generate a JWT token for the user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body		object					true	"User credentials"
//	@Success		200			{object}	map[string]interface{}	"Token and user data"
//	@Failure		400			{object}	string					"Bad Request"
//	@Failure		401			{object}	string					"Unauthorized"
//	@Failure		500			{object}	string					"Internal Server Error"
//	@Router			/user/auth [post]
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
