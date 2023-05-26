package handlers

import (
	"encoding/json"
	"errors"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/ports"
	"net/http"
)

type TransactionHandler struct {
	transactionRepo ports.TransactionRepo
}

func NewTransactionHandler(transactionRepo ports.TransactionRepo) *TransactionHandler {
	return &TransactionHandler{
		transactionRepo: transactionRepo,
	}
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "Invalid transaction data", http.StatusBadRequest)
		return
	}

	err = h.transactionRepo.CreateTransaction(r.Context(), &transaction)
	if err != nil {
		http.Error(w, "Failed to create transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *TransactionHandler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	transactionID := r.FormValue("transactionID")

	err := h.transactionRepo.DeleteTransaction(r.Context(), transactionID)
	if err != nil {
		http.Error(w, "Failed to delete transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TransactionHandler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	transactionID := r.FormValue("transactionID")

	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "Invalid transaction data", http.StatusBadRequest)
		return
	}

	if transaction.TransactionId != transactionID {
		http.Error(w, "Mismatching transaction ID", http.StatusBadRequest)
		return
	}

	err = h.transactionRepo.UpdateTransaction(r.Context(), &transaction)
	if err != nil {
		http.Error(w, "Failed to update transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TransactionHandler) GetTransaction(w http.ResponseWriter, r *http.Request) {
	transactionID := r.FormValue("transactionID")

	transaction, err := h.transactionRepo.GetTransaction(r.Context(), transactionID)
	if err != nil {
		if errors.Is(err, ports.ErrNotFound) {
			http.Error(w, "Transaction not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve transaction", http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(transaction)
}

func (h *TransactionHandler) GetTransactionsByAccount(w http.ResponseWriter, r *http.Request) {
	accountID := r.FormValue("accountID")

	transactions, err := h.transactionRepo.GetTransactionByAccount(r.Context(), accountID)
	if err != nil {
		http.Error(w, "Failed to retrieve transactions by account", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(transactions)
}
