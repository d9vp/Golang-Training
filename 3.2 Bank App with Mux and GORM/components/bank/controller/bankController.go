package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user/components/bank/service"
	"user/models"

	"github.com/gorilla/mux"
)

func NewBankHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		FullName     string `json:"fullName"`
		Abbreviation string `json:"abbreviation"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bank := &models.Bank{
		FullName:     req.FullName,
		Abbreviation: req.Abbreviation,
		Accounts:     []*models.Account{},
		LedgerData:   []*models.LedgerData{},
	}

	bank, err := service.CreateBank(bank)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bank)
}

func GetBanksHandler(w http.ResponseWriter, r *http.Request) {
	banks, err := service.GetAllBanks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(banks)
}

func UpdateBankHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bankIDStr := vars["id"]
	bankID, err := strconv.Atoi(bankIDStr)

	if err != nil {
		http.Error(w, "Invalid bank ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Parameter string `json:"parameter"`
		NewValue  string `json:"newValue"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.UpdateBank(bankID, req.Parameter, req.NewValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Bank updated successfully"))
}

func DeleteBankHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bankID := vars["id"]

	bankIDint, ok := strconv.Atoi(bankID)
	if ok != nil {
		http.Error(w, "invalid type for user id", http.StatusBadRequest)
		return
	}

	err := service.RemoveBank(bankIDint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
