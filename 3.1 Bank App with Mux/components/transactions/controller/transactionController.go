package controller

// import (
// 	"encoding/json"
// 	"net/http"
// 	"strconv"
// 	"user/components/transactions/service"

// 	"github.com/gorilla/mux"
// )

// // Get all transactions
// func GetTransactions(w http.ResponseWriter, r *http.Request) {
// 	transactions := service.GetAllTransactions()
// 	json.NewEncoder(w).Encode(transactions)
// }

// // Get a specific transaction by ID
// func GetTransaction(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	transactionID, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
// 		return
// 	}

// 	transaction, err := service.GetTransactionByID(transactionID)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusNotFound)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(transaction)
// }

// // Create a new transaction
// func CreateTransaction(w http.ResponseWriter, r *http.Request) {
// 	var transaction service.Transaction
// 	err := json.NewDecoder(r.Body).Decode(&transaction)
// 	if err != nil {
// 		http.Error(w, "Invalid input", http.StatusBadRequest)
// 		return
// 	}

// 	err = service.CreateTransaction(&transaction)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(transaction)
// }
