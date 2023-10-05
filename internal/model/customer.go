package models

import (
	"github.com/go-sql-driver/mysql"
	"time"
)

// CustomerLoan adalah model struct untuk tabel customer_loan
type CustomerLoan struct {
	ID           int        `json:"id"`
	CustomerID   int        `json:"customer_id"`
	Status       string     `json:"status"`
	LoanDate     time.Time  `json:"loan_date"`
	Amount       float64    `json:"amount"`
	ApprovedDate time.Time  `json:"approved_date,omitempty"`
	Tenor        int        `json:"tenor"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

// Customer adalah model struct untuk tabel customer
type Customer struct {
	ID         int            `json:"id" db:"id"`
	NIK        string         `json:"NIK" db:"NIK"`
	FullName   string         `json:"full_name" db:"full_name"`
	LegalName  string         `json:"legal_name,omitempty" db:"legal_name"`
	BornPlace  string         `json:"born_place,omitempty" db:"born_place"`
	BornDate   mysql.NullTime `json:"born_date,omitempty" db:"born_date"`
	Salary     float64        `json:"salary,omitempty" db:"salary"`
	Email      string         `json:"email" db:"email"`
	IsAdmin    bool           `json:"is_admin,omitempty" db:"is_admin"`
	Password   string         `json:"password" db:"password"`
	FotoSelfie string         `json:"foto_selfie,omitempty" db:"foto_selfie"`
	FotoKTP    string         `json:"foto_ktp,omitempty" db:"foto_ktp"`
	CreatedAt  mysql.NullTime `json:"created_at" db:"created_at"`
	UpdatedAt  mysql.NullTime `json:"updated_at" db:"updated_at"`
	DeletedAt  *time.Time     `json:"deleted_at,omitempty" db:"deleted_at"`
}

// CustomerParam adalah model struct untuk request
type CustomerParam struct {
	NIK        string    `json:"NIK" validate:"required"`
	FullName   string    `json:"full_name" validate:"required"`
	LegalName  string    `json:"legal_name,omitempty" validate:"required"`
	BornPlace  string    `json:"born_place,omitempty" validate:"required"`
	BornDate   time.Time `json:"born_date,omitempty" validate:"required"`
	Salary     float64   `json:"salary,omitempty" validate:"required"`
	Email      string    `json:"email" validate:"required,email"`
	Password   string    `json:"password" validate:"required"`
	FotoSelfie string    `json:"foto_selfie,omitempty"`
	FotoKTP    string    `json:"foto_ktp,omitempty"`
}

// LoanPayment adalah model struct untuk tabel loan_payment
type LoanPayment struct {
	ID             int        `json:"id"`
	CustomerLoanID int        `json:"customer_loan_id"`
	PaymentDate    time.Time  `json:"payment_date"`
	Amount         float64    `json:"amount"`
	Status         string     `json:"status"`
	DueDate        time.Time  `json:"due_date"`
	LateFee        float64    `json:"late_fee"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
}
