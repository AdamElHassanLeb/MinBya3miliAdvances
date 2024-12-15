package main

import (
	"encoding/json"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Services"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

// GetAllListings handles the request to get all listings, optionally filtering by type.
//	@Summary		Get all listings
//	@Description	Retrieve a list of all listings, optionally filtered by type.
//	@Tags			Listings
//	@Param			type	query		string				false	"Listing type"	Enums(offer, request)
//	@Success		200		{array}		Services.Listing	"List of listings"
//	@Failure		500		{string}	string				"Internal Server Error"
//	@Router			/listings [get]
func (app *application) GetAllListings(w http.ResponseWriter, r *http.Request) {
	// Extract optional 'type' query parameter
	listingType := chi.URLParam(r, "type")

	// Call the service method to get all listings, optionally filtered by type
	listings, err := app.Service.Listings.GetAll(r.Context(), listingType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(listings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetListingByID handles the request to get a listing by its ID.
//	@Summary		Get a listing by ID
//	@Description	Retrieve a listing by its unique ID.
//	@Tags			Listings
//	@Param			id	path		int					true	"Listing ID"
//	@Success		200	{object}	Services.Listing	"Detailed listing"
//	@Failure		400	{string}	string				"Invalid listing ID"
//	@Failure		404	{string}	string				"Listing not found"
//	@Failure		500	{string}	string				"Internal Server Error"
//	@Router			/listings/{id} [get]
func (app *application) GetListingByID(w http.ResponseWriter, r *http.Request) {
	// Extract listing ID from the URL
	listingIDStr := chi.URLParam(r, "id")
	listingID, err := strconv.Atoi(listingIDStr)
	if err != nil {
		http.Error(w, "Invalid listing ID", http.StatusBadRequest)
		return
	}

	// Call service to get the listing by ID
	listing, err := app.Service.Listings.GetByID(r.Context(), listingID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if listing.ListingID == 0 {
		http.Error(w, "Listing not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(listing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetListingsByUserID handles getting listings by user ID and type.
//	@Summary		Get listings by user ID
//	@Description	Retrieve listings created by a specific user, optionally filtered by type.
//	@Tags			Listings
//	@Param			user_id	path		int					true	"User ID"
//	@Param			type	path		string				true	"Listing type"	Enums(offer, request)
//	@Success		200		{array}		Services.Listing	"List of listings"
//	@Failure		400		{string}	string				"Invalid user ID"
//	@Failure		500		{string}	string				"Internal Server Error"
//	@Router			/listings/user/{user_id}/{type} [get]
func (app *application) GetListingsByUserID(w http.ResponseWriter, r *http.Request) {
	// Extract userID and listingType from the URL
	userIDStr := chi.URLParam(r, "user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	listingType := chi.URLParam(r, "type")

	// Call the service to get listings by user ID
	listings, err := app.Service.Listings.GetByUserID(r.Context(), userID, listingType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(listings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// CreateListing handles the request to create a new listing.
//	@Summary		Create a new listing
//	@Description	Create a new listing for the authenticated user.
//	@Tags			Listings
//	@Param			listing	body		Services.Listing	true	"Listing data"
//	@Success		201		{object}	Services.Listing	"Created listing"
//	@Failure		400		{string}	string				"Invalid input"
//	@Failure		401		{string}	string				"Unauthorized"
//	@Failure		500		{string}	string				"Internal Server Error"
//	@Router			/listings [post]
func (app *application) CreateListing(w http.ResponseWriter, r *http.Request) {
	var listing Services.Listing

	tokenUserId, ok := r.Context().Value("token_user_id").(int)

	if !ok {
		http.Error(w, "User ID not found in token", http.StatusUnauthorized)
	}

	// Decode the JSON request body into the Listing struct
	err := json.NewDecoder(r.Body).Decode(&listing)
	if err != nil {
		http.Error(w, string(err.Error()), http.StatusBadRequest)
		return
	}

	if tokenUserId != listing.UserID {
		http.Error(w, "Cannot Create Listing For Another User", http.StatusUnauthorized)
	}

	// Call the service to create the listing
	createdListing, err := app.Service.Listings.Create(r.Context(), &listing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdListing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateListing handles the request to update an existing listing.
//	@Summary		Update a listing
//	@Description	Update an existing listing created by the authenticated user.
//	@Tags			Listings
//	@Param			id		path		int					true	"Listing ID"
//	@Param			listing	body		Services.Listing	true	"Updated listing data"
//	@Success		200		{string}	string				"Listing updated successfully"
//	@Failure		400		{string}	string				"Invalid input"
//	@Failure		401		{string}	string				"Unauthorized"
//	@Failure		500		{string}	string				"Internal Server Error"
//	@Router			/listings/{id} [put]
func (app *application) UpdateListing(w http.ResponseWriter, r *http.Request) {
	var listing Services.Listing

	tokenUserId, ok := r.Context().Value("token_user_id").(int)

	if !ok {
		http.Error(w, "User ID not found in token", http.StatusUnauthorized)
	}

	// Extract listing ID from the URL
	listingIDStr := chi.URLParam(r, "id")
	listingID, err := strconv.Atoi(listingIDStr)
	if err != nil {
		http.Error(w, "Invalid listing ID", http.StatusBadRequest)
		return
	}

	DbListing, err := app.Service.Listings.GetByID(r.Context(), listingID)
	if err != nil {
		http.Error(w, "Invalid listing ID", http.StatusBadRequest)
		return
	}

	if DbListing.UserID != tokenUserId {
		http.Error(w, "Cannot Delete Listing For Another User", http.StatusUnauthorized)
	}

	// Decode the JSON request body into the Listing struct
	err = json.NewDecoder(r.Body).Decode(&listing)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Call the service to update the listing
	err = app.Service.Listings.Update(r.Context(), &listing, listingID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Listing updated successfully"))
}

// DeleteListing handles the HTTP request to delete a listing by ID.
//	@Summary		Delete a listing
//	@Description	Delete a listing by its ID, only if the user is the creator.
//	@Tags			Listings
//	@Param			id	path		int		true	"Listing ID"
//	@Success		204	{string}	string	"No content"
//	@Failure		400	{string}	string	"Invalid listing ID"
//	@Failure		401	{string}	string	"Unauthorized"
//	@Failure		500	{string}	string	"Internal Server Error"
//	@Router			/listings/{id} [delete]
func (app *application) DeleteListing(w http.ResponseWriter, r *http.Request) {
	// Extract listing ID from the URL
	listingIDStr := chi.URLParam(r, "id")
	listingID, err := strconv.Atoi(listingIDStr)
	if err != nil {
		http.Error(w, "Invalid listing ID", http.StatusBadRequest)
		return
	}

	tokenUserId, ok := r.Context().Value("token_user_id").(int)

	if !ok {
		http.Error(w, "User ID not found in token", http.StatusUnauthorized)
	}

	DbListing, err := app.Service.Listings.GetByID(r.Context(), listingID)
	if err != nil {
		http.Error(w, "Invalid listing ID", http.StatusBadRequest)
		return
	}

	if DbListing.UserID != tokenUserId {
		http.Error(w, "Cannot Delete Listing For Another User", http.StatusUnauthorized)
	}

	// Call the service to delete the listing
	err = app.Service.Listings.Delete(r.Context(), listingID)
	if err != nil {
		http.Error(w, "Failed to delete listing", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content for successful deletion
}

// GetListingsBySearch handles the HTTP request to get listings by search query and type.
//	@Summary		Get listings by search query and type
//	@Description	Retrieves a list of listings that match a search query and a specific listing type (e.g., offer or request).
//	@Tags			Listings
//	@Param			query	path		string	true	"Search query"
//	@Param			type	path		string	true	"Type of the listing (e.g., offer or request)"
//	@Success		200		{array}		Listing	"List of listings matching the search query"
//	@Failure		400		{string}	string	"Invalid search query or listing type"
//	@Failure		500		{string}	string	"Internal server error"
//	@Router			/listings/search/{query}/{type} [get]
func (app *application) GetListingsBySearch(w http.ResponseWriter, r *http.Request) {
	query := chi.URLParam(r, "query")
	listingType := chi.URLParam(r, "type")

	// Call the service to get listings by search
	listings, err := app.Service.Listings.GetBySearch(r.Context(), query, listingType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(listings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetListingsByDistance handles the HTTP request to get listings by location and distance.
//	@Summary		Get listings by location and distance
//	@Description	Retrieves a list of listings within a specified distance from a given location (latitude and longitude).
//	@Tags			Listings
//	@Param			latitude		path		float64	true	"Latitude of the location"
//	@Param			longitude		path		float64	true	"Longitude of the location"
//	@Param			max_distance	path		float64	true	"Maximum distance (in kilometers) to search listings within"
//	@Param			type			path		string	true	"Type of the listing (e.g., offer or request)"
//	@Success		200				{array}		Listing	"List of listings"
//	@Failure		400				{string}	string	"Invalid parameters"
//	@Failure		500				{string}	string	"Internal server error"
//	@Router			/listings/distance/{latitude}/{longitude}/{max_distance}/{type} [get]
func (app *application) GetListingsByDistance(w http.ResponseWriter, r *http.Request) {
	latitude, err := strconv.ParseFloat(chi.URLParam(r, "latitude"), 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}
	longitude, err := strconv.ParseFloat(chi.URLParam(r, "longitude"), 64)
	if err != nil {
		http.Error(w, "Invalid longitude", http.StatusBadRequest)
		return
	}
	maxDistance, err := strconv.ParseFloat(chi.URLParam(r, "max_distance"), 64)
	if err != nil {
		http.Error(w, "Invalid maxDistance", http.StatusBadRequest)
		return
	}
	listingType := chi.URLParam(r, "type")

	// Call the service to get listings by distance
	listings, err := app.Service.Listings.GetByDistance(r.Context(), latitude, longitude, maxDistance, listingType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(listings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetListingsByDistanceAndSearch handles the HTTP request to get listings by location, distance, and search query.
//	@Summary		Get listings by location, distance, and search query
//	@Description	Retrieves a list of listings within a specified distance and matching a search query.
//	@Tags			Listings
//	@Param			latitude		path		float64	true	"Latitude of the location"
//	@Param			longitude		path		float64	true	"Longitude of the location"
//	@Param			max_distance	path		float64	true	"Maximum distance (in kilometers) to search listings within"
//	@Param			type			path		string	true	"Type of the listing (e.g., offer or request)"
//	@Param			query			path		string	true	"Search query"
//	@Success		200				{array}		Listing	"List of listings matching the search query"
//	@Failure		400				{string}	string	"Invalid parameters"
//	@Failure		500				{string}	string	"Internal server error"
//	@Router			/listings/distance-search/{latitude}/{longitude}/{max_distance}/{type}/{query} [get]
func (app *application) GetListingsByDistanceAndSearch(w http.ResponseWriter, r *http.Request) {
	latitude, err := strconv.ParseFloat(chi.URLParam(r, "latitude"), 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}
	longitude, err := strconv.ParseFloat(chi.URLParam(r, "longitude"), 64)
	if err != nil {
		http.Error(w, "Invalid longitude", http.StatusBadRequest)
		return
	}
	maxDistance, err := strconv.ParseFloat(chi.URLParam(r, "max_distance"), 64)
	if err != nil {
		http.Error(w, "Invalid maxDistance", http.StatusBadRequest)
		return
	}
	listingType := chi.URLParam(r, "type")

	query := chi.URLParam(r, "query")

	// Call the service to get listings by distance
	listings, err := app.Service.Listings.GetByDistanceAndSearch(r.Context(), latitude, longitude, maxDistance, listingType, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(listings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetListingsByDate handles the HTTP request to get listings by date created, sorted descending.
//	@Summary		Get listings by date created, sorted descending
//	@Description	Retrieves a list of listings sorted by creation date in descending order.
//	@Tags			Listings
//	@Param			type	path		string	true	"Type of the listing (e.g., offer or request)"
//	@Success		200		{array}		Listing	"List of listings sorted by creation date"
//	@Failure		500		{string}	string	"Internal server error"
//	@Router			/listings/date/{type} [get]
func (app *application) GetListingsByDate(w http.ResponseWriter, r *http.Request) {
	listingType := chi.URLParam(r, "type")

	// Call the service to get listings by date created descending
	listings, err := app.Service.Listings.GetByDateCreatedDescending(r.Context(), listingType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(listings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetListingsByDateAndSearch handles the HTTP request to get listings by date and search query.
//	@Summary		Get listings by date and search query
//	@Description	Retrieves a list of listings sorted by creation date in descending order that match a search query.
//	@Tags			Listings
//	@Param			query	path		string	true	"Search query"
//	@Param			type	path		string	true	"Type of the listing (e.g., offer or request)"
//	@Success		200		{array}		Listing	"List of listings matching the search query"
//	@Failure		400		{string}	string	"Invalid parameters"
//	@Failure		500		{string}	string	"Internal server error"
//	@Router			/listings/date-search/{query}/{type} [get]
func (app *application) GetListingsByDateAndSearch(w http.ResponseWriter, r *http.Request) {
	query := chi.URLParam(r, "query")
	listingType := chi.URLParam(r, "type")

	// Call the service to get listings by date created and search query
	listings, err := app.Service.Listings.GetByDateCreatedAndSearchDescending(r.Context(), query, listingType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(listings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
