package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user/components/bank/service"

	"github.com/gorilla/mux"
)

func AddToLedger(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var request struct {
		CorrBankID int     `json:"corrBankId"`
		Amount     float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	bankIDint, ok := strconv.Atoi(vars["id"])
	if ok != nil {
		http.Error(w, "Invalid bank id", http.StatusBadRequest)
		return
	}
	err := service.AddToLedger(bankIDint, request.CorrBankID, request.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetLedger(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bankIDint, ok := strconv.Atoi(vars["id"])
	if ok != nil {
		http.Error(w, "Invalid bank id", http.StatusBadRequest)
		return
	}
	// var ledger []bankService.LedgerData
	ledger, err := service.GetLedger(bankIDint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(ledger)
	w.WriteHeader(http.StatusOK)
}
