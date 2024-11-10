package Controllers

import (
	"encoding/json"
	"net/http"
)

// GetAllUsers handles the request to get all users.
func (app *Application) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Call the service method to get all users.
	users, err := app.store.Users.getAll(r.Context())
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
