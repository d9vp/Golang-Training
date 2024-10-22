package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	bankController "user/components/bank/controller"
	userController "user/components/user/controller"
	"user/components/user/service"
	"user/middleware"

	"user/models"

	"github.com/gorilla/mux"
)

func main() {
	models.ClearDatabase()
	models.InitDB()

	router := mux.NewRouter()
	router.HandleFunc("/banking/login", login).Methods(http.MethodPost)

	adminRoutes := router.PathPrefix("/banking/admin").Subrouter()
	adminRoutes.Use(middleware.TokenAuthorization)
	adminRoutes.Use(middleware.VerifyAdmin)

	adminRoutes.HandleFunc("", userController.NewAdminHandler).Methods("POST")
	adminRoutes.HandleFunc("/users", userController.NewUserHandler).Methods("POST")
	adminRoutes.HandleFunc("/users", userController.GetUsersHandler).Methods("GET")
	adminRoutes.HandleFunc("/users/{id}", userController.GetUserByIDHandler).Methods("GET")
	adminRoutes.HandleFunc("/users/{id}", userController.UpdateUserHandler).Methods("PUT")
	adminRoutes.HandleFunc("/users/{id}", userController.DeleteUserHandler).Methods("DELETE")

	adminRoutes.HandleFunc("/banks", bankController.NewBankHandler).Methods("POST")
	adminRoutes.HandleFunc("/banks", bankController.GetBanksHandler).Methods("GET")
	adminRoutes.HandleFunc("/banks/{id}", bankController.UpdateBankHandler).Methods("PUT")
	adminRoutes.HandleFunc("/banks/{id}", bankController.DeleteBankHandler).Methods("DELETE")

	adminRoutes.HandleFunc("/banks/ledger/{id}", bankController.AddToLedger).Methods("PUT")
	adminRoutes.HandleFunc("/banks/ledger/{id}", bankController.GetLedger).Methods("GET")

	customerRoutes := router.PathPrefix("/banking/customer/{id}/accounts").Subrouter()

	customerRoutes.Use(middleware.TokenAuthorization)
	customerRoutes.Use(middleware.VerifyCustomer)
	customerRoutes.Use(middleware.VerifyCustomerFunctions)

	customerRoutes.HandleFunc("", userController.NewAccountHandler).Methods("POST")
	customerRoutes.HandleFunc("", userController.GetAccountsHandler).Methods("GET")
	customerRoutes.HandleFunc("/total-balance", userController.GetTotalBalance).Methods("GET")
	customerRoutes.HandleFunc("/{accId}", userController.DeleteAccountHandler).Methods("DELETE")
	customerRoutes.HandleFunc("/{accId}", userController.GetAccountPassbookHandler).Methods("GET")
	customerRoutes.HandleFunc("/deposit/{accId}", userController.DepositHandler).Methods("PUT")
	customerRoutes.HandleFunc("/withdraw/{accId}", userController.WithdrawHandler).Methods("PUT")
	customerRoutes.HandleFunc("/transfer/{accId}", userController.TransferHandler).Methods("PUT")

	log.Println("Server is running on :4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}

func login(w http.ResponseWriter, r *http.Request) {
	var user middleware.Claims
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
		return
	}

	users, _ := service.GetAllUsers()
	for _, u := range users {
		if u.UserName == user.UserName && u.Password == user.Password {
			claim := middleware.NewClaims(user.UserName, user.Password, time.Now().Add(time.Hour*200))
			token, err := claim.Signing()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{"error": "Could not generate token"})
				return
			}

			w.Header().Set("Authorization", token)
			json.NewEncoder(w).Encode(map[string]string{"token": token})
			return
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
}
