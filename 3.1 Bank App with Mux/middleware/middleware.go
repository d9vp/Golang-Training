package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"user/components/user/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var secretKey = []byte("helloWorld")

type Claims struct {
	UserID   int    `json:"userId"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func NewClaims(userID int,
	password string,
	expirationDate time.Time) *Claims {
	return &Claims{
		UserID:   userID,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationDate.Unix(),
		},
	}
}

func (c *Claims) Signing() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, _ := token.SignedString(secretKey)
	return tokenString
}

func TokenAuthorisation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Token not found")
			return
		}

		claim := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claim)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func VerifyAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := r.Context().Value("claims").(*Claims)
		for _, user := range service.AllUsers {
			if user.UserID == claims.UserID && user.Password == claims.Password && user.IsActive {
				// Check if the user is an admin
				if claims == nil || !user.IsAdmin {
					http.Error(w, "Unauthorized: admin access required", http.StatusUnauthorized)
					return
				}
				next.ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, "no such user found", http.StatusUnauthorized)
	})
}

func VerifyCustomer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := r.Context().Value("claims").(*Claims)

		for _, user := range service.AllUsers {
			if user.UserID == claims.UserID && user.Password == claims.Password && user.IsActive {
				// Check if the user is an admin
				if claims == nil || user.IsAdmin {
					http.Error(w, "Unauthorized: customer access required", http.StatusUnauthorized)
					return
				}
				next.ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, "no such user found", http.StatusUnauthorized)

		next.ServeHTTP(w, r)
	})
}

func VerifyCustomerFunctions(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := r.Context().Value("claims").(*Claims)
		vars := mux.Vars(r)
		userID := vars["id"]

		id, err := strconv.Atoi(userID)
		if err != nil {
			http.Error(w, "Bad request: user id type should be integer", http.StatusBadRequest)
			return
		}
		if claims.UserID != id {
			http.Error(w, "Unauthorized: can only CRUD on own accounts", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
