package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	accService "user/components/accounts/service"
	"user/components/user/service"
	"user/models"

	"github.com/gorilla/mux"
)

func GetUserByUserId(userId string) (*models.User, error) {
	userIdint, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}
	users, err := service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if int(user.ID) == userIdint {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func NewAdminHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		UserName  string `json:"userName"`
		Password  string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if req.FirstName == "" || req.LastName == "" || req.Password == "" || req.UserName == "" {
		http.Error(w, "Cannot pass empty string as a parameter", http.StatusBadRequest)
		return
	}

	user, err := service.NewAdmin(req.FirstName, req.LastName, req.UserName, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func NewUserHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		UserName  string `json:"userName"`
		Password  string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if req.FirstName == "" || req.LastName == "" || req.Password == "" || req.UserName == "" {
		http.Error(w, "Cannot pass empty string as a parameter", http.StatusBadRequest)
		return
	}

	user, err := service.NewUser(req.FirstName, req.LastName, req.UserName, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := service.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	user, err := GetUserByUserId(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusOK)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	user, err := GetUserByUserId(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req struct {
		Parameter string      `json:"parameter"`
		NewValue  interface{} `json:"newValue"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = service.UpdateUsers(user.UserName, req.Parameter, req.NewValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	user, err := GetUserByUserId(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.DeleteUsers(user.UserName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func NewAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	user, err := GetUserByUserId(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req struct {
		BankID  int     `json:"bankId"`
		Balance float64 `json:"balance"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Balance < 1000 {
		http.Error(w, "Initial balance must be at least Rs. 1000", http.StatusBadRequest)
		return
	}

	account, err := accService.CreateAccount(user, req.BankID, req.Balance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func GetAccountsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	user, err := accService.GetUserWithAccounts(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user.Accounts)
}

func GetTotalBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	user, err := accService.GetUserWithAccounts(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	totalBalance, err := accService.GetTotalBalanceForUser(user.UserName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]float64{"Total Balance for user": totalBalance})
}

func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	accountID, _ := strconv.Atoi(vars["accId"])

	user, err := accService.GetUserWithAccounts(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := accService.GetAccountByID(user, accountID)
	if err != nil {
		http.Error(w, "account not found or doesn't belong to the user", http.StatusBadRequest)
		return
	}

	if err := accService.DeleteAccount(account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "account deactivated successfully"})
}

func DepositHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	accountID, _ := strconv.Atoi(vars["accId"])

	var req struct {
		Amount float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := accService.GetUserWithAccounts(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := accService.GetAccountByID(user, accountID)
	if err != nil {
		http.Error(w, "account not found or doesn't belong to the user", http.StatusBadRequest)
		return
	}

	if err := accService.Deposit(account, req.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "deposit successful"})
}

func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	accountID, _ := strconv.Atoi(vars["accId"])

	var req struct {
		Amount float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := accService.GetUserWithAccounts(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := accService.GetAccountByID(user, accountID)
	if err != nil {
		http.Error(w, "account not found or doesn't belong to the user", http.StatusBadRequest)
		return
	}

	if err := accService.Withdraw(account, req.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "withdrawal successful"})
}

func TransferHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	user, err := accService.GetUserWithAccounts(userId) // Fetch user with accounts
	if err != nil {
		http.Error(w, "Invalid user", http.StatusBadRequest)
		return
	}

	accountID := vars["accId"]
	accIDint, err := strconv.Atoi(accountID)
	if err != nil {
		http.Error(w, "Invalid type for account ID", http.StatusBadRequest)
		return
	}

	var request struct {
		FromBankID int     `json:"fromBankId"`
		Amount     float64 `json:"amount"`
		ToUserName string  `json:"toUserId"`
		ToBankID   int     `json:"toBankId"`
		ToAccNo    int     `json:"toAccNo"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	toUser, err := accService.GetUserWithAccounts(request.ToUserName)
	if err != nil {
		http.Error(w, "Invalid target user", http.StatusBadRequest)
		return
	}

	account, err := accService.GetAccountByIDForTransfer(user, request.FromBankID, accIDint)
	if err != nil {
		http.Error(w, "No such account found for the user", http.StatusBadRequest)
		return
	}

	toAccount, err := accService.GetAccountByIDForTransfer(toUser, request.ToBankID, request.ToAccNo)
	if err != nil {
		http.Error(w, "No such account found for target user", http.StatusBadRequest)
		return
	}

	if err := accService.Transfer(account, toAccount, request.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Transfer successful"})
}

func GetAccountPassbookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	accountID, _ := strconv.Atoi(vars["accId"])

	user, err := accService.GetUserWithAccounts(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := accService.GetAccountByID(user, accountID)
	if err != nil {
		http.Error(w, "account not found or doesn't belong to the user", http.StatusBadRequest)
		return
	}

	passbook, err := accService.GetAccountPassbook(int(account.ID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(passbook)
}
