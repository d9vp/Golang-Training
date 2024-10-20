package user

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

// Utility function to get user by ID
func getUserByUserName(userName string) (*models.User, error) {

	users, err := service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.UserName == userName {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

// Admin Handler - Create a new admin
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

// User Handler - Create a new user
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

// Get All Users Handler
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
	userName := vars["userName"]

	user, err := getUserByUserName(userName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusOK)
}

// Update User Handler
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["userName"]

	_, err := getUserByUserName(userName)
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

	err = service.UpdateUsers(userName, req.Parameter, req.NewValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete User Handler
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["userName"]

	_, err := getUserByUserName(userName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.DeleteUsers(userName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// New Account Handler - Create a new account for user
func NewAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["userName"]

	// Get user by ID
	user, err := getUserByUserName(userName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Decode request payload
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

	// Create account
	account, err := accService.CreateAccount(user, req.BankID, req.Balance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Respond with the newly created account
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

// Get Accounts Handler - Retrieve user accounts
func GetAccountsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["userName"]

	// Fetch user by username and preload their accounts
	user, err := accService.GetUserWithAccounts(userName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return user's accounts
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user.Accounts)
}

// Get Total Balance - Retrieve total balance across all user accounts
func GetTotalBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["userName"]

	// Fetch total balance from account service
	totalBalance, err := accService.GetTotalBalanceForUser(userName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return the total balance
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]float64{"Total Balance for user": totalBalance})
}

// DeleteAccountHandler - Deactivate an account by setting its status to inactive
func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["userName"]
	accountID, _ := strconv.Atoi(vars["accId"])

	// Get user and preload their accounts
	user, err := accService.GetUserWithAccounts(userName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := accService.GetAccountByID(user, accountID)
	if err != nil {
		http.Error(w, "account not found or doesn't belong to the user", http.StatusBadRequest)
		return
	}

	// Deactivate the account
	if err := accService.DeleteAccount(account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "account deactivated successfully"})
}

// DepositHandler - Deposit money into an account
func DepositHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["userName"]
	accountID, _ := strconv.Atoi(vars["accId"])

	var req struct {
		Amount float64 `json:"amount"`
	}

	// Decode deposit amount from the request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Get user and preload their accounts
	user, err := accService.GetUserWithAccounts(userName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := accService.GetAccountByID(user, accountID)
	if err != nil {
		http.Error(w, "account not found or doesn't belong to the user", http.StatusBadRequest)
		return
	}

	// Perform the deposit
	if err := accService.Deposit(account, req.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "deposit successful"})
}

// WithdrawHandler - Withdraw money from an account
func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["userName"]
	accountID, _ := strconv.Atoi(vars["accId"])

	var req struct {
		Amount float64 `json:"amount"`
	}

	// Decode withdrawal amount from the request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Get user and preload their accounts
	user, err := accService.GetUserWithAccounts(userName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := accService.GetAccountByID(user, accountID)
	if err != nil {
		http.Error(w, "account not found or doesn't belong to the user", http.StatusBadRequest)
		return
	}

	// Perform the withdrawal
	if err := accService.Withdraw(account, req.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "withdrawal successful"})
}

// TransferHandler - Transfer money from one account to another
func TransferHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["userName"]
	user, err := accService.GetUserWithAccounts(userName) // Fetch user with accounts
	if err != nil {
		http.Error(w, "Invalid user", http.StatusBadRequest)
		return
	}

	// Convert account ID from string to int
	accountID := vars["accId"]
	accIDint, err := strconv.Atoi(accountID)
	if err != nil {
		http.Error(w, "Invalid type for account ID", http.StatusBadRequest)
		return
	}

	// Decode request payload
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

	// Fetch target user and preload their accounts
	toUser, err := accService.GetUserWithAccounts(request.ToUserName)
	if err != nil {
		http.Error(w, "Invalid target user", http.StatusBadRequest)
		return
	}

	// Get the sender and receiver accounts
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

	// Perform the transfer
	if err := accService.Transfer(account, toAccount, request.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Transfer successful"})
}

// GetAccountPassbookHandler - Retrieve transaction history for an account
func GetAccountPassbookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["userName"]
	accountID, _ := strconv.Atoi(vars["accId"])

	// Get user and preload their accounts
	user, err := accService.GetUserWithAccounts(userName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := accService.GetAccountByID(user, accountID)
	if err != nil {
		http.Error(w, "account not found or doesn't belong to the user", http.StatusBadRequest)
		return
	}

	// Fetch passbook (transaction history)
	passbook, err := accService.GetAccountPassbook(account.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return passbook as JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(passbook)
}
