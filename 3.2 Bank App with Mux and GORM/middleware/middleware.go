package middleware

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"
	"user/components/user/service"
	"user/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var secretKey = []byte("it'sDevthedev")

type Claims struct {
	UserID   int    `json:"userId"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func NewClaims(userID int, password string, expirationDate time.Time) *Claims {
	return &Claims{
		UserID:   userID,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationDate.Unix(),
		},
	}
}

func (c *Claims) Signing() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(secretKey)
}

func parseToken(r *http.Request) (*Claims, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return nil, errors.New("token not found")
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func TokenAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := parseToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func checkUserExistence(claims *Claims) (*models.User, error) {
	users, err := service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.UserID == claims.UserID && user.Password == claims.Password && user.IsActive {
			return user, nil
		}
	}
	return nil, errors.New("no such user found")
}

func VerifyAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := r.Context().Value("claims").(*Claims)
		user, err := checkUserExistence(claims)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if !user.IsAdmin {
			http.Error(w, "Unauthorized: admin access required", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func VerifyCustomer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := r.Context().Value("claims").(*Claims)
		user, err := checkUserExistence(claims)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if user.IsAdmin {
			http.Error(w, "Unauthorized: customer access required", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func VerifyCustomerFunctions(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := r.Context().Value("claims").(*Claims)
		vars := mux.Vars(r)
		userID, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Bad request: user ID must be an integer", http.StatusBadRequest)
			return
		}
		if claims.UserID != userID {
			http.Error(w, "Unauthorized: can only CRUD own accounts", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
