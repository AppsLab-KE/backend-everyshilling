package handlers

import (
	"encoding/json"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/ports"
	"net/http"
)

type AccountHandler struct {
	accountRepo ports.AccountRepo
}

func NewAccountHandler(accountRepo ports.AccountRepo) *AccountHandler {
	return &AccountHandler{
		accountRepo: accountRepo,
	}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, "Invalid account data", http.StatusBadRequest)
		return
	}

	err, _ = h.accountRepo.CreateAccount(r.Context(), &account)
	if err != nil {
		http.Error(w, "Failed to create account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *AccountHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	accountID := r.FormValue("accountID")

	err := h.accountRepo.DeleteAccount(r.Context(), accountID)
	if err != nil {
		http.Error(w, "Failed to delete account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *AccountHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	accountID := r.FormValue("accountID")

	var account models.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, "Invalid account data", http.StatusBadRequest)
		return
	}

	account.AccountID = accountID

	err, _ = h.accountRepo.UpdateAccount(r.Context(), &account)
	if err != nil {
		http.Error(w, "Failed to update account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	accountID := r.FormValue("accountID")

	account, err := h.accountRepo.GetAccount(r.Context(), accountID)
	if err != nil {
		http.Error(w, "Failed to retrieve account", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(account)
}

func (h *AccountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := h.accountRepo.GetAccounts(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve accounts", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(accounts)
}
