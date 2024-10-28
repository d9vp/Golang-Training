package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	UserName  string     `gorm:"not null" json:"userName"`
	FirstName string     `gorm:"not null" json:"firstName"`
	LastName  string     `gorm:"not null" json:"lastName"`
	IsAdmin   bool       `gorm:"default:false" json:"isAdmin"`
	IsActive  bool       `gorm:"default:true" json:"isActive"`
	Password  string     `gorm:"not null" json:"-"`
	Accounts  []*Account `gorm:"foreignKey:UserID;references:ID" json:"accounts"`
}

type Bank struct {
	gorm.Model
	FullName     string        `gorm:"not null" json:"fullName"`
	Abbreviation string        `gorm:"size:5;not null" json:"abbreviation"`
	IsActive     bool          `gorm:"default:true" json:"isActive"`
	Accounts     []*Account    `gorm:"foreignKey:BankID;references:ID" json:"accounts"`
	LedgerData   []*LedgerData `gorm:"foreignKey:BankID;references:ID" json:"ledgerData"`
}

type LedgerData struct {
	gorm.Model
	BankID              int     `gorm:"not null" json:"bankId"`
	CorrespondingBankID int     `gorm:"not null" json:"correspondingBankId"`
	Amount              float64 `gorm:"not null" json:"amount"`
}

type Account struct {
	gorm.Model
	UserID   int                 `gorm:"not null" json:"userId"`
	BankID   int                 `gorm:"not null" json:"bankId"`
	Balance  float64             `gorm:"not null;default:1000" json:"balance"`
	IsActive bool                `gorm:"not null;default:true" json:"isActive"`
	Passbook []*TransactionEntry `gorm:"foreignKey:AccountID;references:ID" json:"passbook"`
}

type TransactionEntry struct {
	gorm.Model

	Type                    string    `json:"type"`
	Amount                  float64   `json:"amount"`
	BalanceAfterTransaction float64   `json:"balanceAfterTransaction"`
	CorrespondingBankID     int       `json:"correspondingBankId"`
	CorrespondingAccountID  int       `json:"correspondingAccountId"`
	Timestamp               time.Time `gorm:"autoCreateTime" json:"timestamp"`
	AccountID               int       `gorm:"not null" json:"accountId"`
}

func InitDB() {
	dsn := "root:{password}@tcp(127.0.0.1:3306)/GoBankingApp?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(&User{}, &Bank{}, &LedgerData{}, &Account{}, &TransactionEntry{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	fmt.Println("Database connection successfully established!")
	AddSuperAdmin()
}

func SetupDatabase() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("Database connection details not set in environment variables")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(&User{}, &Bank{}, &LedgerData{}, &Account{}, &TransactionEntry{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	fmt.Println("Connected to the database and schema migrated successfully!")
}

func AddSuperAdmin() {
	superAdmin := &User{
		FirstName: "Super",
		LastName:  "Admin",
		UserName:  "SuperAdmin",
		IsAdmin:   true,
		IsActive:  true,
		Password:  "password",
	}

	if err := DB.Create(&superAdmin).Error; err != nil {
		log.Fatalf("failed to insert initial user: %v", err)
	}

}

func ClearDatabase() {
	dsn := "root:{password}@tcp(127.0.0.1:3306)/GoBankingApp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.Migrator().DropTable(&User{}, &Bank{}, &Account{}, &TransactionEntry{}, &LedgerData{}); err != nil {
		log.Fatalf("failed to drop tables: %v", err)
	}

	log.Println("All tables dropped successfully!")
}
