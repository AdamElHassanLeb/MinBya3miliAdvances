package main

import (
	"encoding/json"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Services"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (app *application) createTransaction(w http.ResponseWriter, r *http.Request) {

	var transaction Services.Transaction

	//Token Valid

	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, string(err.Error()), http.StatusBadRequest)
		return
	}

	createdTransaction, err := app.Service.Transactions.Create(r.Context(), &transaction)
	if err != nil {
		http.Error(w, string(err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdTransaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) getTransactionByID(w http.ResponseWriter, r *http.Request) {
	// Get transaction ID from URL parameter
	transactionIDStr := chi.URLParam(r, "id")
	transactionID, err := strconv.Atoi(transactionIDStr)
	if err != nil {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	// Retrieve the transaction by ID
	transaction, err := app.Service.Transactions.GetByID(r.Context(), transactionID)
	if err != nil {
		if err.Error() == "transaction not found" {
			http.Error(w, "Transaction not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the transaction
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) getTransactionsByOfferedUserAndStatus(w http.ResponseWriter, r *http.Request) {
	// Get user ID and optional status from query parameters
	userIDStr := chi.URLParam(r, "user_id")
	status := chi.URLParam(r, "status")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Retrieve transactions by user and status
	transactions, err := app.Service.Transactions.GetByOfferedUserAndStatus(r.Context(), userID, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the transactions
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) getTransactionsByOfferingUserAndStatus(w http.ResponseWriter, r *http.Request) {
	// Get user ID and optional status from query parameters
	userIDStr := chi.URLParam(r, "user_id")
	status := chi.URLParam(r, "status")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID "+err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve transactions by user and status
	transactions, err := app.Service.Transactions.GetByOfferingUserAndStatus(r.Context(), userID, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the transactions
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) getTransactionsByListingAndStatus(w http.ResponseWriter, r *http.Request) {
	// Get listing ID and optional status from query parameters
	listingIDStr := chi.URLParam(r, "listing_id")
	status := chi.URLParam(r, "status")

	listingID, err := strconv.Atoi(listingIDStr)
	if err != nil {
		http.Error(w, "Invalid listing ID", http.StatusBadRequest)
		return
	}

	// Retrieve transactions by listing and status
	transactions, err := app.Service.Transactions.GetByListingAndStatus(r.Context(), listingID, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the transactions
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
