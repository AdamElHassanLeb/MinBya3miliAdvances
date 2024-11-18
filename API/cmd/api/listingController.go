package main

import (
	"encoding/json"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Services"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

// GetAllListings handles the request to get all listings, optionally filtering by type.
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
	err = app.Service.Listings.Create(r.Context(), &listing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(listing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateListing handles the request to update an existing listing.
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

// GetListingsBySearch handles the request to get listings by a search query and type.
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

// GetListingsByDistance handles the request to get listings by location and distance.
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

// GetListingsByDistanceAndSearch handles the request to get listings by location and distance.
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

/*
// GetListingsByLocation handles the request to get listings by location and range.
func (app *application) GetListingsByLocation(w http.ResponseWriter, r *http.Request) {
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
	maxRange, err := strconv.ParseFloat(chi.URLParam(r, "max_range"), 64)
	if err != nil {
		http.Error(w, "Invalid maxRange", http.StatusBadRequest)
		return
	}
	listingType := chi.URLParam(r, "type")

	// Call the service to query listings by location
	listings, err := app.Service.Listings.QueryByLocation(r.Context(), latitude, longitude, maxRange, listingType)
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
*/

// GetListingsByDate handles the request to get listings by date, sorted descending.
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

// GetListingsByDateAndSearch handles the request to get listings by date and search query.
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
