package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	accService "user/components/accounts/service"
	"user/components/user/service"

	"github.com/gorilla/mux"
)

// Utility function to get user by ID
func getUserByID(userID string) (*service.User, error) {
	id, err := strconv.Atoi(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID")
	}

	for _, user := range service.AllUsers {
		if user.UserID == id {
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
		Password  string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := service.NewAdmin(req.FirstName, req.LastName, req.Password)
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
		Password  string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := service.NewUser(req.FirstName, req.LastName, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Get All Users Handler
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := service.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	user, err := getUserByID(userID)
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
	userID := vars["id"]

	_, err := getUserByID(userID)
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

	userIDint, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = service.UpdateUsers(userIDint, req.Parameter, req.NewValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete User Handler
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	_, err := getUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userIDint, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = service.DeleteUsers(userIDint)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// New Account Handler - Create a new account for user
func NewAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := getUserByID(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req struct {
		BankID  int     `json:"bankId"`
		Balance float64 `json:"balance"`
	}

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Balance < 1000 {
		http.Error(w, "Initial balance must be at least Rs. 1000", http.StatusBadRequest)
		return
	}

	account, err := accService.CreateAccount(req.BankID, req.Balance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Accounts = append(user.Accounts, account)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

// Get Accounts Handler - Retrieve user accounts
func GetAccountsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	user, err := getUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user.Accounts)
}

// Delete Account Handler
func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	user, err := getUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accountID := vars["accId"]
	var request struct {
		BankID int `json:"bankID"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	accIDint, err := strconv.Atoi(accountID)
	fmt.Println(accountID, accIDint, err)
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}
	for _, acc := range user.Accounts {
		if acc.ID == accIDint && acc.BankID == request.BankID {
			accService.DeleteAccount(acc)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "No such account found for user", http.StatusBadRequest)
}

// Deposit Handler
func Deposit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	user, err := getUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accountID := vars["accId"]
	accIDint, err := strconv.Atoi(accountID)
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	var request struct {
		Amount float64 `json:"amount"`
		BankID int     `json:"bankID"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	for _, acc := range user.Accounts {
		if acc.ID == accIDint && acc.BankID == request.BankID {
			err := accService.Deposit(acc, request.Amount)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			json.NewEncoder(w).Encode(map[string]string{"message": "Deposit successful"})
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "No such account found for user", http.StatusBadRequest)
}

// Withdraw Handler
func Withdraw(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	user, err := getUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accountID := vars["accId"]

	accIDint, err := strconv.Atoi(accountID)
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	var request struct {
		Amount float64 `json:"amount"`
		BankID int     `json:"bankID"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	for _, acc := range user.Accounts {
		if acc.ID == accIDint && acc.BankID == request.BankID {
			err := accService.Withdraw(acc, request.Amount)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			json.NewEncoder(w).Encode(map[string]string{"message": "Withdrawal successful"})
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "No such account found for user", http.StatusBadRequest)
}

func Transfer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	user, err := getUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		ToUserID   int     `json:"toUserId"`
		ToBankID   int     `json:"toBankId"`
		ToAccNo    int     `json:"toAccNo"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Fetch target user
	toUser, err := getUserByID(strconv.Itoa(request.ToUserID))
	if err != nil {
		http.Error(w, "Invalid target user", http.StatusBadRequest)
		return
	}

	// Fetch account and toAccount
	var account, toAccount *accService.Account
	for _, acc := range user.Accounts {
		if acc.ID == accIDint && acc.BankID == request.FromBankID {
			account = acc
			break
		}
	}
	if account == nil {
		http.Error(w, "No such account found for user", http.StatusBadRequest)
		return
	}

	for _, acc := range toUser.Accounts {
		if acc.ID == request.ToAccNo && acc.BankID == request.ToBankID {
			toAccount = acc
			break
		}
	}
	if toAccount == nil {
		http.Error(w, "No such account found for target user", http.StatusBadRequest)
		return
	}

	// Perform the transfer
	if err := accService.Withdraw(account, request.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := accService.Deposit(toAccount, request.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Transfer successful"})
}

func GetTotalBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	user, err := getUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	totalBalance := 0.0
	for _, acc := range service.GetAccountsForUser(user) {
		totalBalance += acc.Balance
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]float64{"Total Balance for user": totalBalance})
}

func GetAccountPassbook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	user, err := getUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	accountID := vars["accId"]
	accIDint, err := strconv.Atoi(accountID)
	if err != nil {
		http.Error(w, "Invalid type for account ID", http.StatusBadRequest)
		return
	}
	var request struct {
		BankID int `json:"bankId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	for _, acc := range service.GetAccountsForUser(user) {
		if acc.BankID == request.BankID && acc.ID == accIDint {
			passbook := accService.GetAccountPassbook(acc)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(passbook)
			return
		}
	}
	http.Error(w, "no such account for user", http.StatusBadRequest)
}
