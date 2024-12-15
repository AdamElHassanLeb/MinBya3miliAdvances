package main

import (
	"encoding/json"
	"github.com/AdamElHassanLeb/279MidtermAdamElHassan/API/Internal/Services"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

//	@Summary		Create a new transaction
//	@Description	Create a new transaction between a user and a listing
//	@Tags			transactions
//	@Accept			json
//	@Produce		json
//	@Param			transaction	body		Services.Transaction	true	"Transaction data"
//	@Success		201			{object}	Services.Transaction
//	@Failure		400			{object}	http.ResponseError
//	@Failure		401			{object}	http.ResponseError
//	@Failure		500			{object}	http.ResponseError
//	@Router			/transactions [post]
func (app *application) createTransaction(w http.ResponseWriter, r *http.Request) {

	var transaction Services.Transaction

	//Token Valid

	tokenUserId, ok := r.Context().Value("token_user_id").(int)

	if !ok {
		http.Error(w, "User ID not found in token", http.StatusUnauthorized)
	}

	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, string(err.Error()), http.StatusBadRequest)
		return
	}

	if transaction.UserOfferedID != tokenUserId {
		http.Error(w, string("User offering ID not found in token"), http.StatusUnauthorized)
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

//	@Summary		Get a transaction by ID
//	@Description	Get a specific transaction by its ID
//	@Tags			transactions
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Transaction ID"
//	@Success		200	{object}	Services.Transaction
//	@Failure		400	{object}	http.ResponseError
//	@Failure		404	{object}	http.ResponseError
//	@Failure		500	{object}	http.ResponseError
//	@Router			/transactions/{id} [get]
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

//	@Summary		Get transactions by offered user and status
//	@Description	Get all transactions for a given user and status
//	@Tags			transactions
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		int		true	"User ID"
//	@Param			status	path		string	true	"Status of the transaction"
//	@Success		200		{array}		Services.Transaction
//	@Failure		400		{object}	http.ResponseError
//	@Failure		500		{object}	http.ResponseError
//	@Router			/transactions/offered/{user_id}/{status} [get]
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

//	@Summary		Get transactions by offering user and status
//	@Description	Get all transactions where the user offered the service and status
//	@Tags			transactions
//	@Accept			json
//	@Produce		json
//	@Param			user_id	path		int		true	"User ID"
//	@Param			status	path		string	true	"Status of the transaction"
//	@Success		200		{array}		Services.Transaction
//	@Failure		400		{object}	http.ResponseError
//	@Failure		500		{object}	http.ResponseError
//	@Router			/transactions/offering/{user_id}/{status} [get]
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

//	@Summary		Get transactions by listing ID and status
//	@Description	Get all transactions related to a specific listing and status
//	@Tags			transactions
//	@Accept			json
//	@Produce		json
//	@Param			listing_id	path		int		true	"Listing ID"
//	@Param			status		path		string	true	"Status of the transaction"
//	@Success		200			{array}		Services.Transaction
//	@Failure		400			{object}	http.ResponseError
//	@Failure		500			{object}	http.ResponseError
//	@Router			/transactions/listing/{listing_id}/{status} [get]
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

//	@Summary		Update an existing transaction
//	@Description	Update details of an existing transaction
//	@Tags			transactions
//	@Accept			json
//	@Produce		json
//	@Param			id			path	int						true	"Transaction ID"
//	@Param			transaction	body	Services.Transaction	true	"Updated transaction data"
//	@Success		204
//	@Failure		400	{object}	http.ResponseError
//	@Failure		401	{object}	http.ResponseError
//	@Failure		500	{object}	http.ResponseError
//	@Router			/transactions/{id} [put]
func (app *application) updateTransaction(w http.ResponseWriter, r *http.Request) {
	// Get transaction ID from URL parameter
	transactionIDStr := chi.URLParam(r, "id")
	transactionID, err := strconv.Atoi(transactionIDStr)
	if err != nil {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	// Decode the updated transaction details from the request body
	var transaction Services.Transaction
	err = json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "Failed to decode request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Ensure the user is authorized
	tokenUserID, ok := r.Context().Value("token_user_id").(int)
	if !ok || transaction.UserOfferedID != tokenUserID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Update the transaction
	err = app.Service.Transactions.Update(r.Context(), transactionID, transaction)
	if err != nil {
		http.Error(w, "Failed to update transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusNoContent)
}

//	@Summary		Delete a transaction
//	@Description	Delete a transaction by its ID
//	@Tags			transactions
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"Transaction ID"
//	@Success		204
//	@Failure		400	{object}	http.ResponseError
//	@Failure		401	{object}	http.ResponseError
//	@Failure		404	{object}	http.ResponseError
//	@Failure		500	{object}	http.ResponseError
//	@Router			/transactions/{id} [delete]
func (app *application) deleteTransaction(w http.ResponseWriter, r *http.Request) {
	// Get transaction ID from URL parameter
	transactionIDStr := chi.URLParam(r, "id")
	transactionID, err := strconv.Atoi(transactionIDStr)
	if err != nil {
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	// Ensure the user is authorized
	tokenUserID, ok := r.Context().Value("token_user_id").(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Retrieve the transaction to ensure ownership
	transaction, err := app.Service.Transactions.GetByID(r.Context(), transactionID)
	if err != nil {
		if err.Error() == "transaction not found" {
			http.Error(w, "Transaction not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if transaction.UserOfferedID != tokenUserID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Delete the transaction
	err = app.Service.Transactions.Delete(r.Context(), transactionID)
	if err != nil {
		http.Error(w, "Failed to delete transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusNoContent)
}
