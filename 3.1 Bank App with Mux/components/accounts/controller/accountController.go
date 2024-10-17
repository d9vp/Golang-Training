package controller

// import (
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// 	"strconv"
// 	"user/components/accounts/service"

// 	"github.com/gorilla/mux"
// )

// func NewAccountHandler(w http.ResponseWriter, r *http.Request) (*service.Account, error) {
// 	var req struct {
// 		BankID  int     `json:"bank_id"`
// 		Balance float64 `json:"balance"`
// 	}
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return nil, err
// 	}

// 	if req.Balance < 1000 {
// 		err := errors.New("initial balance must be at least Rs. 1000")
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return nil, err
// 	}

// 	account, err := service.CreateAccount(req.BankID, req.Balance)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return nil, err
// 	}
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(account)
// 	return account, nil
// }

// func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	userID := vars["user_id"]
// 	user, err := getUserByID(userID)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	err = isNotAdmin(user)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	accountID := vars["acc_id"]
// 	bankID := vars["bank_id"]
// 	accIDint, ok := strconv.Atoi(accountID)
// 	if ok != nil {
// 		http.Error(w, "invalid type for user id", http.StatusBadRequest)
// 	}
// 	bankIDint, ok := strconv.Atoi(bankID)
// 	if ok != nil {
// 		http.Error(w, "invalid type for bank id", http.StatusBadRequest)
// 	}
// 	for _, acc := range user.Accounts {
// 		if acc.ID == accIDint && acc.BankID == bankIDint {
// 			accService.DeleteAccount(acc)
// 			w.WriteHeader(http.StatusNoContent)
// 			return
// 		}
// 	}

// 	http.Error(w, "no such account found for user", http.StatusBadRequest)
// }

// func Deposit(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	userID := vars["user_id"]
// 	user, err := getUserByID(userID)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	err = isNotAdmin(user)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	accountID := vars["acc_id"]
// 	bankID := vars["bank_id"]
// 	accIDint, ok := strconv.Atoi(accountID)
// 	if ok != nil {
// 		http.Error(w, "invalid type for user id", http.StatusBadRequest)
// 	}
// 	bankIDint, ok := strconv.Atoi(bankID)
// 	if ok != nil {
// 		http.Error(w, "invalid type for bank id", http.StatusBadRequest)
// 	}

// 	var request struct {
// 		Amount float64 `json:"amount"`
// 	}

// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}

// 	for _, acc := range user.Accounts {
// 		if acc.ID == accIDint && acc.BankID == bankIDint {
// 			err := accService.Deposit(acc, request.Amount)
// 			if err != nil {
// 				http.Error(w, err.Error(), http.StatusBadRequest)
// 				return
// 			}
// 			json.NewEncoder(w).Encode(map[string]string{"message": "Deposit successful"})
// 			w.WriteHeader(http.StatusNoContent)
// 			return
// 		}
// 	}
// 	http.Error(w, "no such account found for user", http.StatusBadRequest)
// }

// func Withdraw(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	userID := vars["user_id"]
// 	user, err := getUserByID(userID)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	err = isNotAdmin(user)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	accountID := vars["acc_id"]
// 	bankID := vars["bank_id"]
// 	accIDint, ok := strconv.Atoi(accountID)
// 	if ok != nil {
// 		http.Error(w, "invalid type for user id", http.StatusBadRequest)
// 	}
// 	bankIDint, ok := strconv.Atoi(bankID)
// 	if ok != nil {
// 		http.Error(w, "invalid type for bank id", http.StatusBadRequest)
// 	}

// 	var request struct {
// 		Amount float64 `json:"amount"`
// 	}

// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}

// 	for _, acc := range user.Accounts {
// 		if acc.ID == accIDint && acc.BankID == bankIDint {
// 			err := accService.Withdraw(acc, request.Amount)
// 			if err != nil {
// 				http.Error(w, err.Error(), http.StatusBadRequest)
// 				return
// 			}
// 			json.NewEncoder(w).Encode(map[string]string{"message": "Withdraw successful"})
// 			w.WriteHeader(http.StatusNoContent)
// 			return
// 		}
// 	}

// 	http.Error(w, "no such account found for user", http.StatusBadRequest)
// }

// func Transfer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	userID := vars["user_id"]
// 	user, err := getUserByID(userID)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	err = isNotAdmin(user)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	accountID := vars["acc_id"]
// 	bankID := vars["bank_id"]
// 	accIDint, ok := strconv.Atoi(accountID)
// 	if ok != nil {
// 		http.Error(w, "invalid type for user id", http.StatusBadRequest)
// 	}
// 	bankIDint, ok := strconv.Atoi(bankID)
// 	if ok != nil {
// 		http.Error(w, "invalid type for bank id", http.StatusBadRequest)
// 	}

// 	var request struct {
// 		Amount   float64 `json:"amount"`
// 		ToUserID int     `json:"to_userid"`
// 		ToBankID int     `json:"to_bankid"`
// 		ToAccNo  int     `json:"to_accno"`
// 	}

// 	toUser, err := getUserByID(strconv.Itoa(request.ToUserID))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	var account, toAccount *accService.Account
// 	for _, acc := range user.Accounts {
// 		if acc.ID == accIDint && acc.BankID == bankIDint {
// 			account = acc
// 		}
// 	}
// 	if account == nil {
// 		http.Error(w, "no such account found for user", http.StatusBadRequest)
// 		return
// 	}

// 	for _, acc := range toUser.Accounts {
// 		if acc.ID == request.ToAccNo && acc.BankID == request.ToBankID {
// 			toAccount = acc
// 		}
// 	}

// 	if toAccount == nil {
// 		http.Error(w, "no such account found for corresponding user", http.StatusBadRequest)
// 		return
// 	}

// 	accService.Withdraw(account, request.Amount)
// 	accService.Deposit(toAccount, request.Amount)

// 	json.NewEncoder(w).Encode(map[string]string{"message": "Transfer successful"})
// }
