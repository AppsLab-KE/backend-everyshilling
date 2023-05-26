package handlers

import (
	"encoding/json"

	"net/http"

	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/models"
	"github.com/AppsLab-KE/backend-everyshilling/services/app-db/internal/core/ports"
)

type TradeHandler struct {
	tradeRepo ports.TradeRepo
}

func NewTradeHandler(tradeRepo ports.TradeRepo) *TradeHandler {
	return &TradeHandler{
		tradeRepo: tradeRepo,
	}
}

func (h *TradeHandler) CreateTrade(w http.ResponseWriter, r *http.Request) {
	var trade models.Trade
	err := json.NewDecoder(r.Body).Decode(&trade)
	if err != nil {
		http.Error(w, "Invalid trade data", http.StatusBadRequest)
		return
	}

	err = h.tradeRepo.CreateTrade(r.Context(), &trade)
	if err != nil {
		http.Error(w, "Failed to create trade", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *TradeHandler) DeleteTrade(w http.ResponseWriter, r *http.Request) {
	tradeID := r.FormValue("tradeID")

	err := h.tradeRepo.DeleteTrade(r.Context(), tradeID)
	if err != nil {
		http.Error(w, "Failed to delete trade", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TradeHandler) UpdateTrade(w http.ResponseWriter, r *http.Request) {
	tradeID := r.FormValue("tradeID")

	var trade models.Trade
	err := json.NewDecoder(r.Body).Decode(&trade)
	if err != nil {
		http.Error(w, "Invalid trade data", http.StatusBadRequest)
		return
	}

	trade.TradeID = tradeID

	err = h.tradeRepo.UpdateTrade(r.Context(), &trade)
	if err != nil {
		http.Error(w, "Failed to update trade", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TradeHandler) GetTrade(w http.ResponseWriter, r *http.Request) {
	tradeID := r.FormValue("tradeID")

	trade, err := h.tradeRepo.GetTrade(r.Context(), tradeID)
	if err != nil {
		http.Error(w, "Failed to retrieve trade", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(trade)
}

func (h *TradeHandler) GetTradeByAccount(w http.ResponseWriter, r *http.Request) {
	accountID := r.FormValue("accountID")

	trades, err := h.tradeRepo.GetTradeByAccount(r.Context(), accountID)
	if err != nil {
		http.Error(w, "Failed to retrieve trades by account", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(trades)
}
